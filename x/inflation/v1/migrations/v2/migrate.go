// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)

package v2

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2types "github.com/silcprotocol/silc/x/inflation/v1/migrations/v2/types"
	"github.com/silcprotocol/silc/x/inflation/v1/types"
)

// MigrateStore migrates the x/inflation module state from the consensus version 1 to
// version 2. Specifically, it takes the parameters that are currently stored
// and managed by the Cosmos SDK params module and stores them directly into the x/inflation module state.
func MigrateStore(
	ctx sdk.Context,
	storeKey storetypes.StoreKey,
	legacySubspace types.Subspace,
	cdc codec.BinaryCodec,
) error {
	store := ctx.KVStore(storeKey)
	var params v2types.V2Params

	legacySubspace = legacySubspace.WithKeyTable(v2types.ParamKeyTable())
	legacySubspace.GetParamSetIfExists(ctx, &params)
	if err := params.Validate(); err != nil {
		return err
	}

	bz, err := cdc.Marshal(&params)
	if err != nil {
		return err
	}

	store.Set(types.ParamsKey, bz)

	return nil
}
