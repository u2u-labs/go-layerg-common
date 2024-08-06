package api

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Send a Metamask Sign In token to the server. Used with authenticate/link/unlink.
type AccountMetamask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Ethereum address associated with the Metamask account.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The signature received from Metamask to validate.
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `protobuf:"bytes,3,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (a *AccountMetamask) verifySignature() (bool, error) {

	// Message that was signed.
	message := "Sample message to sign"

	// Hash the message.
	msgHash := crypto.Keccak256Hash([]byte(message))

	// Decode the signature.
	sigBytes := []byte(a.Signature)
	if len(sigBytes) != 65 {
		return false, errors.New("invalid signature length")
	}

	// Adjust the V value in the signature if needed.
	sigBytes[64] -= 27

	// Recover the public key.
	pubKey, err := crypto.SigToPub(msgHash.Bytes(), sigBytes)
	if err != nil {
		return false, err
	}

	// Get the Ethereum address from the public key.
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	// Check if the recovered address matches the provided address.
	return recoveredAddr == a.Address, nil
}

func (x *AccountMetamask) Reset() {
	*x = AccountMetamask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountMetamask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountMetamask) ProtoMessage() {}

func (x *AccountMetamask) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*AccountMetamask) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *AccountMetamask) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AccountMetamask) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AccountMetamask) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}
