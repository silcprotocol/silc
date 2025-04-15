// Copyright Tharsis Labs Ltd.(Silc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package contracts

import (
	contractutils "github.com/silcprotocol/silc/contracts/utils"
	evmtypes "github.com/silcprotocol/silc/x/evm/types"
)

func LoadCounterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("Counter.json")
}
