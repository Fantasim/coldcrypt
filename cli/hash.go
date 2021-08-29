package cli

import (
	"coldcrypt/util"
	"crypto"
	"flag"
	"fmt"
	"log"
	"os"
)

func hashing(data []byte) []byte {
	return util.Sha256(util.Ripemd160(crypto.BLAKE2b_512.New().Sum((data))))
}

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
			if doesMnemonicHashFileExist() {
				if getMnemonicHashFile() != hash {
					log.Fatal("a mnemonic different than the current one has already been hashed. Please remove the folder .data to perform this action.")
				}
			} else {
				os.WriteFile(file, []byte(hash), 0644)
			}
		} else {
			fmt.Println("hash:", hash)
		}
	} else {
		HashPrintUsage()
	}

}
