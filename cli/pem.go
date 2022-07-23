package cli

import (
	"coldcrypt/util"
	"coldcrypt/util/bip32"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func PemPrintUsage() {
	fmt.Println(" Options:")
	fmt.Println(`	--new `)
	fmt.Println(`	--mnemonic [mnemonic-string]`)
	fmt.Println(`	--index [index-int]		Create and add new pem into list`)
}

func newPem(mnemonic string, childIndex int) {

	toB64 := func(key *bip32.Key) string {
		serial, _ := key.Serialize()
		return base64.URLEncoding.EncodeToString([]byte(serial))
	}

	writeFile := func(s string) error {
		content := "-----BEGIN EC PRIVATE KEY-----\n"
		for i := 64; i < len(s); i += 65 {
			s = s[:i] + "\n" + s[i:]
		}
		content += s
		content += "\n-----END EC PRIVATE KEY-----"

		file := FOLDER_PATH + strconv.Itoa(childIndex) + ".pem"
		return ioutil.WriteFile(file, []byte(content), 0644)
	}

	m, _, err := util.NewMasterKeyFromMnemonic(mnemonic, getPassphrase())
	handleError(err)
	k, err := m.NewChildKey(uint32(childIndex))
	handleError(err)
	handleError(writeFile(toB64(k)))

	fmt.Println("Successfully generated")
}

func PEMPrintCli() {
	PemCMD := flag.NewFlagSet("pem", flag.ExitOnError)
	new := PemCMD.Bool("new", false, "Create and mine new block")
	mnemonic := PemCMD.String("mnemonic", "", "mnemonic")
	index := PemCMD.Int("index", 1, "master child index")

	handleParsingError(PemCMD)

	if *new == true {
		checkSession()
		if *index < 1 {
			log.Fatal("Wrong child index")
		}
		checkMnemonic(*mnemonic)
		checkPassphraseEnv()

		checkHashEqualMnemonic(*mnemonic)
		newPem(*mnemonic, *index)
	} else {
		PemPrintUsage()
	}
}
