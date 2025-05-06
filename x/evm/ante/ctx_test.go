


package ante_test

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/silcprotocol/silc/testutil/integration/evmos/network"
	evmante "github.com/silcprotocol/silc/x/evm/ante"
)

func (suite *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New()

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	suite.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	suite.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
