package eth_sign

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

// PersonalSignature
// Sign calculates an Ethereum ECDSA signature for:
// keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
func PersonalSignature(message string, privateKeyStr string) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signature[64] += 27
	return hexutil.Encode(signature), nil
}

// VerifySignature
// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be 27 or 28 for legacy reasons.
func VerifySignature(message, address, signature string) (bool, error) {
	signatureBytes := hexutil.MustDecode(signature)
	if len(signatureBytes) != crypto.SignatureLength {
		return false, fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}
	if signatureBytes[crypto.RecoveryIDOffset] != 27 && signatureBytes[crypto.RecoveryIDOffset] != 28 {
		return false, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	signatureBytes[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	hash := accounts.TextHash([]byte(message))
	recovered, err := crypto.SigToPub(hash, signatureBytes)
	if err != nil {
		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	return strings.ToLower(address) == strings.ToLower(recoveredAddr.Hex()), err
}
