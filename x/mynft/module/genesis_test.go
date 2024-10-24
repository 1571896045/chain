package mynft_test

import (
	"testing"

	keepertest "github.com/1571896045/chain/testutil/keeper"
	"github.com/1571896045/chain/testutil/nullify"
	mynft "github.com/1571896045/chain/x/mynft/module"
	"github.com/1571896045/chain/x/mynft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MynftKeeper(t)
	mynft.InitGenesis(ctx, k, genesisState)
	got := mynft.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
