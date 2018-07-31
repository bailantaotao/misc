package main

import (
	//"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/getamis/sirius/log"
	"github.com/spf13/viper"
)

func main() {
	Execute()
}

func Execute() {
	if err := ServerCmd.Execute(); err != nil {
		log.Crit("ServerCmd Execution failed", "err", err)
	}
}

var ServerCmd = &cobra.Command{
	Use:   "indexer",
	Short: "blockchain data indexer",
	Long:  `blockchain data indexer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info("message","i", eth)
		return nil
	},
}

const (
	Eth         = "eth"
	EthProtocol = "eth.protocol"
	EthHost     = "eth.host"
	EthPort     = "eth.port"
)

var (
	ethProtocol string
	eth bool
	ethHost     string
	ethPort     int
)


func init() {
	ServerCmd.Flags().StringVar(&ethProtocol, EthProtocol, "ws", "The eth-client protocol")
	ServerCmd.Flags().StringVar(&ethHost, EthHost, "127.0.0.1", "The eth-client host")
	ServerCmd.Flags().IntVar(&ethPort, EthPort, 8546, "The eth-client port")
	ServerCmd.Flags().BoolVar(&eth, Eth, false, "The eth-client port")

	viper.BindPFlag("eth", ServerCmd.Flags().Lookup("eth"))
	log.Info("m", "eth", viper.GetBool("eth"))
}
