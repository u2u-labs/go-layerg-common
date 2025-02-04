package masterdb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/unicornultrafoundation/go-u2u/common/hexutil"
	"github.com/unicornultrafoundation/go-u2u/crypto"
)

// SignEthereumMessage signs data using Ethereum's personal_sign format
func SignEthereumMessage(privateKeyHex string, data []byte) (string, error) {
	// Remove 0x prefix if present
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	// Convert private key from hex
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	// Prepare the message according to Ethereum's personal_sign format
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	hash := crypto.Keccak256Hash([]byte(msg))

	// Sign the hash
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %v", err)
	}

	// Convert to hex with 0x prefix
	return hexutil.Encode(signature), nil
}

// VerifyEthereumSignature verifies an Ethereum signature
func VerifyEthereumSignature(address string, data []byte, signatureHex string) (bool, error) {
	// Remove 0x prefix
	address = strings.TrimPrefix(address, "0x")
	signatureHex = strings.TrimPrefix(signatureHex, "0x")

	// Decode signature
	signature, err := hexutil.Decode("0x" + signatureHex)
	if err != nil {
		return false, fmt.Errorf("invalid signature: %v", err)
	}

	// Prepare the message as was done during signing
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	hash := crypto.Keccak256Hash([]byte(msg))

	// Recover public key from signature
	sigPublicKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	// Derive address from public key
	recoveredAddress := crypto.PubkeyToAddress(*sigPublicKey)

	// Compare with expected address
	return strings.EqualFold(recoveredAddress.Hex(), "0x"+address), nil
}

// ConvertToBytes converts a payload of type any to a byte slice.
func ConvertToBytes(payload any) ([]byte, error) {
	// Marshal the payload into JSON format
	bytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err // Return an error if marshaling fails
	}
	return bytes, nil // Return the byte slice
}
