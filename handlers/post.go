package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"unicode/utf8"

	"github.com/suite911/cloud911/database"
	"github.com/suite911/cloud911/vars"

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
	captcha := args.Peek("g-recaptcha-response")

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

	if len(captcha) > 0 {
		pass, err := VerifyCaptchaSolution(ctx, captcha)
		if err != nil {
			ctx.Redirect("#captcha-not-working", 302)
			return
		}
		if !pass {
			ctx.Redirect("#captcha-failed", 302)
			return
		}
	}

	if childAccount := args.Peek("child-account"); true {
		log.Printf("child-account: <%T>: \"%v\"", childAccount, childAccount)
	}

	switch action {
	case "register":
		if len(email) < 1 {
			ctx.Redirect("#email-missing", 302)
			return
		}
		if len(captcha) < 1 {
			ctx.Redirect("#captcha-missing", 302)
			return
		}

		if url, err := database.Register(email); err != nil {
			if len(url) > 0 {
				ctx.Redirect(url, 302)
				return
			}
			ctx.Redirect("#something-went-wrong", 302)
			return
		}
		if url := vars.Pass.Registered; len(url) > 0 {
			ctx.Redirect(url, 302)
			return
		}
		ctx.Redirect("#registered", 302)
		return
	}
}

type GoogleCaptchaResponse struct {
	Success        bool        `json:"success"`
	ChallengeTS    string      `json:"challenge_ts"`
	HostName       string      `json:"hostname"`
	APKPackageName string      `json:"apk_package_name"`
	ErrorCodes     interface{} `json:"error-codes"`
}

func VerifyCaptchaSolution(ctx *fasthttp.RequestCtx, solution []byte) (bool, error) {
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
	var resp GoogleCaptchaResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, resp)
		return false, pkgErrors.Wrap(err, "json.Unmarshal(body, &resp)")
	}
	return resp.Success, nil
}
