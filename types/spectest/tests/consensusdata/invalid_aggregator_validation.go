package consensusdata

import "github.com/bloxapp/ssv-spec-genesis/types/testingutils"

// InvalidAggregatorValidation tests an invalid consensus data with AggregateAndProof
func InvalidAggregatorValidation() *ConsensusDataTest {

	ks := testingutils.Testing4SharesSet()

	cd := testingutils.TestAggregatorWithJustificationsConsensusData(ks)

	cd.DataSSZ = testingutils.TestingSyncCommitteeBlockRoot[:]

	return &ConsensusDataTest{
		Name:          "invalid aggregator data",
		ConsensusData: *cd,
		ExpectedError: "could not unmarshal ssz: incorrect size",
	}
}
