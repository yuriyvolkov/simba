package smb

// SecurityMode represents the security mode of an SMB packet
type SecurityMode uint8

const (
	// SecurityModeUserLevel indicates user-level security
	SecurityModeUserLevel SecurityMode = 0x01

	// SecurityModeEncryptPasswords indicates encrypted password support
	SecurityModeEncryptPasswords SecurityMode = 0x02

	// SecurityModeSignaturesEnabled indicates signature support
	SecurityModeSignaturesEnabled SecurityMode = 0x04

	// SecurityModeSignaturesRequired indicates signature support
	SecurityModeSignaturesRequired SecurityMode = 0x08
)
