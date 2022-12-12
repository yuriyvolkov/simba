package smb

import (
	"encoding/binary"
	"io"
)

// Dialect represents an SMB dialect
type Dialect uint16

const (
	// DialectUnknown indicates an unknown dialect
	DialectUnknown Dialect = 0

	// DialectSMB202 indicates the SMB 2.0.2 dialect
	DialectSMB202 Dialect = 0x0202

	// DialectSMB210 indicates the SMB 2.1 dialect
	DialectSMB210 Dialect = 0x0210

	// DialectSMB300 indicates the SMB 3.0 dialect
	DialectSMB300 Dialect = 0x0300

	// DialectSMB302 indicates the SMB 3.0.2 dialect
	DialectSMB302 Dialect = 0x0302

	// DialectSMB311 indicates the SMB 3.1.1 dialect
	DialectSMB311 Dialect = 0x0311
)

// Dialects represents a list of SMB dialects
type Dialects []Dialect

// Read implements the io.Reader interface
func (d *Dialects) Read(r io.Reader) (int, error) {
	// Read the number of dialects
	var numDialects uint16
	if err := binary.Read(r, binary.LittleEndian, &numDialects); err != nil {
		return 0, err
	}

	// Read the dialects
	*d = make([]Dialect, numDialects)
	buf := make([]byte, numDialects*2) // each Dialect is a uint16, so each dialect takes 2 bytes
	n, err := r.Read(buf)
	if err != nil {
		return 0, err
	}

	// Convert the byte slice to a Dialects slice
	for i := 0; i < int(numDialects); i++ {
		(*d)[i] = Dialect(binary.LittleEndian.Uint16(buf[i*2:]))
	}

	return n + 2, nil // 2 bytes were read for the numDialects field
}

func isDialectSupported(d Dialect) bool {
	return d == DialectSMB202 || d == DialectSMB210 || d == DialectSMB300 || d == DialectSMB302 || d == DialectSMB311
}
