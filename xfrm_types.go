package netlink

// Proto is an enum representing an ipsec protocol.
type Proto uint8

// Mode is an enum representing an ipsec transport.
type Mode uint8

// XfrmMark represents the mark associated to the state or policy
type XfrmMark struct {
	Value uint32
	Mask  uint32
}
