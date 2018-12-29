package dhcpv4

import (
	"fmt"

	"github.com/u-root/u-root/pkg/uio"
)

// This option implements the relay agent information option
// https://tools.ietf.org/html/rfc3046

// OptRelayAgentInformation is a "container" option for specific agent-supplied
// sub-options.
type OptRelayAgentInformation struct {
	Options Options
}

// ParseOptRelayAgentInformation returns a new OptRelayAgentInformation from a
// byte stream, or error if any.
func ParseOptRelayAgentInformation(data []byte) (*OptRelayAgentInformation, error) {
	options, err := OptionsFromBytesWithParser(data, ParseOptionGeneric, false /* don't check for OptionEnd tag */)
	if err != nil {
		return nil, err
	}
	return &OptRelayAgentInformation{Options: options}, nil
}

// Code returns the option code.
func (o *OptRelayAgentInformation) Code() OptionCode {
	return OptionRelayAgentInformation
}

// ToBytes returns a serialized stream of bytes for this option.
func (o *OptRelayAgentInformation) ToBytes() []byte {
	return uio.ToBigEndian(o.Options)
}

// String returns a human-readable string for this option.
func (o *OptRelayAgentInformation) String() string {
	return fmt.Sprintf("Relay Agent Information -> %v", o.Options)
}

// Length returns the length of the data portion (excluding option code and byte
// for length, if any).
func (o *OptRelayAgentInformation) Length() int {
	l := 0
	for _, opt := range o.Options {
		l += 2 + opt.Length()
	}
	return l
}