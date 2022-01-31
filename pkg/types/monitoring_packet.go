package types

import (
	"fmt"
	"github.com/pkg/errors"
)

func (m MonitoringPacket) ValidateBasic() error {
	if err := m.SignatureCounts.Validate(); err != nil {
		return errors.Wrap(err, "invalid signature counts")
	}

	if m.BlockHeight < m.SignatureCounts.BlockCount {
		return fmt.Errorf(
			"block height %d must be greater or equal to block count %d",
			m.BlockHeight,
			m.SignatureCounts.BlockCount,
		)
	}
	return nil
}
