package utils

import (
	"Talker/global"
	"fmt"
	"net"
)

func Debug(text ...string) {
	if global.DEBUG_TYPE {
		for _, str := range text {
			fmt.Print(str)
		}
	}
}

func WriteBanner(conn *net.Conn) {
	(*conn).Write([]byte(fmt.Sprintf("%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n",
		"ooooooooo.                                                       .oooooo.   oooo                      .  ",
		"`888   `Y88.                                                    d8P'  `Y8b  `888                    .o8  ",
		" 888   .d88'  .ooooo.   .oooo.   oo.ooooo.   .ooooo.  oooo d8b 888           888 .oo.    .oooo.   .o888oo ",
		" 888ooo88P'  d88' `88b `P  )88b   888' `88b d88' `88b `888\"\"8P 888           888P\"Y88b  `P  )88b    888  ",
		" 888`88b.    888ooo888  .oP\"888   888   888 888ooo888  888     888           888   888   .oP\"888    888  ",
		" 888  `88b.  888    .o d8(  888   888   888 888    .o  888     `88b    ooo   888   888  d8(  888    888 .",
		"o888o  o888o `Y8bod8P' `Y888\"\"8o  888bod8P' `Y8bod8P' d888b     `Y8bood8P'  o888o o888o `Y888\"\"8o   \"888\"",
		"                                  888                                                                     ",
		"                                 o888o                                                                    ",
		" 																										   ")))
}
