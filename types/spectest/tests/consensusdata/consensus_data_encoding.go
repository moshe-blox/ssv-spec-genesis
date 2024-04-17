package consensusdata

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// ConsensusDataEncoding tests encoding and decoding ConsensusData for all duties
func ConsensusDataEncoding(name string, cd *types.ConsensusData) *EncodingTest {

	byts, err := cd.Encode()
	if err != nil {
		panic(err.Error())
	}
	root, err := cd.HashTreeRoot()
	if err != nil {
		panic(err.Error())
	}

	return &EncodingTest{
		Name:         name,
		Data:         byts,
		ExpectedRoot: root,
	}
}

func ProposerConsensusDataEncoding() *EncodingTest {
	return ConsensusDataEncoding("proposer encoding", testingutils.TestProposerBlindedBlockConsensusDataV(spec.DataVersionCapella))
}
func BlindedProposerConsensusDataEncoding() *EncodingTest {
	return ConsensusDataEncoding("blinded proposer encoding", testingutils.TestProposerBlindedBlockConsensusDataV(spec.DataVersionCapella))
}
func AttestationConsensusDataEncoding() *EncodingTest {
	return ConsensusDataEncoding("attestation encoding", testingutils.TestAttesterConsensusData)
}
func AggregatorConsensusDataEncoding() *EncodingTest {
	return ConsensusDataEncoding("aggregation encoding", testingutils.TestAggregatorConsensusData)
}
func SyncCommitteeConsensusDataEncoding() *EncodingTest {
	return ConsensusDataEncoding("sync committee encoding", testingutils.TestSyncCommitteeConsensusData)
}
func SyncCommitteeContributionConsensusDataEncoding() *EncodingTest {
	return ConsensusDataEncoding("sync committee contribution encoding", testingutils.TestSyncCommitteeContributionConsensusData)
}
