package cloud911

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/suite911/cloud911/run"
	"github.com/suite911/cloud911/vars"
)

func child(fns []func() error) error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &vars.Pass); err != nil {
		return err
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return run.Listen()
}
