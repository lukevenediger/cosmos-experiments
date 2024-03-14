package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "experi/testutil/keeper"
	"experi/x/experi/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ExperiKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
