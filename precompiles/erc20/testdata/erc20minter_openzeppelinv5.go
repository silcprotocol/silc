// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/silcprotocol/silc/contracts/utils"
	evmtypes "github.com/silcprotocol/silc/x/evm/types"
)

func LoadERC20MinterV5Contract() (evmtypes.CompiledContract, error) {
	return contractutils.LegacyLoadContractFromJSONFile("ERC20Minter_OpenZeppelinV5.json")
}
