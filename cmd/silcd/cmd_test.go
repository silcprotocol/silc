package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/silcprotocol/silc/app"
	silcd "github.com/silcprotocol/silc/cmd/silcd"
	"github.com/silcprotocol/silc/utils"
)

func TestInitCmd(t *testing.T) {
	target := t.TempDir()

	rootCmd, _ := silcd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"silc-test", // Moniker
		fmt.Sprintf("--home=%s", target),
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, utils.TestnetChainID+"-1"),
	})

	err := svrcmd.Execute(rootCmd, "silcd", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := silcd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "SILCD", app.DefaultNodeHome)
	require.Error(t, err)
}
