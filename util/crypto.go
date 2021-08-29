package util

import (
	"coldcrypt/util/bip32"
	"coldcrypt/util/bip39"
)

type Seed []byte

func NewMasterKey() (*bip32.Key, Seed, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, nil, err
	}
	seed := bip39.NewSeed(mnemonic, "")
	mk, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, nil, err
	}
	return mk, seed, nil
}

func NewMasterKeyFromMnemonic(mnemonic string, password string) (*bip32.Key, Seed, error) {
	seed := bip39.NewSeed(mnemonic, password)
	mk, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, nil, err
	}
	return mk, seed, nil
}
