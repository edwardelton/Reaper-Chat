package utils

import "net"

func RemoveElement(clientConn *[]*net.Conn, index int) {
	(*clientConn)[index] = (*clientConn)[len((*clientConn))-1]
	(*clientConn) = (*clientConn)[:len((*clientConn))-1]
}

func RemoveConnection(conn *net.Conn, clientConn *[]*net.Conn) {
	for index, client := range *clientConn {
		if conn == client {
			RemoveElement(clientConn, index)
		}
	}
}
