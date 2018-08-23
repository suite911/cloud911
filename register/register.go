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

func Try(ctx *fasthttp.RequestCtx, email, captcha []byte) (attempt bool, err error) {


}
