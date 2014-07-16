package encryption

import (
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
)

func Encrypt(log chan string, dest_pubkey []byte, plainText string) *EncryptedMessage {
	// Generate New Public/Private Key Pair
	D1, X1, Y1 := CreateKey(log)
	// Unmarshal the Destination's Pubkey
	X2, Y2 := elliptic.Unmarshal(elliptic.P256(), dest_pubkey)

	// Point Multiply to get new Pubkey
	PubX, PubY := elliptic.P256().ScalarMult(X2, Y2, D1)

	// Generate Pubkey hashes
	PubHash := sha512.Sum512(elliptic.Marshal(elliptic.P256(), PubX, PubY))
	PubHash_E := PubHash[:32]
	PubHash_M := PubHash[32:64]

	IV, cipherText, _ := SymmetricEncrypt(PubHash_E, plainText)

	// Generate HMAC
	mac := hmac.New(sha256.New, PubHash_M)
	mac.Write(cipherText)
	HMAC := mac.Sum(nil)

	ret := new(EncryptedMessage)
	copy(ret.IV[:], IV[:])
	copy(ret.PublicKey[:], elliptic.Marshal(elliptic.P256(), X1, Y1))
	ret.CipherText = cipherText
	copy(ret.HMAC[:], HMAC)

	return ret
}

// checkMAC returns true if messageMAC is a valid HMAC tag for message.
func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func Decrypt(log chan string, privKey []byte, encrypted *EncryptedMessage) []byte {
	if encrypted == nil || privKey == nil || log == nil {
		return nil
	}

	// Unmarshal the Sender's Pubkey
	X2, Y2 := elliptic.Unmarshal(elliptic.P256(), encrypted.PublicKey[:])

	// Point Multiply to get the new Pubkey
	PubX, PubY := elliptic.P256().ScalarMult(X2, Y2, privKey)

	// Generate Pubkey hashes
	PubHash := sha512.Sum512(elliptic.Marshal(elliptic.P256(), PubX, PubY))
	PubHash_E := PubHash[:32]
	PubHash_M := PubHash[32:64]

	// Check HMAC
	if !checkMAC(encrypted.CipherText[:], encrypted.HMAC[:], PubHash_M) {
		log <- "Invalid HMAC Message"
		return nil
	}

	return SymmetricDecrypt(encrypted.IV, PubHash_E, encrypted.CipherText)
}