package smb

import (
	"bytes"
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

// readVarString function takes a bytes.Reader object as its first argument,
// the length of the string in bytes as its second argument, and a boolean indicating
// whether the string is in Unicode format as its third argument. It reads the string
// from the reader, converts it to a Go string, and returns it along with any error that occurred.
// If the string is in Unicode format, it uses the utf16 package to decode the string from a byte
// slice to a Go string. Otherwise, it simply converts the byte slice to a Go string using the string()
// function.
func readVarString(buf *bytes.Reader, n int, unicode bool) (string, error) {
	// Create a byte slice to read the string into
	var str []byte

	// Check if the string is in Unicode format
	if unicode {
		// Read 2 bytes for each character in the string
		str = make([]byte, n*2)
		if err := binary.Read(buf, binary.LittleEndian, &str); err != nil {
			return "", err
		}
		// Convert the byte slice to a Unicode string
		return string(utf16.Decode(str)), nil
	} else {
		// Read 1 byte for each character in the string
		str = make([]byte, n)
		if err := binary.Read(buf, binary.LittleEndian, &str); err != nil {
			return "", err
		}
		// Convert the byte slice to a string
		return string(str), nil
	}
}

// readVarBytes function takes a bytes.Reader object as its only argument and uses it to read the length
// of a byte slice, followed by the byte slice itself. It returns the byte slice along with any error
// that occurred.
func readVarBytes(buf *bytes.Reader) ([]byte, error) {
	// Read the length of the byte slice
	var length uint16
	if err := binary.Read(buf, binary.LittleEndian, &length); err != nil {
		return nil, err
	}

	// Create a byte slice to read the data into
	data := make([]byte, length)

	// Read the data into the byte slice
	if err := binary.Read(buf, binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	// Return the byte slice
	return data, nil
}

// writeVarString function takes a bytes.Buffer and a string as its arguments and writes the string to the buffer in the appropriate format. If the string is in Unicode format (indicated by the presence of the Unicode format marker 0xFEFF at the beginning of the string), the function writes the Unicode format marker to the buffer and then writes the Unicode string in little-endian order. If the string is not in Unicode format, the function simply writes the ASCII string to the buffer. This function does not return any value, as it writes directly to the provided bytes.Buffer.
func writeVarString(buf *bytes.Buffer, s string) {
	// Check if the string is in Unicode format
	isUnicode := false
	if len(s) > 1 && s[0] == 0xFE && s[1] == 0xFF {
		isUnicode = true
	}

	// Write the string to the buffer
	if isUnicode {
		// Write the Unicode format marker
		binary.Write(buf, binary.LittleEndian, uint16(0xFEFF))

		// Write the Unicode string
		for i := 2; i < len(s); i += 2 {
			binary.Write(buf, binary.LittleEndian, uint16(s[i]<<8|s[i+1]))
		}
	} else {
		// Write the ASCII string
		buf.WriteString(s)
	}
}

// writeVarBytes function takes a bytes.Buffer and a byte slice as its arguments and writes the length of the byte slice and the byte slice itself to the buffer. The length of the byte slice is written as a 16-bit integer in little-endian order, followed by the byte slice itself. This function does not return any value, as it writes directly to the provided bytes.Buffer.
func writeVarBytes(buf *bytes.Buffer, b []byte) {
	// Write the length of the byte slice
	binary.Write(buf, binary.LittleEndian, uint16(len(b)))

	// Write the byte slice
	buf.Write(b)
}
