package messages

import (
	"github.com/moshe-blox/ssv-spec/qbft"
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
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
