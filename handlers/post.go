package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"unicode/utf8"

	"github.com/suite911/cloud911/database"
	"github.com/suite911/cloud911/vars"

	"github.com/suite911/maths911/maths"

	"github.com/badoux/checkmail"
	pkgErrors "github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var OverridePost func(*fasthttp.RequestCtx)

func Post(ctx *fasthttp.RequestCtx) {
	switch {
	case OverridePost == nil:
		post(ctx)
	default:
		OverridePost(ctx)
	}
}

func post(ctx *fasthttp.RequestCtx) {
	args := ctx.PostArgs()
	var action, email string
	{
		actionBytes := args.Peek("action")
		if !utf8.Valid(actionBytes) {
			actionBytes = actionBytes[:0]
		}
		action = string(actionBytes)
	}

	if emailBytes := args.Peek("email"); len(emailBytes) > 0 {
		if !utf8.Valid(emailBytes) {
			ctx.Redirect("#email-invalid", 302)
			return
		}
		email = string(emailBytes)
		if err := checkmail.ValidateFormat(email); err != nil {
			ctx.Redirect("#email-invalid", 302)
			return
		}
		if err := checkmail.ValidateHost(email); err != nil {
			ctx.Redirect("#email-invalid", 302)
			return
		}
	}

	if childAccount := args.Peek("child-account"); len(childAccount) > 0 {
		childTypeBytes := args.Peek("child-type")
		if !utf8.Valid(childTypeBytes) {
			ctx.Redirect("#something-went-wrong", 302)
			return
		}
		switch string(childTypeBytes) {
		case "under-13-us":
			if perm := args.Peek("child-perm-under-13-us"); len(perm) < 1 {
				ctx.Redirect("#permission-under-13-us-missing", 302)
				return
			}
		case "under-16-eu":
			if perm := args.Peek("child-perm-under-16-eu"); len(perm) < 1 {
				ctx.Redirect("#permission-under-16-eu-missing", 302)
				return
			}
		case "under-18":
		case "default":
			ctx.Redirect("#child-type-missing", 302)
			return
		default:
			ctx.Redirect("#something-went-wrong", 302)
			return
		}
	}

	switch action {
	case "register":
		if len(email) < 1 {
			ctx.Redirect("#email-missing", 302)
			return
		}
		netScore := 1.0
		if len(vars.Pass.CaptchaSecret) > 0 {
			captchaOnLoad := args.Peek("captcha-onload")
			captchaOnChange := args.Peek("captcha-onchange")
			captchaOnSubmit := args.Peek("captcha-onsubmit")
			for _, captcha := range [][]byte{captchaOnLoad, captchaOnChange, captchaOnSubmit} {
				if len(captcha) < 1 {
					ctx.Redirect("#captcha-missing", 302)
					return
				}
				score, err := VerifyCaptchaSolution(ctx, captcha)
				if err != nil {
					ctx.Redirect("#captcha-not-working", 302)
					return
				}
				if score < vars.CaptchaThresholdRegister {
					ctx.Redirect("#captcha-failed", 302)
					return
				}
				score *= netScore
			}
		}

		ctx.Redirect(database.Register(email, netScore), 302)
		return
	}
}

type GoogleCaptchaResponse struct {
	Success        bool        `json:"success"`
	Score          float32     `json:"score"`
	Action         string      `json:"action"`
	ChallengeTS    string      `json:"challenge_ts"`
	HostName       string      `json:"hostname"`
	APKPackageName string      `json:"apk_package_name"`
	ErrorCodes     interface{} `json:"error-codes"`
}

func VerifyCaptchaSolution(ctx *fasthttp.RequestCtx, solution []byte, action string) (float32, error) {
	var args fasthttp.Args
	args.SetBytesV("secret", vars.Pass.CaptchaSecret)
	args.SetBytesV("response", solution)
	args.SetBytesV("remoteip", ctx.RequestURI())
	/*
	var statusCode int
	var body []byte
	*/
	statusCode, body, err := fasthttp.Post(nil, "https://www.google.com/recaptcha/api/siteverify", &args)
	if err != nil {
		return false, pkgErrors.Wrap(err, "fasthttp.Post")
	}
	if statusCode != 200 {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, string(body))
		return false, pkgErrors.Wrap(errors.New(strconv.Itoa(statusCode)), "fasthttp.Post")
	}
	var resp = GoogleCaptchaResponse{Success: true, Score: 1.0}
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, resp)
		return false, pkgErrors.Wrap(err, "json.Unmarshal(body, &resp)")
	}
	if !resp.Success || resp.Action != action {
		return 0.0, nil
	}
	return maths.ClampFloat32(resp.Score, 0, 1), nil
}
