package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetFormatter(&nested.Formatter{
		HideKeys:      true,
		NoColors:      true,
		ShowFullLevel: true,
	})
}
