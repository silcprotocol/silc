


package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/silcprotocol/silc/x/evm/types"
)

var (
	//go:embed compiled_contracts/WEVMOS.json
	WSILCJSON []byte

	// WSILCContract is the compiled contract of WSILC
	WSILCContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WSILCJSON, &WSILCContract)
	if err != nil {
		panic(err)
	}

	if len(WSILCContract.Bin) == 0 {
		panic("failed to load WSILC smart contract")
	}
}
