package experi_test

import (
	"testing"

	keepertest "experi/testutil/keeper"
	"experi/testutil/nullify"
	experi "experi/x/experi/module"
	"experi/x/experi/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExperiKeeper(t)
	experi.InitGenesis(ctx, k, genesisState)
	got := experi.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
