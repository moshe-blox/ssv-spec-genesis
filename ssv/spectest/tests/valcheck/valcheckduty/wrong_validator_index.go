package valcheckduty

import (
	"encoding/json"

	"github.com/attestantio/go-eth2-client/spec"

	"github.com/bloxapp/ssv-spec-genesis/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec-genesis/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
)

// WrongValidatorIndex tests duty.ValidatorIndex wrong
func WrongValidatorIndex() tests.SpecTest {
	consensusDataBytsF := func(cd *types.ConsensusData) []byte {
		cdCopy := &types.ConsensusData{}
		b, _ := json.Marshal(cd)
		if err := json.Unmarshal(b, cdCopy); err != nil {
			panic(err.Error())
		}
		cdCopy.Duty.ValidatorIndex = testingutils.TestingValidatorIndex + 100

		ret, _ := cdCopy.Encode()
		return ret
	}

	expectedErr := "duty invalid: wrong validator index"
	return &valcheck.MultiSpecTest{
		Name: "wrong validator index",
		Tests: []*valcheck.SpecTest{
			{
				Name:          "sync committee aggregator",
				Network:       types.BeaconTestNetwork,
				BeaconRole:    types.BNRoleSyncCommitteeContribution,
				Input:         consensusDataBytsF(testingutils.TestSyncCommitteeContributionConsensusData),
				ExpectedError: expectedErr,
			},
			{
				Name:          "sync committee",
				Network:       types.BeaconTestNetwork,
				BeaconRole:    types.BNRoleSyncCommittee,
				Input:         consensusDataBytsF(testingutils.TestSyncCommitteeConsensusData),
				ExpectedError: expectedErr,
			},
			{
				Name:          "aggregator",
				Network:       types.BeaconTestNetwork,
				BeaconRole:    types.BNRoleAggregator,
				Input:         consensusDataBytsF(testingutils.TestAggregatorConsensusData),
				ExpectedError: expectedErr,
			},
			{
				Name:          "proposer",
				Network:       types.BeaconTestNetwork,
				BeaconRole:    types.BNRoleProposer,
				Input:         consensusDataBytsF(testingutils.TestProposerConsensusDataV(spec.DataVersionDeneb)),
				ExpectedError: expectedErr,
			},
			{
				Name:          "attester",
				Network:       types.BeaconTestNetwork,
				BeaconRole:    types.BNRoleAttester,
				Input:         consensusDataBytsF(testingutils.TestAttesterConsensusData),
				ExpectedError: "duty invalid: wrong validator index",
			},
		},
	}
}
