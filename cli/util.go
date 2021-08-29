package cli

import (
	"encoding/hex"
	"flag"
	"log"
	"os"
	"strings"
)

func handleParsingError(set *flag.FlagSet) {
	err := set.Parse(os.Args[2:])
	if err != nil {
		log.Panic(err)
		os.Exit(2)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkPassphraseEnv() {
	if os.Getenv("emeraud") == "" {
		log.Fatal("emeraud env is empty")
	}
}

func getPassphrase() string {
	checkPassphraseEnv()
	return os.Getenv("emeraud")
}

func doesMnemonicHashFileExist() bool {
	file := FOLDER_PATH + APP_ID
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func getMnemonicHashFile() string {
	file := FOLDER_PATH + APP_ID
	if doesMnemonicHashFileExist() {
		content, err := os.ReadFile(file)
		handleError(err)
		return string(content)
	}
	log.Fatal("mnemonic hash file has not been created. Please create it with the command `hash`")
	return ""
}

func checkMnemonic(mnemonic string) {
	lenMne := len(strings.Split(mnemonic, " "))
	if lenMne != 12 && lenMne != 24 {
		log.Fatal("wrong mnemonic length")
	}
}

func hashMnemonic(mnemonic string) string {
	return hex.EncodeToString(([]byte(mnemonic + getPassphrase())))
}

func checkHashEqualMnemonic(mnemonic string) {
	if hashMnemonic(mnemonic) != getMnemonicHashFile() {
		log.Fatal("current hashed mnemonic and used mnemonic in this action are different. Please fix the difference or remove the hash file.")
	}
}
