package masterdb

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
)

// ECDSASignature holds the r and s values of the signature.
type ECDSASignature struct {
	R, S *big.Int
}

// SignECDSA signs the data using the private key provided in hex format.
func SignECDSA(privateKeyHex string, data []byte) (string, error) {
	// Decode the hex string to bytes
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex private key: %v", err)
	}

	// Create a private key from the byte slice
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = elliptic.P256() // Use the appropriate curve
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)

	// Hash the data
	hashed := sha256.Sum256(data)

	// Sign the hashed data
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// Create a combined signature
	signature := ECDSASignature{R: r, S: s}

	// Marshal the signature to ASN.1 DER format
	signatureBytes, err := asn1.Marshal(signature)
	if err != nil {
		return "", fmt.Errorf("failed to marshal signature: %v", err)
	}

	// Return the signature as a hex string
	return hex.EncodeToString(signatureBytes), nil
}

// VerifyECDSASignature verifies the ECDSA signature using the public key in hex format.
func VerifyECDSASignature(hexPublicKey, hexSignature string, data []byte) (bool, error) {
	publicKey, err := ConvertHexToECDSAPublicKey(hexPublicKey)
	if err != nil {
		return false, err
	}

	signature, err := ConvertHexToECDSASignature(hexSignature)
	if err != nil {
		return false, err
	}

	// Hash the data
	hashed := sha256.Sum256(data)

	// Verify the signature
	return ecdsa.Verify(publicKey, hashed[:], signature.R, signature.S), nil
}

// ConvertHexToECDSAPublicKey converts a hex string to an ECDSA public key.
func ConvertHexToECDSAPublicKey(hexPublicKey string) (*ecdsa.PublicKey, error) {
	pubKeyBytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex public key: %v", err)
	}

	if len(pubKeyBytes) != 65 {
		return nil, fmt.Errorf("invalid public key length: expected 65 bytes, got %d", len(pubKeyBytes))
	}

	// The first byte is the prefix (0x04 for uncompressed)
	x := new(big.Int).SetBytes(pubKeyBytes[1:33])
	y := new(big.Int).SetBytes(pubKeyBytes[33:65])

	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, nil
}

// ConvertHexToECDSASignature converts a hex string to an ECDSA signature.
func ConvertHexToECDSASignature(hexSignature string) (*ECDSASignature, error) {
	signatureBytes, err := hex.DecodeString(hexSignature)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex signature: %v", err)
	}

	var sig ECDSASignature
	_, err = asn1.Unmarshal(signatureBytes, &sig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal signature: %v", err)
	}

	return &sig, nil
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
