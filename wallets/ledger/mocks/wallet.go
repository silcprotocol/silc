// Copyright Tharsis Labs Ltd.(Silc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	apitypes "github.com/ethereum/go-ethereum/signer/core/apitypes"
	accounts "github.com/silcprotocol/silc/wallets/accounts"

	big "math/big"

	go_ethereumaccounts "github.com/ethereum/go-ethereum/accounts"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Wallet is an autogenerated mock type for the Wallet type
type Wallet struct {
	mock.Mock
}

// Accounts provides a mock function with given fields:
func (_m *Wallet) Accounts() []accounts.Account {
	ret := _m.Called()

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func() []accounts.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Wallet) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Contains provides a mock function with given fields: account
func (_m *Wallet) Contains(account accounts.Account) bool {
	ret := _m.Called(account)

	var r0 bool
	if rf, ok := ret.Get(0).(func(accounts.Account) bool); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Derive provides a mock function with given fields: path, pin
func (_m *Wallet) Derive(path go_ethereumaccounts.DerivationPath, pin bool) (accounts.Account, error) {
	ret := _m.Called(path, pin)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(go_ethereumaccounts.DerivationPath, bool) accounts.Account); ok {
		r0 = rf(path, pin)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(go_ethereumaccounts.DerivationPath, bool) error); ok {
		r1 = rf(path, pin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Open provides a mock function with given fields: passphrase
func (_m *Wallet) Open(passphrase string) error {
	ret := _m.Called(passphrase)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(passphrase)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignTx provides a mock function with given fields: account, tx, chainID
func (_m *Wallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) ([]byte, error) {
	ret := _m.Called(account, tx, chainID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(accounts.Account, *types.Transaction, *big.Int) []byte); ok {
		r0 = rf(account, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accounts.Account, *types.Transaction, *big.Int) error); ok {
		r1 = rf(account, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignTypedData provides a mock function with given fields: account, typedData
func (_m *Wallet) SignTypedData(account accounts.Account, typedData apitypes.TypedData) ([]byte, error) {
	ret := _m.Called(account, typedData)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(accounts.Account, apitypes.TypedData) []byte); ok {
		r0 = rf(account, typedData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accounts.Account, apitypes.TypedData) error); ok {
		r1 = rf(account, typedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields:
func (_m *Wallet) Status() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URL provides a mock function with given fields:
func (_m *Wallet) URL() go_ethereumaccounts.URL {
	ret := _m.Called()

	var r0 go_ethereumaccounts.URL
	if rf, ok := ret.Get(0).(func() go_ethereumaccounts.URL); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(go_ethereumaccounts.URL)
	}

	return r0
}

type mockConstructorTestingTNewWallet interface {
	mock.TestingT
	Cleanup(func())
}

// NewWallet creates a new instance of Wallet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWallet(t mockConstructorTestingTNewWallet) *Wallet {
	mock := &Wallet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
