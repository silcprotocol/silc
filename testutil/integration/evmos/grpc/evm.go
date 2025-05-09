// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)
package grpc

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/silcprotocol/silc/x/evm/core/vm"
	evmtypes "github.com/silcprotocol/silc/x/evm/types"
)

// GetEvmAccount returns the EVM account for the given address.
func (gqh *IntegrationHandler) GetEvmAccount(address common.Address) (*evmtypes.QueryAccountResponse, error) {
	evmClient := gqh.network.GetEvmClient()
	return evmClient.Account(context.Background(), &evmtypes.QueryAccountRequest{
		Address: address.String(),
	})
}

// EstimateGas returns the estimated gas for the given call args.
func (gqh *IntegrationHandler) EstimateGas(args []byte, gasCap uint64) (*evmtypes.EstimateGasResponse, error) {
	evmClient := gqh.network.GetEvmClient()
	res, err := evmClient.EstimateGas(context.Background(), &evmtypes.EthCallRequest{
		Args:   args,
		GasCap: gasCap,
	})
	if err != nil {
		return nil, err
	}

	// handle case where there's a revert related error
	if res.Failed() {
		if (res.VmError != vm.ErrExecutionReverted.Error()) || len(res.Ret) == 0 {
			return nil, errors.New(res.VmError)
		}
		return nil, evmtypes.NewExecErrorWithReason(res.Ret)
	}

	return res, err
}

// GetEvmParams returns the EVM module params.
func (gqh *IntegrationHandler) GetEvmParams() (*evmtypes.QueryParamsResponse, error) {
	evmClient := gqh.network.GetEvmClient()
	return evmClient.Params(context.Background(), &evmtypes.QueryParamsRequest{})
}
