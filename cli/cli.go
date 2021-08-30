package cli

import (
	"flag"
	"fmt"
	"os"
)

const (
	FOLDER_PATH = "./.data/"
	APP_ID      = "399c183f-394a-4d16-93f2-c52e3eafca9f"
)

func init() {
	if _, err := os.Stat(FOLDER_PATH); os.IsNotExist(err) {
		err := os.Mkdir(FOLDER_PATH, os.ModePerm)
		handleError(err)
	}
}

type CLI struct {
	//empty struct
}

func (cli *CLI) printUsage() {
	fmt.Println("Commands:")
	fmt.Println(" pem \t Manage private key")
	fmt.Println(" mnemonic \t Manage mnemonic")
	fmt.Println(" hash \t Manage hashing")
	fmt.Println(" session \t Manage session")
}

//Demarre le cli
func Start() {
	cli := new(CLI)
	cli.validateArgs()
	flag.Parse()
	cli.listMenu()

}

//Verifie les arguments
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

//la liste des commandes
func (cli *CLI) listMenu() {
	switch os.Args[1] {
	case "pem":
		PEMPrintCli()
	case "mnemonic":
		MenmonicCli()
	case "hash":
		HashCli()
	default:
		cli.printUsage()
	}
}
