package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/1571896045/chain/testutil/keeper"
	"github.com/1571896045/chain/x/chain/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.ChainKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
