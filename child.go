package cloud911

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/suite911/cloud911/run"
	"github.com/suite911/cloud911/vars"

	"github.com/pkg/errors"
)

func child(fns []func() error) error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadAll(os.Stdin)")
	}
	if err := json.Unmarshal(b, &vars.Pass); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return run.Listen()
}
