package main

import (
	"TalkerClient/global"
	"TalkerClient/utils"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial(global.CONN_TYPE, fmt.Sprintf("%s:%d", global.CONN_HOST, global.CONN_PORT))

	if err != nil {
		utils.Debug("[Error] Listening:", err.Error(), "\nExiting the program.")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	defer conn.Close()

	utils.GetUsername()

	handleLogin(&conn)
	go readChat(&conn)
	writeChat(&conn)
}
