package modul7_logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func Test_log1(t *testing.T) {
	logger := logrus.New()

	logger.Println("hi")
}

func Test_log2(t *testing.T) {
	//pake level
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logger.Debug("hahha")
	logger.Warnln("hahha")
}

func Test_log3(t *testing.T) {
	//pake file
	logger := logrus.New()

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {

	}
	logger.SetOutput(file)

	logger.Infoln("hahha pake file1")
}

func Test_log_json(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Infoln("pake json")
}

func Test_log_custem_field(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "akulah").Infoln("login")

	logger.WithField("username1", "akulah1").
		WithField("jam", "sekian").
		Infoln("login")

	logger.WithFields(logrus.Fields{
		"data1": "aku",
		"data2": "akulagi",
	}).Infoln("hahha")

}

func Test_hook(t *testing.T) {
	//callback saat ada log level tertentu
	logger := logrus.New()

	logger.AddHook(&SampleHook{})

	logger.Infoln("msg dari info")
	logger.Warnln("msg dari warn")

}

func Test_struct(t *testing.T) {

	//logger := logrus.New()
	//
	//file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//
	//}
	//logger.SetOutput(file)
	//
	//logger.Infoln(respJson)
}
