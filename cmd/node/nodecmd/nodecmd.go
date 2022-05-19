package nodecmd

import (
	"github.com/forta-network/forta-node/cmd/botadmin"
	"github.com/forta-network/forta-node/cmd/hostnet"
	json_rpc "github.com/forta-network/forta-node/cmd/json-rpc"
	"github.com/forta-network/forta-node/cmd/publisher"
	"github.com/forta-network/forta-node/cmd/scanner"
	"github.com/forta-network/forta-node/cmd/supervisor"
	"github.com/forta-network/forta-node/cmd/updater"
	"github.com/spf13/cobra"
)

var (
	cmdFortaNode = &cobra.Command{
		Use: "forta-node",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceUsage: true,
	}

	cmdUpdater = &cobra.Command{
		Use: "updater",
		RunE: func(cmd *cobra.Command, args []string) error {
			updater.Run()
			return nil
		},
	}

	cmdSupervisor = &cobra.Command{
		Use: "supervisor",
		RunE: func(cmd *cobra.Command, args []string) error {
			supervisor.Run()
			return nil
		},
	}

	cmdScanner = &cobra.Command{
		Use: "scanner",
		RunE: func(cmd *cobra.Command, args []string) error {
			scanner.Run()
			return nil
		},
	}

	cmdPublisher = &cobra.Command{
		Use: "publisher",
		RunE: func(cmd *cobra.Command, args []string) error {
			publisher.Run()
			return nil
		},
	}

	cmdJsonRpc = &cobra.Command{
		Use: "json-rpc",
		RunE: func(cmd *cobra.Command, args []string) error {
			json_rpc.Run()
			return nil
		},
	}

	cmdHostnet = &cobra.Command{
		Use: "detect-host-networking",
		RunE: func(cmd *cobra.Command, args []string) error {
			hostnet.Run()
			return nil
		},
	}

	cmdBotAdmin = &cobra.Command{
		Use: "bot-admin",
		RunE: func(cmd *cobra.Command, args []string) error {
			botadmin.Run()
			return nil
		},
	}
)

func init() {
	cmdFortaNode.AddCommand(cmdUpdater)
	cmdFortaNode.AddCommand(cmdSupervisor)
	cmdFortaNode.AddCommand(cmdScanner)
	cmdFortaNode.AddCommand(cmdPublisher)
	cmdFortaNode.AddCommand(cmdJsonRpc)
	cmdFortaNode.AddCommand(cmdHostnet)
	cmdFortaNode.AddCommand(cmdBotAdmin)
}

func Run() error {
	return cmdFortaNode.Execute()
}
