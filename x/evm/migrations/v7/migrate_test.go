// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)

package v7_test

import (
	"encoding/json"
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/stretchr/testify/require"

	"github.com/silcprotocol/silc/encoding"
	v7 "github.com/silcprotocol/silc/x/evm/migrations/v7"
	v6types "github.com/silcprotocol/silc/x/evm/migrations/v7/types"
	"github.com/silcprotocol/silc/x/evm/types"
)

func TestMigrate(t *testing.T) {
	encCfg := encoding.MakeConfig()
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	kvStore := ctx.KVStore(storeKey)

	chainConfig := types.DefaultChainConfig()
	bz, err := json.Marshal(chainConfig)
	require.NoError(t, err)
	var chainCfgV6 v6types.V6ChainConfig
	err = json.Unmarshal(bz, &chainCfgV6)
	require.NoError(t, err)

	// Create a pre migration environment with default params.
	paramsV6 := v6types.V6Params{
		EvmDenom:            types.DefaultEVMDenom,
		ChainConfig:         chainCfgV6,
		ExtraEIPs:           v6types.DefaultExtraEIPs,
		AllowUnprotectedTxs: types.DefaultAllowUnprotectedTxs,
		ActivePrecompiles:   types.DefaultStaticPrecompiles,
		EVMChannels:         types.DefaultEVMChannels,
	}
	paramsV6Bz := cdc.MustMarshal(&paramsV6)
	kvStore.Set(types.KeyPrefixParams, paramsV6Bz)

	err = v7.MigrateStore(ctx, storeKey, cdc)
	require.NoError(t, err)

	paramsBz := kvStore.Get(types.KeyPrefixParams)
	var params types.Params
	cdc.MustUnmarshal(paramsBz, &params)

	require.Equal(t, types.DefaultEVMDenom, params.EvmDenom)
	require.False(t, params.AllowUnprotectedTxs)
	require.Equal(t, chainConfig, params.ChainConfig)
	require.Equal(t, types.DefaultExtraEIPs, params.ExtraEIPs)
	require.Equal(t, types.DefaultEVMChannels, params.EVMChannels)
	require.Equal(t, types.DefaultAccessControl, params.AccessControl)
}
