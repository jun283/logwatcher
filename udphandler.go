package main

import (
	"fmt"
	"time"

	"github.com/jeromer/syslogparser"
)

type LogHole struct {
}

func ts(s interface{}) string {

	if s == nil {
		return time.Now().Format(time.RFC3339)
	}

	return s.(time.Time).Local().Format(time.RFC3339)

}

func Logswitcher(logparts syslogparser.LogParts) {
	//fmt.Println(logparts)
	//fmt.Printf("%s %s %s %s %s\n", ts(logparts["timestamp"]), logparts["severity"], logparts["hostname"], logparts["tag"], logparts["content"])
	fmt.Println(ts(logparts["timestamp"]), logparts["source"])
	fmt.Println(logparts["content"])
}
