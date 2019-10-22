package main

import (
	"fmt"

	"github.com/jeromer/syslogparser"
)

type LogHole struct {
}

func Logswitcher(logparts syslogparser.LogParts) {
	fmt.Printf("%s %s %s %s %s\n", ts(logparts["timestamp"]), logparts["severity"], logparts["hostname"], logparts["tag"], logparts["content"])

}
