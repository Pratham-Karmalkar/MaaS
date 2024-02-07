package models

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

//Log Struct can be customized however we like
type Log struct {
	IP          string
	UserID      int
	HTTPVersion string
	HTTPMethod  string
	HTTPStatus  int
	URL         string
	TimeStamp   string
}

//Generates Data to Log struct using gofakeit library
func (log *Log) GenerateLog(flag bool) Log {

	if flag {
		log.IP = gofakeit.IPv4Address()
		log.UserID = gofakeit.Number(1, 100301)
		log.HTTPVersion = gofakeit.HTTPVersion()
		log.HTTPMethod = gofakeit.HTTPMethod()
		log.HTTPStatus = gofakeit.HTTPStatusCode()
		log.URL = gofakeit.URL()
		log.TimeStamp = time.Now().Local().Format("2006/01/02 03:04:05") //time format YY/MM/DD HH:MM:SS
	}
	return *log
}

func (log *Log) StopGeneration() bool {

	return true
}

//  # paths:
//   #   - /var/log/*.log
