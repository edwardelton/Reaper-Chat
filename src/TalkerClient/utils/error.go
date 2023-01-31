package utils

import (
	"fmt"
	"os"
	"time"
)

func HandleError(err error, text string) {
	if err != nil {
		Debug("[Error] Receiving:", err.Error())
		fmt.Println(text)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}
