package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func HashPrintUsage() {
	fmt.Println(" Options:")
	fmt.Println(` --mnemonic`)
	fmt.Println(` --print-only`)
}

func HashCli() {
	hashCMD := flag.NewFlagSet("hash", flag.ExitOnError)
	mnemonic := hashCMD.String("mnemonic", "", "Mnemonic to hash")
	printOnly := hashCMD.Bool("print-only", false, "Print without storing a new hashed mnemonic")

	handleParsingError(hashCMD)

	if *mnemonic != "" {
		checkMnemonic(*mnemonic)
		checkPassphraseEnv()

		hash := hashMnemonic(*mnemonic)
		file := FOLDER_PATH + APP_ID
		if !*printOnly {
			checkSession()
			if doesMnemonicHashFileExist() {
				if getMnemonicHashFile() != hash {
					log.Fatal("a mnemonic different than the current one has already been hashed. Please remove the folder .data to perform this action.")
				} else {
					fmt.Println("Hash already recorded")
				}
			} else {
				handleError(os.WriteFile(file, []byte(hash), 0644))
				fmt.Printf("SUCCESSFULLY saved in %s\nhash:%s/n", file, hash)
			}
		} else {
			fmt.Println("hash:", hash)
		}
	} else {
		HashPrintUsage()
	}
}
