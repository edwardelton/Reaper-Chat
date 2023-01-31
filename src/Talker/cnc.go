package main

import (
	"Talker/global"
	"Talker/utils"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var clientConn []*net.Conn

func main() {
	listenerServer, err := net.Listen(global.CONN_TYPE, fmt.Sprintf("%s:%d", global.CONN_HOST, global.CONN_PORT))

	if err != nil {
		utils.Debug("[Error] Listening:", err.Error())
		os.Exit(1)
	}
	defer listenerServer.Close()

	acceptConnection(&listenerServer)
}

func handleLogin(conn *net.Conn) {
	for {
		(*conn).Write([]byte("Key: \r\n"))
		response, err := bufio.NewReader(*conn).ReadString('\n')

		if err != nil {
			utils.Debug("[Error] Receiving:", err.Error())
			return
		}

		if response == "1234\r\n" {
			utils.WriteBanner(conn)
			clientConn = append(clientConn, conn)
			handleCommand(conn)
		} else {
			(*conn).Write([]byte("[Error] Wrong Key.\r\n"))
		}
	}
}

func handleCommand(conn *net.Conn) {
	defer (*conn).Close()

	for {
		msg, err := bufio.NewReader(*conn).ReadString('\n')

		if err != nil {
			utils.Debug("[Error] Receiving:", err.Error())
			return
		}

		utils.Debug("Received a message.", msg)

		if strings.Compare(msg, "exit\r\n") == 0 {
			utils.RemoveConnection(conn, &clientConn)
			(*conn).Close()
			return
		} else {
			go broadcast(msg, conn)
		}
	}
}
