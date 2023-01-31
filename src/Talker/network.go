package main

import (
	"Talker/utils"
	"fmt"
	"net"
)

func acceptConnection(listener *net.Listener) {
	for {
		conn, _ := (*listener).Accept()
		utils.Debug("Accepted a new connection.")
		go handleLogin(&conn)
	}
}

func broadcast(msg string, conn *net.Conn) {
	utils.Debug("Broadcasting the message.")

	for index, client := range clientConn {
		utils.Debug("Looping through connection.")

		if client != conn {
			_, err := (*client).Write([]byte(fmt.Sprintf("%s\r\n", msg)))

			if err != nil {
				utils.Debug("[Error] Sending: ", err.Error())
				utils.RemoveElement(&clientConn, index)
			}
		}

	}
}
