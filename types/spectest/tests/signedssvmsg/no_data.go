package signedssvmsg

import (
	"github.com/moshe-blox/ssv-spec/types"
	"github.com/moshe-blox/ssv-spec/types/testingutils"
)

// NoData tests an invalid SignedSSVMessageTest with empty data
func NoData() *SignedSSVMessageTest {

	return &SignedSSVMessageTest{
		Name: "no data",
		Messages: []*types.SignedSSVMessage{
			{
				OperatorID: 1,
				Signature:  testingutils.TestingSignedSSVMessageSignature,
				Data:       []byte{},
			},
		},
		ExpectedError: "Data has length 0 in SignedSSVMessage",
	}
}
