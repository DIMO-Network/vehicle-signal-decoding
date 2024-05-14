package common

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/volatiletech/null/v8"
)

func JSONOrDefault(j null.JSON) json.RawMessage {
	if !j.Valid || len(j.JSON) == 0 {
		return []byte(`{}`)
	}
	return j.JSON
}

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

func RemoveSpecialCharacter(input string) string {
	expresionRegular := regexp.MustCompile(`[^\w]`)

	cadenaSinEspeciales := expresionRegular.ReplaceAllString(input, "")

	return cadenaSinEspeciales
}
