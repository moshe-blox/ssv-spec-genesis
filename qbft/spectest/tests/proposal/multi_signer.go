package proposal

import (
	"github.com/bloxapp/ssv-spec-genesis/qbft"
	"github.com/bloxapp/ssv-spec-genesis/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec-genesis/types"
	"github.com/bloxapp/ssv-spec-genesis/types/testingutils"
	"github.com/herumi/bls-eth-go-binary/bls"
)

// MultiSigner tests a proposal msg with > 1 signers
func MultiSigner() tests.SpecTest {
	pre := testingutils.BaseInstance()
	ks := testingutils.Testing4SharesSet()
	msgs := []*qbft.SignedMessage{
		testingutils.TestingMultiSignerProposalMessage(
			[]*bls.SecretKey{ks.Shares[1], ks.Shares[2]},
			[]types.OperatorID{1, 2},
		),
	}
	return &tests.MsgProcessingSpecTest{
		Name:           "proposal multi signer",
		Pre:            pre,
		PostRoot:       "eaa7264b5d6f05cfcdec3158fcae4ff58c3de1e7e9e12bd876177a58686994d4",
		InputMessages:  msgs,
		OutputMessages: []*qbft.SignedMessage{},
		ExpectedError:  "invalid signed message: msg allows 1 signer",
	}
}
