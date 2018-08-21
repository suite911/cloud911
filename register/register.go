package register

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/suite911/cloud911/database"
	"github.com/suite911/cloud911/vars"

	pkgErrors "github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

type GoogleCaptchaResponse struct {
	Success        bool        `json:"success"`
	ChallengeTS    string      `json:"challenge_ts"`
	HostName       string      `json:"hostname"`
	APKPackageName string      `json:"apk_package_name"`
	ErrorCodes     interface{} `json:"error-codes"`
}

func Try(ctx *fasthttp.RequestCtx) (attempt bool, err error) {
	attempt = ctx.IsPost()
	if !attempt {
		return
	}
	argsRecv := ctx.PostArgs()
	email := argsRecv.Peek("email")
	captcha := argsRecv.Peek("g-recaptcha-response")
	if len(email) < 0 {
		ctx.Redirect("?email=missing", 302)
		return
	}
	if len(captcha) < 0 {
		ctx.Redirect("?captcha=missing", 302)
		return
	}
	var args fasthttp.Args
	args.SetBytesV("secret", vars.Pass.CaptchaSecret)
	args.SetBytesV("response", captcha)
	args.SetBytesV("remoteip", ctx.RequestURI())
	var statusCode int
	var body []byte
	if statusCode, body, err = fasthttp.Post(nil, "https://www.google.com/recaptcha/api/siteverify", &args); err != nil {
		return
	}
	if statusCode != 200 {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, string(body))
		return attempt, pkgErrors.Wrap(errors.New(strconv.Itoa(statusCode)), "fasthttp.Post")
	}
	var resp GoogleCaptchaResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, resp)
		return attempt, pkgErrors.Wrap(err, "json.Unmarshal(body, &resp)")
	}
	if !resp.Success {
		ctx.Redirect("?captcha=failed", 302)
		return
	}
	var url string
	if url, err = database.Register(email); err != nil {
		if len(url) > 0 {
			ctx.Redirect(url, 302)
		}
		return
	}
	if url := vars.Pass.Registered; len(url) > 0 {
		ctx.Redirect(url, 302)
		return
	}
	ctx.Redirect("?registered=true", 302)
	return
}
