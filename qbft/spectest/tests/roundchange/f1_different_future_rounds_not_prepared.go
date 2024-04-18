package roundchange

import (
	"github.com/moshe-blox/ssv-spec/qbft"
	"github.com/moshe-blox/ssv-spec/qbft/spectest/tests"
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
	"github.com/moshe-blox/ssv-spec/types/testingutils/comparable"
)

// F1DifferentFutureRoundsNotPrepared tests f+1 speedup (not prev prepared)
func F1DifferentFutureRoundsNotPrepared() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()
	sc := f1DifferentFutureRoundsNotPreparedStateComparison()

	msgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[1], types.OperatorID(1), 5),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[2], types.OperatorID(2), 10),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "round change f+1 not prepared",
		Pre:           pre,
		PostRoot:      sc.Root(),
		PostState:     sc.ExpectedState,
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingRoundChangeMessageWithParams(ks.Shares[1], types.OperatorID(1), 5, qbft.FirstHeight,
				[32]byte{}, 0, [][]byte{}),
		},
	}
}

func f1DifferentFutureRoundsNotPreparedStateComparison() *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()

	msgs := []*qbft.SignedMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[1], types.OperatorID(1), 5),
		testingutils.TestingRoundChangeMessageWithRound(ks.Shares[2], types.OperatorID(2), 10),
	}

	instance := &qbft.Instance{
		State: &qbft.State{
			Share: testingutils.TestingShare(testingutils.Testing4SharesSet()),
			ID:    testingutils.TestingIdentifier,
			Round: 5,
		},
	}
	comparable.SetSignedMessages(instance, msgs)
	return &comparable.StateComparison{ExpectedState: instance.State}
}
