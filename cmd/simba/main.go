package main

import (
	"net"

	"github.com/yuriyvolkov/simba/pkg/smb"
)

// shareMap maps share names to directories on the local filesystem
var shareMap = map[string]string{
	"share1": "/path/to/share1",
	"share2": "/path/to/share2",
	// ...
}

func main() {
	// Listen for incoming connections on port 445
	ln, err := net.Listen("tcp", ":445")
	if err != nil {
		// Handle error
	}

	for {
		// Accept incoming connections
		conn, err := ln.Accept()
		if err != nil {
			// Handle error
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

// handleConnection handles an incoming SMB connection
func handleConnection(conn net.Conn) {
	// Close the connection when this function returns
	defer conn.Close()

	// Receive the incoming SMB packet
	data := make([]byte, 4096)
	n, err := conn.Read(data)
	if err != nil {
		// Handle error
	}

	// Parse the SMB packet
	packet, err := smb.PacketParse(data[:n])
	if err != nil {
		// Handle error
	}

	// Handle the SMB packet according to its command
	switch packet.Header.Command {
	case smb.CommandNegotiate:
		handleNegotiateCommand(conn, packet)
	case smb.CommandSessionSetup:
		handleSessionSetupCommand(conn, packet)
	case smb.CommandTreeConnect:
		handleTreeConnectCommand(conn, packet)
	case smb.CommandCreate:
		handleCreateCommand(conn, packet)
	case smb.CommandClose:
		handleCloseCommand(conn, packet)
	case smb.CommandWrite:
		handleWriteCommand(conn, packet)
	case smb.CommandRead:
		handleReadCommand(conn, packet)
	case smb.CommandTreeDisconnect:
		handleTreeDisconnectCommand(conn, packet)
	case smb.CommandLogoff:
		handleLogoffCommand(conn, packet)
	default:
		// Send error response for unknown command
		sendErrorResponse(conn, packet.Header, smb.StatusNotImplemented)
	}
}
