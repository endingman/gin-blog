package app

import (
	"gin-blog/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) (msg string) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
		msg = err.Message
		break
	}

	return msg
}
