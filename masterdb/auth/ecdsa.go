package auth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"math/big"
)

func SignPayload(priv *ecdsa.PrivateKey, payloadBytes []byte) (string, error) {
	// Hash the JSON payload
	hash := sha256.Sum256(payloadBytes)

	// Sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return "", err
	}

	// Encode the signature to a base64 string
	signature := append(r.Bytes(), s.Bytes()...)
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySignature verifies the signature of the given payload using the public key
func VerifySignature(pub *ecdsa.PublicKey, payloadBytes []byte, signature string) (bool, error) {
	// Hash the JSON payload
	hash := sha256.Sum256(payloadBytes)

	// Decode the signature from base64
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	// Split the signature into r and s
	r := new(big.Int).SetBytes(sigBytes[:len(sigBytes)/2])
	s := new(big.Int).SetBytes(sigBytes[len(sigBytes)/2:])

	// Verify the signature
	return ecdsa.Verify(pub, hash[:], r, s), nil
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