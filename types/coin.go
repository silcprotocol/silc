

package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoSilc defines the default coin denomination used in Silc in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Silc.
	AttoSilc string = "sillet"

	// BaseDenomUnit defines the base denomination unit for Silc.
	// 1 silc = 1x10^{BaseDenomUnit} sillet
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewSilcCoin is a utility function that returns an "sillet" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewSilcCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoSilc, amount)
}

// NewSilcDecCoin is a utility function that returns an "sillet" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewSilcDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoSilc, amount)
}

// NewSilcCoinInt64 is a utility function that returns an "sillet" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewSilcCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoSilc, amount)
}
