package roundchange

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft"
	"github.com/bloxapp/ssv-spec-genesis/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
)

// F1SpeedupPrevPrepared tests catching up to higher rounds via f+1 speedup, other peers are all at the same round (one prev prepared)
func F1SpeedupPrevPrepared() tests.SpecTest {
	pre := testingutils.BaseInstance()

	ks := testingutils.Testing4SharesSet()
	prepareMsgs := []*qbft.SignedMessage{
		testingutils.TestingPrepareMessage(ks.Shares[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.Shares[2], types.OperatorID(2)),
		testingutils.TestingPrepareMessage(ks.Shares[3], types.OperatorID(3)),
	}
	msgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[2], types.OperatorID(2), 10),
		testingutils.TestingRoundChangeMessageWithRoundAndRC(ks.Shares[3], types.OperatorID(3), 10,
			testingutils.MarshalJustifications(prepareMsgs)),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "f+1 speed up prev prepared",
		Pre:           pre,
		PostRoot:      "50f32d85be78af424db9a42c06e43efa15c49dedd0b95b9a1348a2b3346674da",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingRoundChangeMessageWithParams(ks.Shares[1], types.OperatorID(1), 10, qbft.FirstHeight,
				[32]byte{}, 0, [][]byte{}),
		},
	}
}
