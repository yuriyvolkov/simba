package smb

import (
	"bytes"
	"encoding/binary"
)

// TreeConnectRequest represents an SMB tree connect request
// The TreeConnectRequest struct has three fields:

//     Path: The path of the share to connect to.
//     Password: The password for the share, if any.
//     Service: The service associated with the share.
type TreeConnectRequest struct {
	Path     string
	Password string
	Service  string
}

// TreeConnectRequestParse parses an SMB tree connect request
func TreeConnectRequestParse(data []byte) (*TreeConnectRequest, error) {
	// Create a new bytes reader
	r := bytes.NewReader(data)

	// Read the word count
	var wordCount uint8
	if err := binary.Read(r, binary.LittleEndian, &wordCount); err != nil {
		return nil, err
	}

	// Read the flags
	var flags uint16
	if err := binary.Read(r, binary.LittleEndian, &flags); err != nil {
		return nil, err
	}

	// Read the path
	path, err := readSMBString(r, flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}

	// Read the password
	password, err := readSMBString(r, flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}

	// Read the service
	service, err := readSMBString(r, flags&FlagUnicode != 0)
	if err != nil {
		return nil, err
	}

	// Create the request
	request := &TreeConnectRequest{
		Path:     path,
		Password: password,
		Service:  service,
	}

	return request, nil
}
