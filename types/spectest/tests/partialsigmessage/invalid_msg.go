package partialsigmessage

import (
	"github.com/moshe-blox/ssv-spec/qbft"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// InvalidMsg tests a signed msg with 1 invalid message
func InvalidMsg() *MsgSpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.PostConsensusAttestationMsg(ks.Shares[1], 1, qbft.FirstHeight)
	msg.Message.Messages = append(msg.Message.Messages, &types.PartialSignatureMessage{})

	return &MsgSpecTest{
		Name:          "invalid message",
		Messages:      []*types.SignedPartialSignatureMessage{msg},
		ExpectedError: "inconsistent signers",
	}
}
