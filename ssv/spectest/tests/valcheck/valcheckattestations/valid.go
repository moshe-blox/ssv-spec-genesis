package valcheckattestations

import (
	"github.com/bloxapp/ssv-spec-genesis/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec-genesis/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
)

// Valid tests valid data
func Valid() tests.SpecTest {
	return &valcheck.SpecTest{
		Name:       "attestation value check valid",
		Network:    types.PraterNetwork,
		BeaconRole: types.BNRoleAttester,
		Input:      testingutils.TestAttesterConsensusDataByts,
	}
}
