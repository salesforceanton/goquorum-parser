package main

import (
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var dummyAction = func(c *cli.Context) {}

func main() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()

	app.Usage = "Backend integration testing tool"
	app.Commands = []cli.Command{
		{Name: "cryptogate", Usage: "Cryptogate actions", Action: dummyAction, Subcommands: []cli.Command{
			{
				Name:  "test-send-transaction",
				Usage: "Test send transaction to blockchain",
				Action: func(c *cli.Context) {
					contractType := c.String("contract-type")
					nameFunction := c.String("name-function")

					testSendTransaction(contractType, nameFunction)

				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "contract-type",
						Usage: "Contract type",
					},
					&cli.StringFlag{
						Name:  "name-function",
						Usage: "Name function",
					},
				},
			},
			{
				Name:  "test-get-balance-native",
				Usage: "Test get balance native",
				Action: func(c *cli.Context) {
					address := c.String("address")
					testGetNativeBalance(address)

				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Usage: "wallet address",
					},
				},
			},
		}},
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
