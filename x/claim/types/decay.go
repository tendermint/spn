package types

import (
	"fmt"
)

// Validate validates the decay information
func (m DecayInformation) Validate() error {
	if m.Enabled && m.DecayStart.After(m.DecayEnd) {
		return fmt.Errorf("decay starts after decay end %s > %s", m.DecayStart.String(), m.DecayEnd.String())
	}

	return nil
}
