package keeper_test

import (
	"testing"

	testkeeper "github.com/seosigoto/hello/testutil/keeper"
	"github.com/seosigoto/hello/x/hello/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HelloKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
