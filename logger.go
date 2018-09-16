package cloudskine

import (
	"log"
	"errors"
)

type Logger struct {
	logger log.Logger
}

func (l *Logger) Log(msg string, die bool) error {
	var err Error
	if die == true {
		l.logger.Panic(msg)
	} else {
		l.logger.Print(msg)
		err = error.New(msg)
	}
	return err
}
