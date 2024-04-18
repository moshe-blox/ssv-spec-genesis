package messages

import (
	"github.com/moshe-blox/ssv-spec/qbft"
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// MsgNonZeroIdentifier tests Message with len(Identifier) == 0
func MsgNonZeroIdentifier() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	msg := testingutils.SignQBFTMsg(ks.Shares[1], types.OperatorID(1), &qbft.Message{
		MsgType:    qbft.CommitMsgType,
		Height:     qbft.FirstHeight,
		Round:      qbft.FirstRound,
		Identifier: []byte{},
		Root:       testingutils.TestingQBFTRootData,
	})

	return &tests.MsgSpecTest{
		Name: "msg identifier len == 0",
		Messages: []*qbft.SignedMessage{
			msg,
		},
		ExpectedError: "message identifier is invalid",
	}
}
