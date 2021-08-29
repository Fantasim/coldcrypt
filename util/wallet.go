package util

import (
	"bytes"
	"coldcrypt/util/basen"
	"encoding/hex"
	"strconv"
)

const (
	Version            = byte(0x00)
	AddressChecksumLen = 4 //checksumlen du Bitcoin
)

//Vérifie qu'une adresse est correcte (processus utilisé par le BTC)
func IsAddressValid(addr string) bool {
	//base58 to pubkey hash
	pubKeyHash, err := basen.Base58.DecodeString(string(addr[1:]))
	if err != nil {
		return false
	}
	//on recupere le checksum de la clé publique hashé
	actualChecksum := pubKeyHash[len(pubKeyHash)-AddressChecksumLen:]
	//on recupere la version
	version := []byte(addr)[0] - 48 - 1
	//on recupere le contenu de la clé public hashé entre la version (Index = 1) et le checksum (Index = len - 4)
	pubKeyHash = pubKeyHash[:len(pubKeyHash)-AddressChecksumLen]

	//on créer un checksum correspondant au resultat de la reconstution de la clé public hashé
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	//si les deux checksum sont identiques l'adresse est valide
	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

func GetAddressFromPubKeyHash(pubKeyHash []byte) []byte {
	versionedPayload := append([]byte{Version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(pubKeyHash, checksum...)
	address := basen.Base58.EncodeToString(fullPayload)
	return []byte(strconv.Itoa(int(Version)+1) + address)
}

func PubKeyHashFromAddress(address []byte) []byte {
	pubKeyHash, _ := basen.Base58.DecodeString(string(address[1:]))
	pubKeyHash = pubKeyHash[:len(pubKeyHash)-AddressChecksumLen]
	return pubKeyHash
}

func ToPubKH(pubk []byte) []byte {
	return Ripemd160(Sha256(pubk))
}

func checksum(payload []byte) []byte {
	doubleSha := Sha256(Sha256(payload))
	return doubleSha[:AddressChecksumLen]
}

func AddrSliceToPKHHexSlice(addrs []string) []string {
	ret := make([]string, len(addrs))
	for i := 0; i < len(addrs); i++ {
		ret[i] = hex.EncodeToString(PubKeyHashFromAddress([]byte(addrs[i])))
	}
	return ret
}
