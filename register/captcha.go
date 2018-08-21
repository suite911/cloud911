package register

import (
	"io/ioutil"

	"github.com/suite911/cloud911/vars"

	"github.com/pkg/errors"
)

func LoadCaptchaSecret() error {
	s, err := ioutil.ReadFile(vars.CaptchaSecretPath)
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadFile(vars.CaptchaSecretPath)")
	}
	for len(s) > 0 && s[len(s) - 1] == '\n' {
		s = s[:len(s) - 1]
	}
	vars.Pass.CaptchaSecret = s
	return nil
}
