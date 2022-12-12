package smb

import (
	"bytes"
	"encoding/binary"
)

// Header represents an SMB header
type Header struct {
	ProtocolID  [4]byte
	Command     Command
	Status      Status
	Flags       uint32
	Flags2      uint32
	Reserved    uint32
	TreeID      uint32
	ProcessID   uint16
	UserID      uint16
	MultiplexID uint16
	Data        []byte
}

// Marshal serializes an SMB header into a byte slice
func (h *Header) Marshal() ([]byte, error) {
	// Create a new bytes buffer
	buf := new(bytes.Buffer)

	// Write the protocol ID
	if err := binary.Write(buf, binary.LittleEndian, h.ProtocolID); err != nil {
		return nil, err
	}

	// Write the command
	if err := binary.Write(buf, binary.LittleEndian, h.Command); err != nil {
		return nil, err
	}

	// Write the status
	if err := binary.Write(buf, binary.LittleEndian, h.Status); err != nil {
		return nil, err
	}

	// Write the flags
	if err := binary.Write(buf, binary.LittleEndian, h.Flags); err != nil {
		return nil, err
	}

	// Write the flags 2
	if err := binary.Write(buf, binary.LittleEndian, h.Flags2); err != nil {
		return nil, err
	}

	// Write the reserved field
	if err := binary.Write(buf, binary.LittleEndian, h.Reserved); err != nil {
		return nil, err
	}

	// Write the tree ID
	// TODO

	return buf.Bytes(), nil
}

// HeaderParse parses an SMB packet header from a byte slice
func HeaderParse(data []byte) (*Header, error) {
	// Create a new bytes reader
	r := bytes.NewReader(data)

	// Read the protocol ID
	var protocolID [4]byte
	if err := binary.Read(r, binary.LittleEndian, &protocolID); err != nil {
		return nil, err
	}

	// Read the command
	var command Command
	if err := binary.Read(r, binary.LittleEndian, &command); err != nil {
		return nil, err
	}

	// Read the status
	var status Status
	if err := binary.Read(r, binary.LittleEndian, &status); err != nil {
		return nil, err
	}

	// Read the flags
	var flags uint32
	if err := binary.Read(r, binary.LittleEndian, &flags); err != nil {
		return nil, err
	}

	// Read the flags 2
	var flags2 uint32
	if err := binary.Read(r, binary.LittleEndian, &flags2); err != nil {
		return nil, err
	}

	// Read the reserved field
	var reserved uint32
	if err := binary.Read(r, binary.LittleEndian, &reserved); err != nil {
		return nil, err
	}

	// Read the tree ID
	var treeID uint32
	if err := binary.Read(r, binary.LittleEndian, &treeID); err != nil {
		return nil, err
	}

	// Read the process ID
	var processID uint16
	if err := binary.Read(r, binary.LittleEndian, &processID); err != nil {
		return nil, err
	}

	// Read the user ID
	var userID uint16
	if err := binary.Read(r, binary.LittleEndian, &userID); err != nil {
		return nil, err
	}

	// Read the multiplex ID
	var multiplexID uint16
	if err := binary.Read(r, binary.LittleEndian, &multiplexID); err != nil {
		return nil, err
	}

	// Create the header
	header := &Header{
		ProtocolID:  protocolID,
		Command:     command,
		Status:      status,
		Flags:       flags,
		Flags2:      flags2,
		Reserved:    reserved,
		TreeID:      treeID,
		ProcessID:   processID,
		UserID:      userID,
		MultiplexID: multiplexID,
	}

	return header, nil
}
