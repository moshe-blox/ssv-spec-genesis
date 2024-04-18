package partialsigmessage

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
)

// SigValid tests SignedPostConsensusMessage sig == 96 bytes
func SigValid() *MsgSpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight)

	return &MsgSpecTest{
		Name: "sig valid",
		Messages: []*types.SignedPartialSignatureMessage{
			msg,
		},
	}
}
