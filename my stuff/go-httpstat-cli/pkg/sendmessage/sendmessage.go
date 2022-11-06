package sendmessage

import (
	"fmt"
	"net"
)

func SendMessage(proto, address, port, message string) {
	if proto == "" {
		proto = "tcp"
	}
	if address == "" {
		address = "127.0.0.1"
	}
	if port == "" {
		port = "514"
	}
	if message == "" {
		message = "default message"
	}

	connection, err := net.Dial(proto, address+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(connection, message)

}
