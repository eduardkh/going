// https://stackoverflow.com/questions/71573526/send-log-to-rsyslog-server-with-go
// linux only
package main

import (
	"log"
	"log/syslog"
)

func main() {
	s, err := syslog.Dial("udp", "localhost:514", syslog.LOG_INFO, "TAG")
	if err != nil {
		log.Fatal(err)
	}

	s.Info("Just some info")
}
