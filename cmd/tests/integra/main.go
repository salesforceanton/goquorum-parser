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
				Name:  "test-mint-token",
				Usage: "Test mint token",
				Action: func(c *cli.Context) {
					contractType := c.String("contract-type")
					amount := c.String("amount")
					isPrivate := c.Bool("private")

					testMintToken(contractType, amount, isPrivate)

				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "contract-type",
						Usage: "Contract type",
					},
					&cli.StringFlag{
						Name:  "amount",
						Usage: "amount",
					},
					&cli.StringFlag{
						Name:  "private",
						Usage: "private",
					},
				},
			},
			{
				Name:  "test-transfer-token",
				Usage: "Test transfer token",
				Action: func(c *cli.Context) {
					contractType := c.String("contract-type")
					userAddress := c.String("user-address")
					amount := c.String("amount")

					testTransferToken(contractType, amount, userAddress)

				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "contract-type",
						Usage: "Contract type",
					},
					&cli.StringFlag{
						Name:  "user-address",
						Usage: "user address",
					},
					&cli.StringFlag{
						Name:  "amount",
						Usage: "amount",
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
			{
				Name:  "test-get-balance-token",
				Usage: "Test get balance token",
				Action: func(c *cli.Context) {
					address := c.String("address")
					contract := c.String("contract-type")

					testGetBalanceToken(contract, address)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Usage: "wallet address",
					},
					&cli.StringFlag{
						Name:  "contract-type",
						Usage: "Contract type",
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
