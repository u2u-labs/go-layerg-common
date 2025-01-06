package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type HmacPayload struct {
	WalletAddress string `json:"walletAddress"`
	IpAddress     string `json:"ip"`
}

type ClientHmacRequest struct {
	Signature string `json:"signature"`
	IP        string `json:"ip"`
}

func getMessageDigestOrSignature(msg, key []byte) (string, error) {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(msg)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func GenHmacSignature(payload HmacPayload, secretKey string) ClientHmacRequest {

	body, err := json.Marshal(payload)
	if err != nil {
		return ClientHmacRequest{}
	}

	signature, err := getMessageDigestOrSignature(body, []byte(secretKey))
	if err != nil {
		return ClientHmacRequest{}

	}
	return ClientHmacRequest{
		Signature: signature,
		IP:        payload.IpAddress,
	}
}

func VerifyHmac(msg, key []byte, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil)), nil
}
