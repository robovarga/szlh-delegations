package config

import (
	"io"
	"os"

	// "github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetLevel(5)
	l.SetOutput(io.MultiWriter(os.Stdout))

	// if len(strings.TrimSpace(conf.SentryDSN)) > 0 {
	// 	hook, err := logrus_sentry.NewSentryHook(conf.SentryDSN, []logrus.Level{
	// 		logrus.PanicLevel,
	// 		logrus.FatalLevel,
	// 		logrus.ErrorLevel,
	// 	})
	//
	// 	if err == nil {
	// 		hook.SetEnvironment(conf.Environment)
	// 		hook.SetRelease(conf.Release)
	// 		hook.StacktraceConfiguration.Enable = true
	// 		hook.Timeout = 0
	// 		l.Hooks.Add(hook)
	// 	} else {
	// 		l.Error(err)
	// 	}
	// }

	return l
}
