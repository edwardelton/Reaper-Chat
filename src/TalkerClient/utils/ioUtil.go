package utils

import (
	"TalkerClient/global"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Debug(text ...string) {
	if global.DEBUG_TYPE {
		for _, str := range text {
			fmt.Print(str)
		}
	}
}

func ClearScreen() {
	const OS = runtime.GOOS
	var cmd *exec.Cmd

	if strings.Compare(OS, "linux") == 0 {
		cmd = exec.Command("clear")
	} else if strings.Compare(OS, "windows") == 0 {
		fmt.Print("here")
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetUsername() {
	fmt.Print("What's the username you want to display: ")
	username, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		Debug("[Error] Reading:", err.Error())
		GetUsername()
	}
	global.Username = username[:len(username)-2]
}
