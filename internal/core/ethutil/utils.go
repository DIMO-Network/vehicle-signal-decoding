package ethutil

import (
	"encoding/hex"
	"fmt"
)

func ResolveEtherumAddressFromString(etherumAddress string) ([]byte, error) {
	var ethAddr string
	if len(etherumAddress) > 2 && etherumAddress[:2] == "0x" {
		ethAddr = ethAddr[2:]
	}

	ethAddrBytes, err := hex.DecodeString(ethAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid ethereum address: %w", err)
	}

	return ethAddrBytes, nil
}
