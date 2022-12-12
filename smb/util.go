package smb

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"unicode/utf16"
)

func randomBytes(n int) []byte {
	buf := make([]byte, n)
	rand.Read(buf)
	return buf
}

func readSMBString(r io.Reader, isUnicode bool) (string, error) {
	// Read the string length
	var length uint8
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return "", err
	}

	// Read the string data
	data := make([]byte, length)
	if _, err := r.Read(data); err != nil {
		return "", err
	}

	if isUnicode {
		// If the string is in Unicode format, convert it to a Go string
		return string(utf16.Decode(data)), nil
	}

	// Otherwise, return the raw string data
	return string(data), nil
}
