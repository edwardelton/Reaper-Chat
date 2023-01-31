package main

import (
	"TalkerClient/global"
	"TalkerClient/utils"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func handleLogin(conn *net.Conn) {
	key, err := bufio.NewReader(*conn).ReadString('\n')
	fmt.Print(key)

	utils.HandleError(err, "Exiting the program. The server maybe down.")

	msg, err := bufio.NewReader(os.Stdin).ReadString('\n')

	utils.HandleError(err, "Exiting the program. The server maybe down.")

	(*conn).Write([]byte(msg))

	response, err := bufio.NewReader(*conn).ReadString('\n')

	utils.HandleError(err, "Exiting the program. The server maybe down.")

	if strings.Contains(response, "Error") {
		handleLogin(conn)
	}
}

func readChat(conn *net.Conn) {
	defer (*conn).Close()

	for {
		msg, err := bufio.NewReader(*conn).ReadString('\n')

		utils.HandleError(err, "Exiting the program. The server maybe down.")

		msg = DecryptText([]byte(msg[:len(msg)-2]))
		fmt.Print(msg)
		fmt.Print("-> ")
	}
}

func writeChat(conn *net.Conn) {
	defer (*conn).Close()

	for {
		fmt.Print("-> ")
		msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		if strings.Compare(msg, "/exit\r\n") == 0 {
			fmt.Println("Closing the client.")
			time.Sleep(5 * time.Second)
			os.Exit(0)
		} else if strings.Compare(msg, "\r\n") == 0 || strings.Compare(msg, " \r\n") == 0 {
			continue
		} else if strings.Compare(msg, "/clear\r\n") == 0 {
			utils.ClearScreen()
		} else {
			msg = fmt.Sprintf("%s | %s", global.Username, msg)
			encryptedMessage := EncryptText([]byte(msg))
			returnString := "\r\n"
			(*conn).Write(append(encryptedMessage, returnString...))
		}
	}
}
