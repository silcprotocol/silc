// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/silcprotocol/silc/x/inflation/v1/types"
)

// GetEpochIdentifier gets the epoch identifier
func (k Keeper) GetEpochIdentifier(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixEpochIdentifier)
	if len(bz) == 0 {
		return ""
	}

	return string(bz)
}

// SetEpochsPerPeriod stores the epoch identifier
func (k Keeper) SetEpochIdentifier(ctx sdk.Context, epochIdentifier string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixEpochIdentifier, []byte(epochIdentifier))
}

// GetEpochsPerPeriod gets the epochs per period
func (k Keeper) GetEpochsPerPeriod(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixEpochsPerPeriod)
	if len(bz) == 0 {
		return 0
	}

	return int64(sdk.BigEndianToUint64(bz)) //#nosec G115
}

// SetEpochsPerPeriod stores the epochs per period
func (k Keeper) SetEpochsPerPeriod(ctx sdk.Context, epochsPerPeriod int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixEpochsPerPeriod, sdk.Uint64ToBigEndian(uint64(epochsPerPeriod))) //nolint:gosec // G115
}

// GetSkippedEpochs gets the number of skipped epochs
func (k Keeper) GetSkippedEpochs(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixSkippedEpochs)
	if len(bz) == 0 {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetSkippedEpochs stores the number of skipped epochs
func (k Keeper) SetSkippedEpochs(ctx sdk.Context, skippedEpochs uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixSkippedEpochs, sdk.Uint64ToBigEndian(skippedEpochs))
}
