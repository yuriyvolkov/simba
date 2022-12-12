package main

import (
	"net"

	"github.com/yuriyvolkov/simba/pkg/smb"
)

// sendErrorResponse sends an error response to an SMB client
func sendErrorResponse(conn net.Conn, header smb.Header, status smb.Status) error {
	// Create the error response packet
	packet := &smb.Packet{
		Header: header,
		Data: smb.ErrorResponse{
			ErrorClass:    smb.ErrorClassDOS,
			Reserved:      0,
			ErrorCode:     uint32(status),
			ErrorReserved: 0,
		},
	}

	// Marshal the packet
	data, err := packet.Marshal()
	if err != nil {
		return err
	}

	// Send the response
	_, err = conn.Write(data)
	return err
}
