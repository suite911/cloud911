package register

import (
	"encoding/json"
	"errors"

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

func Maybe(ctx *fasthttp.RequestCtx) (attempt bool, err error) {
	attempt = ctx.IsPost()
	if !attempt {
		return
	}
	argsRecv := ctx.PostArgs()
	email := argsRecv.Peek("email")
	captcha := argsRecv.Peek("g-recaptcha-response")
	var args fasthttp.Args
	args.Set("secret", vars.Pass.CaptchaSecret)
	args.SetBytesV("response", response)
	args.SetBytesV("remoteip", ctx.RequestURI())
	var statusCode int
	var body []byte
	if statusCode, body, err = fasthttp.Post(nil, "https://www.google.com/recaptcha/api/siteverify", args); err != nil {
		return
	}
	if statusCode != 200 {
		log.Printf("Google replied %d\n--- REPLY BODY ---\n%v\n--- END OF REPLY BODY ---", statusCode, string(body))
		return true, pkgErrors.Wrap(errors.New("non-200 status code"), "fasthttp.Post")
	}
	var resp GoogleCaptchaResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return
	}
	if success {
		ctx.Redirect("/success", 302)
	} else {
		ctx.Redirect("/failure", 302)
	}
	return
}
