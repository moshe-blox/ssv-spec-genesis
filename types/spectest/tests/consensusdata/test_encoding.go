package consensusdata

import (
	reflect2 "reflect"
	"testing"

	"github.com/bloxapp/ssv-spec-genesis/types"
	comparable2 "github.com/bloxapp/ssv-spec-genesis/types/testingutils/comparable"
	"github.com/stretchr/testify/require"
)

type EncodingTest struct {
	Name         string
	Data         []byte
	ExpectedRoot [32]byte
}

func (test *EncodingTest) TestName() string {
	return test.Name
}

func (test *EncodingTest) Run(t *testing.T) {
	// decode
	decoded := &types.ConsensusData{}
	require.NoError(t, decoded.Decode(test.Data))

	// encode
	byts, err := decoded.Encode()
	require.NoError(t, err)
	require.EqualValues(t, test.Data, byts)

	// test root
	r, err := decoded.HashTreeRoot()
	require.NoError(t, err)
	require.EqualValues(t, test.ExpectedRoot, r)

	comparable2.CompareWithJson(t, test, test.TestName(), reflect2.TypeOf(test).String())
}
