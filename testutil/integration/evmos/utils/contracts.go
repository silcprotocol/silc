// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)
package utils

import (
	"fmt"
	"slices"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/silcprotocol/silc/testutil/integration/silc/factory"
	evmostypes "github.com/silcprotocol/silc/types"
	evmtypes "github.com/silcprotocol/silc/x/evm/types"
)

// CheckTxTopics checks if all expected topics are present in the transaction response
func CheckTxTopics(res abcitypes.ExecTxResult, expectedTopics []string) error {
	msgEthResponse, err := DecodeExecTxResult(res)
	if err != nil {
		return err
	}

	// Collect all topics within the transaction
	availableLogs := make([]string, 0, len(msgEthResponse.Logs))
	for _, log := range msgEthResponse.Logs {
		availableLogs = append(availableLogs, log.Topics...)
	}

	// Check if all expected topics are present
	for _, expectedTopic := range expectedTopics {
		if !slices.Contains(availableLogs, expectedTopic) {
			return fmt.Errorf("expected topic %s not found in tx response", expectedTopic)
		}
	}
	return nil
}

// IsContractAccount checks if the given account is a contract account
func IsContractAccount(acc sdktypes.AccountI) error {
	contractETHAccount, ok := acc.(evmostypes.EthAccountI)
	if !ok {
		return fmt.Errorf("account is not an eth account")
	}

	if contractETHAccount.Type() != evmostypes.AccountTypeContract {
		return fmt.Errorf("account is not a contract account")
	}
	return nil
}

// DecodeContractCallResponse decodes the response of a contract call query
func DecodeContractCallResponse(response interface{}, callArgs factory.CallArgs, res abcitypes.ExecTxResult) error {
	msgEthResponse, err := DecodeExecTxResult(res)
	if err != nil {
		return err
	}

	err = callArgs.ContractABI.UnpackIntoInterface(response, callArgs.MethodName, msgEthResponse.Ret)
	if err != nil {
		return err
	}
	return nil
}

func DecodeExecTxResult(res abcitypes.ExecTxResult) (*evmtypes.MsgEthereumTxResponse, error) {
	msgEthResponse, err := evmtypes.DecodeTxResponse(res.Data)
	if err != nil {
		return nil, err
	}
	return msgEthResponse, nil
}
