package smb

import (
	"errors"
	"unsafe"
)

// FlagUnicode indicates that the data in the SMB packet is encoded in Unicode.
const FlagUnicode uint16 = 0x8000

type Marshaller interface {
	Marshal() ([]byte, error)
}

// Packet represents an SMB packet
type Packet struct {
	Header Header
	Data   Marshaller
}

// Marshal serializes an SMB packet into a byte slice
func (p *Packet) Marshal() ([]byte, error) {
	// Serialize the packet data
	data, err := p.Data.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the header
	header, err := p.Header.Marshal()
	if err != nil {
		return nil, err
	}

	// Concatenate the header and data
	return append(header, data...), nil
}

// PacketParse parses an SMB packet from a byte slice
func PacketParse(data []byte) (*Packet, error) {
	// Parse the header
	header, err := HeaderParse(data)
	if err != nil {
		return nil, err
	}

	// Get the data section of the packet
	headerSize := unsafe.Sizeof(Header{})
	packetData := data[headerSize:]

	// Parse the packet data
	parsedData, err := ParseData(header.Command, packetData)
	if err != nil {
		return nil, err
	}

	// Create the packet
	packet := &Packet{
		Header: *header,
		Data:   parsedData,
	}

	return packet, nil
}

// ParseData parses the data section of an SMB packet
func ParseData(command Command, data []byte) (Marshaller, error) {
	switch command {
	case CommandNegotiate:
		return NegotiateRequestParse(data)
	case CommandSessionSetup:
		return SessionSetupRequestParse(data)
	case CommandTreeConnect:
		return TreeConnectRequestParse(data)
	case CommandTreeDisconnect:
		return TreeDisconnectRequestParse(data)
	// case CommandCreate:
	// 	return CreateRequestParse(data)
	// case CommandClose:
	// 	return CloseRequestParse(data)
	// case CommandFlush:
	// 	return FlushRequestParse(data)
	// case CommandRead:
	// 	return ReadRequestParse(data)
	// case CommandWrite:
	// 	return WriteRequestParse(data)
	// case CommandDelete:
	// 	return DeleteRequestParse(data)
	// case CommandRename:
	// 	return RenameRequestParse(data)
	// case CommandQueryInfo:
	// 	return QueryInfoRequestParse(data)
	// case CommandSetInfo:
	// 	return SetInfoRequestParse(data)
	default:
		return nil, errors.New("invalid command")
	}
}
