package main

import (
	"context"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	println("kmstool - usage: kmstool <key>")

	if len(os.Args) < 2 {
		panic("missing <key>")
	}

	keyName := os.Args[1]

	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to create kms client: %w", err))
	}
	defer client.Close()

	addr, err := resolveAddr(ctx, client, keyName)
	if err != nil {
		panic("failed to retrieve the key")
	}
	fmt.Printf("ethereum addr: %s", addr)
	println()
	println()
}

func resolveAddr(ctx context.Context, client *kms.KeyManagementClient, keyName string) (common.Address, error) {
	resp, err := client.GetPublicKey(ctx, &kmspb.GetPublicKeyRequest{Name: keyName})
	if err != nil {
		return common.Address{}, fmt.Errorf("Google KMS public key %q lookup: %w", keyName, err)
	}

	block, _ := pem.Decode([]byte(resp.Pem))
	if block == nil {
		return common.Address{}, fmt.Errorf("Google KMS public key %q PEM empty: %.130q", keyName, resp.Pem)
	}

	var info struct {
		AlgID pkix.AlgorithmIdentifier
		Key   asn1.BitString
	}
	_, err = asn1.Unmarshal(block.Bytes, &info)
	if err != nil {
		return common.Address{}, fmt.Errorf("Google KMS public key %q PEM block %q: %v", keyName, block.Type, err)
	}

	wantAlg := asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
	if gotAlg := info.AlgID.Algorithm; !gotAlg.Equal(wantAlg) {
		return common.Address{}, fmt.Errorf("Google KMS public key %q ASN.1 algorithm %s intead of %s", keyName, gotAlg, wantAlg)
	}

	return pubKeyAddr(info.Key.Bytes), nil
}

// PubKeyAddr returns the Ethereum address for the (uncompressed) key bytes.
func pubKeyAddr(bytes []byte) common.Address {
	digest := crypto.Keccak256(bytes[1:])
	var addr common.Address
	copy(addr[:], digest[12:])
	return addr
}
