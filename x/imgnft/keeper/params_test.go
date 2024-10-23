package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/1571896045/chain/testutil/keeper"
	"github.com/1571896045/chain/x/imgnft/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ImgnftKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
