package messages

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft"
	"github.com/bloxapp/ssv-spec-genesis/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
)

// InvalidPrepareJustificationsUnmarshalling tests unmarshalling invalid prepare justifications (during message.validate())
func InvalidPrepareJustificationsUnmarshalling() tests.SpecTest {

	ks := testingutils.Testing4SharesSet()

	msg := testingutils.SignQBFTMsg(ks.Shares[1], types.OperatorID(1), &qbft.Message{
		MsgType:              qbft.ProposalMsgType,
		Height:               qbft.FirstHeight,
		Round:                qbft.FirstRound,
		Identifier:           []byte{1, 2, 3, 4},
		Root:                 testingutils.DifferentRoot,
		PrepareJustification: [][]byte{{1}},
	})

	msg.FullData = testingutils.TestingQBFTFullData

	return &tests.MsgSpecTest{
		Name: "invalid prepare justification unmarshalling",
		Messages: []*qbft.SignedMessage{
			msg,
		},
		ExpectedError: "incorrect size",
	}
}
