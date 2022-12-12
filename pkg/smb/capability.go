package smb

// Capability represents the capabilities of an SMB packet
type Capability uint32

const (
	// CapabilityNTStatus indicates support for NT error codes
	CapabilityNTStatus Capability = 0x00000001

	// CapabilityRPCRemoteAPIs indicates support for remote RPC
	CapabilityRPCRemoteAPIs Capability = 0x00000002

	// CapabilityUnicode indicates support for Unicode strings
	CapabilityUnicode Capability = 0x00000004

	// CapabilityLargeFiles indicates support for large files
	CapabilityLargeFiles Capability = 0x00000008

	// CapabilityExtendedSecurity indicates support for extended security
	CapabilityExtendedSecurity Capability = 0x80000000
)
