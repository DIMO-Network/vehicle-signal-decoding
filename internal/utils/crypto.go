package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const sigLen = 65

var zeroAddr common.Address

// This is exact copy of the function from internal/controllers/helpers/handlers.go
// Ecrecover mimics the ecrecover opcode, returning the address that signed
// hash with signature. sig must have length 65 and the last byte, the recovery
// byte usually denoted v, must be 27 or 28.
func Ecrecover(hash, sig []byte) (common.Address, error) {
	if len(sig) != sigLen {
		return zeroAddr, fmt.Errorf("signature has invalid length %d", len(sig))
	}

	// Defensive copy: the caller shouldn't have to worry about us modifying
	// the signature. We adjust because crypto.Ecrecover demands 0 <= v <= 4.
	fixedSig := make([]byte, sigLen)
	copy(fixedSig, sig)
	fixedSig[64] -= 27

	rawPk, err := crypto.Ecrecover(hash, fixedSig)
	if err != nil {
		return zeroAddr, err
	}

	pk, err := crypto.UnmarshalPubkey(rawPk)
	if err != nil {
		return zeroAddr, err
	}

	return crypto.PubkeyToAddress(*pk), nil
}

// VerifySignature godoc
// @Description checks if the provided signature corresponds to the given message and Ethereum address.
// @param body []byte The original message that was signed.
// @param signature string The signature that needs to be verified.
// @param ethAddr string The Ethereum address of the signer.
// @return bool Indicates whether the signature is valid or not.
// @return error If there was an issue during the verification process.
func VerifySignature(body []byte, signature string, ethAddr string) (bool, error) {
	addr := common.HexToAddress(ethAddr)
	sig := common.FromHex(signature)
	hash := crypto.Keccak256Hash(body)

	recAddr, err := Ecrecover(hash.Bytes(), sig)
	if err != nil {
		return false, err
	}

	return recAddr == addr, nil
}