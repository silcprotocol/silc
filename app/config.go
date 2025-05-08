


package app

import (
	"github.com/silcprotocol/silc/app/eips"
	evmconfig "github.com/silcprotocol/silc/x/evm/config"
	"github.com/silcprotocol/silc/x/evm/core/vm"
)

// The init function of the config file allows to setup the global
// configuration for the EVM, modifying the custom ones defined in silc.
func init() {
	err := evmconfig.NewEVMConfigurator().
		WithExtendedEips(silcActivators).
		Configure()
	if err != nil {
		panic(err)
	}
}

// SilcActivators defines a map of opcode modifiers associated
// with a key defining the corresponding EIP.
var silcActivators = map[string]func(*vm.JumpTable){
	"silc_0": eips.Enable0000,
	"silc_1": eips.Enable0001,
	"silc_2": eips.Enable0002,
}
