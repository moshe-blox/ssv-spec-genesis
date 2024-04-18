package testingutils

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft"
)

func UnknownDutyValueCheck() qbft.ProposedValueCheckF {
	return func(data []byte) error {
		return nil
	}
}
