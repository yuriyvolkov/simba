package smb

import (
	"net"
	"time"
)

// handleNegotiateCommand handles an SMB negotiate request
func handleNegotiateCommand(conn net.Conn, packet *Packet) error {
	data, err := packet.Data.Marshal()
	if err != nil {
		return err
	}

	// Parse the request data
	request, err := NegotiateRequestParse(data)
	if err != nil {
		return err
	}

	// Check if the request contains any supported dialects
	var dialect Dialect = DialectUnknown
	for _, d := range request.Dialects {
		if isDialectSupported(d) {
			dialect = d
			break
		}
	}

	// If no supported dialects were found, return an error
	if dialect == DialectUnknown {
		return sendErrorResponse(conn, StatusNotSupported)
	}

	// Create the response
	response := &NegotiateResponse{
		Dialect: dialect,
		SecurityMode: SecurityModeUserLevel | SecurityModeEncryptPasswords |
			SecurityModeSignaturesEnabled | SecurityModeSignaturesRequired,
		MaxBufferSize: 64 * 1024,
		MaxMpxCount:   2,
		Capabilities: CapabilityUnicode | CapabilityNTStatus |
			CapabilityLargeFiles | CapabilityInfoLevelPassthru,
		SystemTime:          time.Now(),
		ServerTimeZone:      -60,
		Challenge:           randomBytes(8),
		DomainName:          "EXAMPLE",
		ServerName:          "FILESERVER",
		EncryptionKeyLength: 8,
	}

	// Marshal the response
	data, err = response.Marshal()
	if err != nil {
		return err
	}

	// Send the response
	return sendResponse(conn, packet, data)
}

func sendResponse(conn net.Conn, packet *Packet, data []byte) error {
	// Set the data field of the packet's header
	packet.Header.Data = data

	// Marshal the packet
	buf, err := packet.Marshal()
	if err != nil {
		return err
	}

	// Send the packet
	_, err = conn.Write(buf)
	return err
}

// sendErrorResponse sends an error response to the client
func sendErrorResponse(conn net.Conn, status Status) error {
	// Create the response packet
	response := &Packet{
		Header: Header{
			ProtocolID: [4]byte{'N', 'B', 'L', 'S'},
			Command:    CommandNegotiate,
			Status:     status,
		},
	}

	// Marshal the packet
	data, err := response.Marshal()
	if err != nil {
		return err
	}

	// Send the response
	if _, err := conn.Write(data); err != nil {
		return err
	}

	return nil
}
