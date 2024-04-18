package partialsigmessage

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
)

// PartialRootValid tests PostConsensusMessage root == 32 bytes
func PartialRootValid() *MsgSpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight)

	return &MsgSpecTest{
		Name: "partial root valid",
		Messages: []*types.SignedPartialSignatureMessage{
			msg,
		},
	}
}
