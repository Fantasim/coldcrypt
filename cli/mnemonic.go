package cli

import (
	"coldcrypt/util/bip39"
	"flag"
	"fmt"
)

func MnemonicPrintUsage() {
	fmt.Println(" Options:")
	fmt.Println(` --generate`)
	fmt.Println(` --entropy (default 128)`)
}

func MenmonicCli() {
	mnemonicCMD := flag.NewFlagSet("mnemonic", flag.ExitOnError)
	generate := mnemonicCMD.Bool("generate", false, "Generate a new mnemonic")
	entropy := mnemonicCMD.Uint("entropy", 128, "mnemonic entropy (256 or 128) default: 128")

	handleParsingError(mnemonicCMD)

	if *generate == true {
		en, err := bip39.NewEntropy(int(*entropy))
		handleError(err)
		mne, err := bip39.NewMnemonic(en)
		handleError(err)
		fmt.Println(mne)
	} else {
		MnemonicPrintUsage()
	}
}
