package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type CustomerFormatter struct {
	Prefix string
}

func (s *CustomerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("[%s] %s [%s] %s\n", s.Prefix, timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}
