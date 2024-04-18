package timeout

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils/comparable"
)

// Round15 tests calling UponRoundTimeout for round 15, testing state and broadcasted msgs
func Round15() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	sc := round15StateComparison()

	pre := testingutils.BaseInstance()
	pre.State.Round = 15
	pre.State.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessageWithRound(ks.Shares[1],
		types.OperatorID(1), 15)

	return &SpecTest{
		Name:      "round 15",
		Pre:       pre,
		PostRoot:  sc.Root(),
		PostState: sc.ExpectedState,

		ExpectedTimerState: &testingutils.TimerState{
			Timeouts: 0,
			Round:    0,
		},
		ExpectedError: "instance stopped processing timeouts",
	}
}

func round15StateComparison() *comparable.StateComparison {
	ks := testingutils.Testing4SharesSet()
	state := testingutils.BaseInstance().State
	state.Round = 15
	state.ProposalAcceptedForCurrentRound = testingutils.TestingProposalMessageWithRound(ks.Shares[1],
		types.OperatorID(1), 15)

	return &comparable.StateComparison{ExpectedState: state}
}
