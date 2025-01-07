package masterdb

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
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

// PrivateKeyFromHex converts a hexadecimal private key string to an *ecdsa.PrivateKey
func PrivateKeyFromHex(hexKey string) (*ecdsa.PrivateKey, error) {
	// Decode the hex string into raw bytes
	privKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, errors.New("invalid hexadecimal string")
	}

	// Use the elliptic P-256 curve (secp256r1) to generate the private key
	curve := elliptic.P256()

	// Create a new big.Int from the private key bytes
	priv := new(big.Int).SetBytes(privKeyBytes)

	// Generate the private key using the big.Int value
	privKey := new(ecdsa.PrivateKey)
	privKey.PublicKey.Curve = curve
	privKey.D = priv
	privKey.PublicKey.X, privKey.PublicKey.Y = curve.ScalarBaseMult(privKey.D.Bytes())

	return privKey, nil
}

// PublicKeyFromHex converts a hexadecimal string to *ecdsa.PublicKey
func PublicKeyFromHex(hexKey string) (*ecdsa.PublicKey, error) {
	// Decode the hex string into raw bytes
	pubKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, errors.New("invalid hexadecimal string")
	}

	// Use the elliptic P-256 curve (secp256r1)
	curve := elliptic.P256()

	// The public key is usually represented as a concatenation of X and Y coordinates
	if len(pubKeyBytes) != 64 { // 32 bytes for X and 32 bytes for Y
		return nil, errors.New("public key must be 64 bytes (32 bytes for X and 32 bytes for Y)")
	}

	// Split the byte slice into X and Y components
	x := new(big.Int).SetBytes(pubKeyBytes[:32])
	y := new(big.Int).SetBytes(pubKeyBytes[32:])

	// Create the public key
	pubKey := &ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	return pubKey, nil
}
