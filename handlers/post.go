package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
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
			ctx.Redirect("#encoding", 302)
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

	switch action {
	case "register":
		if len(email) < 1 {
			ctx.Redirect("#email-missing", 302)
			return
		}
		usernameBytes := args.Peek("username")
		if len(usernameBytes) < 1 {
			ctx.Redirect("#username-missing", 302)
			return
		}
		if !utf8.Valid(usernameBytes) {
			ctx.Redirect("#encoding", 302)
			return
		}
		username := string(usernameBytes)
		if len(username) > 64 {
			ctx.Redirect("#something-went-wrong-with-username", 302)
			return
		}
		username = strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1 // delete
			}
			if r == '@' {
				return -1 // delete
			}
			return r
		}, username)
		if len(username) < 1 {
			ctx.Redirect("#username-invalid", 302)
			return
		}
		hasNonDigit := false
		for i := 0; i < len(username); i++ {
			if username[i] < '0' || username[i] > '9' {
				hasNonDigit = true
				break
			}
		}
		if !hasNonDigit {
			ctx.Redirect("#username-invalid", 302)
			return
		}
		var minor bool
		switch ageClass := string(args.Peek("age-class")); ageClass {
		case "adult":
			minor = false
		case "minor":
			minor = true
		default:
			ctx.Redirect("#age-class-missing", 302)
			return
		}
		if minor {
			for _, checkbox := range []string{
				"parent-email-yours",
				"parent-read-terms",
				"parent-notify",
			} {
				value := string(args.Peek(checkbox))
				if value != "on" {
					ctx.Redirect("#child-account-permission-form-incomplete", 302)
					return
				}
			}
		}
		emwhoBytes := args.Peek("emergency-name")
		if !utf8.Valid(emwhoBytes) {
			ctx.Redirect("#encoding", 302)
			return
		}
		emwho := string(emwhoBytes)
		emhowBytes := args.Peek("emergency-contact")
		if !utf8.Valid(emhowBytes) {
			ctx.Redirect("#encoding", 302)
			return
		}
		emhow := string(emhowBytes)
		emrelBytes := args.Peek("emergency-relation")
		if !utf8.Valid(emrelBytes) {
			ctx.Redirect("#encoding", 302)
			return
		}
		emrel := string(emrelBytes)
		netScore := 1.0
		if len(vars.Pass.CaptchaSecret) > 0 {
			log.Printf("onload:   \t\"%s\"", string(args.Peek("captcha-onload")))
			log.Printf("onchange: \t\"%s\"", string(args.Peek("captcha-onchange")))
			log.Printf("onsubmit: \t\"%s\"", string(args.Peek("captcha-onsubmit")))
			for i, captchaAction := range []string{"load", "change", "submit"} {
				captcha := args.Peek("captcha-on" + captchaAction)
				if len(captcha) < 1 {
					ctx.Redirect("#something-went-wrong-with-captcha-code-1"+strconv.Itoa(i), 302)
					return
				}
				score, err := VerifyCaptchaSolution(ctx, captcha, captchaAction)
				if err != nil {
					ctx.Redirect("#something-went-wrong-with-captcha-code-2"+strconv.Itoa(i), 302)
					return
				}
				if score < vars.Pass.CaptchaThresholdRegister {
					ctx.Redirect("#something-went-wrong-with-captcha-code-3"+strconv.Itoa(i), 302)
					return
				}
				score *= netScore
			}
		}
		ctx.Redirect(database.Register(username, email, netScore, minor, emwho, emhow, emrel), 302)
		return
	}
}

type GoogleCaptchaResponse struct {
	Success        bool        `json:"success"`
	Score          float64     `json:"score"`
	Action         string      `json:"action"`
	ChallengeTS    string      `json:"challenge_ts"`
	HostName       string      `json:"hostname"`
	APKPackageName string      `json:"apk_package_name"`
	ErrorCodes     interface{} `json:"error-codes"`
}

func VerifyCaptchaSolution(ctx *fasthttp.RequestCtx, solution []byte, action string) (float64, error) {
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
		return 0.0, pkgErrors.Wrap(err, "fasthttp.Post")
	}
	if statusCode != 200 {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, string(body))
		return 0.0, pkgErrors.Wrap(errors.New(strconv.Itoa(statusCode)), "fasthttp.Post")
	}
	var resp = GoogleCaptchaResponse{Success: true, Score: 1.0}
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, resp)
		return 0.0, pkgErrors.Wrap(err, "json.Unmarshal(body, &resp)")
	}
	if !resp.Success || resp.Action != action {
		return 0.0, nil
	}
	return maths.ClampFloat64(resp.Score, 0, 1), nil
}
