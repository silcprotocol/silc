package keeper_test

import (
	testkeyring "github.com/silcprotocol/silc/testutil/integration/silc/keyring"
	"github.com/silcprotocol/silc/testutil/integration/silc/network"
	evmtypes "github.com/silcprotocol/silc/x/evm/types"
)

func (suite *KeeperTestSuite) TestEndBlock() {
	keyring := testkeyring.New(2)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	ctx := unitNetwork.GetContext()
	preEventManager := ctx.EventManager()
	suite.Require().Equal(0, len(preEventManager.Events()))

	err := unitNetwork.App.EvmKeeper.EndBlock(ctx)
	suite.Require().NoError(err)

	postEventManager := unitNetwork.GetContext().EventManager()
	// should emit 1 EventTypeBlockBloom event on EndBlock
	suite.Require().Equal(1, len(postEventManager.Events()))
	suite.Require().Equal(evmtypes.EventTypeBlockBloom, postEventManager.Events()[0].Type)
}
