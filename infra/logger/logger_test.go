package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	f, err := os.OpenFile("../../logs/log.log", os.O_RDONLY|os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	ok := New("servoce", "testing")
	ok.SetFormat(&logrus.JSONFormatter{})
	ok.SetOutputs(f)
	ok.WithParam("with", "param").Info()
	// With(logrus.Fields{"hehe": "hehe"})
}

func TestLogExisting(t *testing.T) {
	WithParams(Params{"existing": "test"}).Info("test")
	// With(logrus.Fields{"hehe": "hehe"})
}
