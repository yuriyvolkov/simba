package smb

import (
	"bytes"
	"encoding/binary"
	"time"
)

// NegotiateResponse represents an SMB negotiate response
type NegotiateResponse struct {
	Dialect             Dialect
	SecurityMode        SecurityMode
	MaxBufferSize       uint32
	MaxMpxCount         uint16
	Capabilities        Capability
	SystemTime          time.Time
	ServerTimeZone      int16
	Challenge           [8]byte
	DomainName          string
	ServerName          string
	EncryptionKeyLength uint16
}

// Marshal serializes an SMB negotiate response into a byte slice
func (r *NegotiateResponse) Marshal() ([]byte, error) {
	// Create a new bytes buffer
	buf := new(bytes.Buffer)

	// Write the word count
	if err := binary.Write(buf, binary.LittleEndian, uint8(17)); err != nil {
		return nil, err
	}

	// Write the dialect
	if err := binary.Write(buf, binary.LittleEndian, r.Dialect); err != nil {
		return nil, err
	}

	// Write the security mode
	if err := binary.Write(buf, binary.LittleEndian, r.SecurityMode); err != nil {
		return nil, err
	}

	// Write the maximum buffer size
	if err := binary.Write(buf, binary.LittleEndian, r.MaxBufferSize); err != nil {
		return nil, err
	}

	// Write the maximum multiplex count
	if err := binary.Write(buf, binary.LittleEndian, r.MaxMpxCount); err != nil {
		return nil, err
	}

	// Write the capabilities
	if err := binary.Write(buf, binary.LittleEndian, r.Capabilities); err != nil {
		return nil, err
	}

	// Write the system time
	if err := binary.Write(buf, binary.LittleEndian, uint64(time.Date(
		r.SystemTime.Year(),
		r.SystemTime.Month(),
		r.SystemTime.Day(),
		r.SystemTime.Hour(),
		r.SystemTime.Minute(),
		r.SystemTime.Second(),
		0,
		time.UTC,
	).Unix())); err != nil {
		return nil, err
	}

	// Write the server time zone
	if err := binary.Write(buf, binary.LittleEndian, r.ServerTimeZone); err != nil {
		return nil, err
	}

	// FIXME add missing Writes

	return buf.Bytes(), nil
}
