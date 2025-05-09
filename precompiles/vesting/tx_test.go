// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)
package vesting_test

import (
	"fmt"
	"math/big"
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	cmn "github.com/silcprotocol/silc/precompiles/common"
	"github.com/silcprotocol/silc/precompiles/testutil"
	"github.com/silcprotocol/silc/precompiles/vesting"
	evmosutil "github.com/silcprotocol/silc/testutil"
	evmosutiltx "github.com/silcprotocol/silc/testutil/tx"
	evmostypes "github.com/silcprotocol/silc/types"
	"github.com/silcprotocol/silc/utils"
	"github.com/silcprotocol/silc/x/evm/core/vm"
	vestingtypes "github.com/silcprotocol/silc/x/vesting/types"
)

var (
	balances         = []cmn.Coin{{Denom: utils.BaseDenom, Amount: big.NewInt(1000)}}
	quarter          = []cmn.Coin{{Denom: utils.BaseDenom, Amount: big.NewInt(250)}}
	balancesSdkCoins = sdk.NewCoins(sdk.NewInt64Coin(utils.BaseDenom, 1000))
	quarterSdkCoins  = sdk.NewCoins(sdk.NewInt64Coin(utils.BaseDenom, 250))
	toAddr           = evmosutiltx.GenerateAddress()
	funderAddr       = evmosutiltx.GenerateAddress()
	diffFunderAddr   = evmosutiltx.GenerateAddress()
	lockupPeriods    = []vesting.Period{{Length: 5000, Amount: balances}}
	sdkLockupPeriods = []sdkvesting.Period{{Length: 5000, Amount: balancesSdkCoins}}
	vestingPeriods   = []vesting.Period{
		{Length: 2000, Amount: quarter},
		{Length: 2000, Amount: quarter},
		{Length: 2000, Amount: quarter},
		{Length: 2000, Amount: quarter},
	}
	sdkVestingPeriods = []sdkvesting.Period{
		{Length: 2000, Amount: quarterSdkCoins},
		{Length: 2000, Amount: quarterSdkCoins},
		{Length: 2000, Amount: quarterSdkCoins},
		{Length: 2000, Amount: quarterSdkCoins},
	}
)

