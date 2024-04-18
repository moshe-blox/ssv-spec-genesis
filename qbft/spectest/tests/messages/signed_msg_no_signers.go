package messages

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft"
	"github.com/bloxapp/ssv-spec-genesis/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
)

// SignedMsgNoSigners tests SignedMessage len(signers) == 0
func SignedMsgNoSigners() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	msg := testingutils.TestingCommitMessage(ks.Shares[1], types.OperatorID(1))
	msg.Signers = nil

	return &tests.MsgSpecTest{
		Name: "no signers",
		Messages: []*qbft.SignedMessage{
			msg,
		},
		ExpectedError: "message signers is empty",
	}
}
