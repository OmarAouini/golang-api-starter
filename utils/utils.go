package utils

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func PrettyPrintDebug(message, i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	logrus.Debugf("%s: %s", message, string(s))
}