func (s *PrecompileTestSuite) TestCreateClawbackVestingAccount() {
	var ctx sdk.Context

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		gas         uint64
		postCheck   func(data []byte)
		expError    bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			200000,
			func([]byte) {},
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 3, 0),
		},
		{
			name: "fail - different origin than vesting address",
			malleate: func() []interface{} {
				differentAddr := evmosutiltx.GenerateAddress()
				return []interface{}{
					funderAddr,
					differentAddr,
					false,
				}
			},
			gas:         200000,
			expError:    true,
			errContains: "does not match the from address",
		},
		{
			"success",
			func() []interface{} {
				return []interface{}{
					funderAddr,
					s.keyring.GetAddr(0),
					false,
				}
			},
			20000,
			func(data []byte) {
				success, err := s.precompile.Unpack(vesting.CreateClawbackVestingAccountMethod, data)
				s.Require().NoError(err)
				s.Require().Equal(success[0], true)

				// Check if the vesting account was created
				_, err = s.network.App.VestingKeeper.Balances(ctx, &vestingtypes.QueryBalancesRequest{Address: s.keyring.GetAccAddr(0).String()})
				s.Require().NoError(err)
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest(2)
			ctx = s.network.GetContext()
			method := s.precompile.Methods[vesting.CreateClawbackVestingAccountMethod]

			createArgs := tc.malleate()

			bz, err := s.precompile.CreateClawbackVestingAccount(
				ctx,
				s.keyring.GetAddr(0),
				s.network.GetStateDB(),
				&method,
				createArgs,
			)

			if tc.expError {
				s.Require().ErrorContains(err, tc.errContains)
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestFundVestingAccount() {
	var ctx sdk.Context

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		gas         uint64
		postCheck   func(data []byte)
		expError    bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			200000,
			func([]byte) {},
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 5, 0),
		},
		{
			name: "fail - different origin than funder address",
			malleate: func() []interface{} {
				differentAddr := evmosutiltx.GenerateAddress()
				return []interface{}{
					differentAddr,
					toAddr,
					uint64(time.Now().Unix()), //nolint:gosec // G115
					lockupPeriods,
					vestingPeriods,
				}
			},
			gas:         200000,
			expError:    true,
			errContains: "does not match the from address",
		},
		{
			"success",
			func() []interface{} {
				s.CreateTestClawbackVestingAccount(ctx, s.keyring.GetAddr(0), toAddr)
				err = evmosutil.FundAccount(ctx, s.network.App.BankKeeper, toAddr.Bytes(), sdk.NewCoins(sdk.NewCoin(utils.BaseDenom, math.NewInt(100))))
				return []interface{}{
					s.keyring.GetAddr(0),
					toAddr,
					uint64(time.Now().Unix()), //nolint:gosec // G115
					lockupPeriods,
					vestingPeriods,
				}
			},
			20000,
			func(data []byte) {
				success, err := s.precompile.Unpack(vesting.CreateClawbackVestingAccountMethod, data)
				s.Require().NoError(err)
				s.Require().Equal(success[0], true)

				// Check if the vesting account was created
				vestingAcc, err := s.network.App.VestingKeeper.Balances(ctx, &vestingtypes.QueryBalancesRequest{Address: sdk.AccAddress(toAddr.Bytes()).String()})
				s.Require().NoError(err)
				s.Require().Equal(vestingAcc.Locked, balancesSdkCoins)
				s.Require().Equal(vestingAcc.Unvested, balancesSdkCoins)
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest(2)
			ctx = s.network.GetContext()
			method := s.precompile.Methods[vesting.FundVestingAccountMethod]

			var contract *vm.Contract
			contract, ctx = testutil.NewPrecompileContract(s.T(), ctx, s.keyring.GetAddr(0), s.precompile, tc.gas)

			bz, err := s.precompile.FundVestingAccount(ctx, contract, s.keyring.GetAddr(0), s.network.GetStateDB(), &method, tc.malleate())

			if tc.expError {
				s.Require().ErrorContains(err, tc.errContains)
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestClawback() {
	var ctx sdk.Context

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		gas         uint64
		postCheck   func(data []byte)
		expError    bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			200000,
			func([]byte) {},
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 3, 0),
		},
		{
			name: "fail - different origin than funder address",
			malleate: func() []interface{} {
				differentAddr := evmosutiltx.GenerateAddress()
				return []interface{}{
					differentAddr,
					toAddr,
					s.keyring.GetAddr(0),
				}
			},
			gas:         200000,
			expError:    true,
			errContains: "does not match the funder address",
		},
		{
			"success",
			func() []interface{} {
				s.CreateTestClawbackVestingAccount(ctx, s.keyring.GetAddr(0), toAddr)
				s.FundTestClawbackVestingAccount()
				return []interface{}{
					s.keyring.GetAddr(0),
					toAddr,
					s.keyring.GetAddr(0),
				}
			},
			20000,
			func(data []byte) {
				var co vesting.ClawbackOutput
				err := s.precompile.UnpackIntoInterface(&co, vesting.ClawbackMethod, data)
				s.Require().NoError(err, "failed to unpack clawback output")
				s.Require().Equal(co.Coins, balances, "expected different clawed back coins")
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest(2)
			ctx = s.network.GetContext()
			method := s.precompile.Methods[vesting.ClawbackMethod]

			var contract *vm.Contract
			contract, ctx = testutil.NewPrecompileContract(s.T(), ctx, s.keyring.GetAddr(0), s.precompile, tc.gas)

			bz, err := s.precompile.Clawback(ctx, contract, s.keyring.GetAddr(0), s.network.GetStateDB(), &method, tc.malleate())

			if tc.expError {
				s.Require().ErrorContains(err, tc.errContains)
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestUpdateVestingFunder() {
	var ctx sdk.Context

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		gas         uint64
		postCheck   func(data []byte)
		expError    bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			200000,
			func([]byte) {},
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 3, 0),
		},
		{
			name: "fail - different origin than funder address",
			malleate: func() []interface{} {
				differentAddr := evmosutiltx.GenerateAddress()
				return []interface{}{
					differentAddr,
					toAddr,
					s.keyring.GetAddr(0),
				}
			},
			gas:         200000,
			expError:    true,
			errContains: "does not match the funder address",
		},
		{
			"success",
			func() []interface{} {
				s.CreateTestClawbackVestingAccount(ctx, s.keyring.GetAddr(0), toAddr)
				vestingAcc := s.network.App.AccountKeeper.GetAccount(ctx, toAddr.Bytes())
				va, ok := vestingAcc.(*vestingtypes.ClawbackVestingAccount)
				s.Require().True(ok)
				s.Require().Equal(va.FunderAddress, s.keyring.GetAccAddr(0).String())
				return []interface{}{
					s.keyring.GetAddr(0),
					diffFunderAddr,
					toAddr,
				}
			},
			20000,
			func(data []byte) {
				success, err := s.precompile.Unpack(vesting.UpdateVestingFunderMethod, data)
				s.Require().NoError(err)
				s.Require().Equal(success[0], true)

				// Check if the vesting account has a new funder address
				vestingAcc := s.network.App.AccountKeeper.GetAccount(ctx, toAddr.Bytes())
				va, ok := vestingAcc.(*vestingtypes.ClawbackVestingAccount)
				s.Require().True(ok)
				s.Require().Equal(va.FunderAddress, sdk.AccAddress(diffFunderAddr.Bytes()).String())
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest(2)
			ctx = s.network.GetContext()
			method := s.precompile.Methods[vesting.UpdateVestingFunderMethod]

			var contract *vm.Contract
			contract, ctx = testutil.NewPrecompileContract(s.T(), ctx, s.keyring.GetAddr(0), s.precompile, tc.gas)

			bz, err := s.precompile.UpdateVestingFunder(ctx, contract, s.keyring.GetAddr(0), s.network.GetStateDB(), &method, tc.malleate())

			if tc.expError {
				s.Require().ErrorContains(err, tc.errContains)
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestConvertVestingAccount() {
	var ctx sdk.Context

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		gas         uint64
		postCheck   func(data []byte)
		expError    bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			200000,
			func([]byte) {},
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 1, 0),
		},
		{
			"fail - incorrect address",
			func() []interface{} {
				return []interface{}{
					"asda412",
				}
			},
			200000,
			func([]byte) {},
			true,
			"invalid type for vestingAddress",
		},
		{
			"success",
			func() []interface{} {
				s.CreateTestClawbackVestingAccount(ctx, s.keyring.GetAddr(0), toAddr)
				return []interface{}{
					toAddr,
				}
			},
			20000,
			func(data []byte) {
				success, err := s.precompile.Unpack(vesting.ConvertVestingAccountMethod, data)
				s.Require().NoError(err)
				s.Require().Equal(success[0], true)

				// Check if the vesting account was converted back to an EthAccountI
				account := s.network.App.AccountKeeper.GetAccount(s.network.GetContext(), toAddr.Bytes())
				_, ok := account.(evmostypes.EthAccountI)
				s.Require().True(ok)
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest(2)
			ctx = s.network.GetContext()
			method := s.precompile.Methods[vesting.ConvertVestingAccountMethod]

			bz, err := s.precompile.ConvertVestingAccount(ctx, s.network.GetStateDB(), &method, tc.malleate())

			if tc.expError {
				s.Require().ErrorContains(err, tc.errContains)
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				tc.postCheck(bz)
			}
		})
	}
}
