


package contracts

import (
	contractutils "github.com/silcprotocol/silc/contracts/utils"
	evmtypes "github.com/silcprotocol/silc/x/evm/types"
)

func LoadInterchainSenderContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("InterchainSender.json")
}
