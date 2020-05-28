//
// gpg
//
// - https://tools.ietf.org/html/rfc4880
//
package main

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func main() {
	// rsa
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	// https://tools.ietf.org/html/rfc4880#section-2.1

	// 1.  The sender creates a message.
	const message string = "Hello world, from Gopher"

	// 2.  The sending OpenPGP generates a random number to be used as a
	// session key for this message only.
	sessionKey, err := rand.Prime(rand.Reader, 128)
	if err != nil {
		panic(err)
	}

	// 3.  The session key is encrypted using each recipient's public key.
	// These "encrypted session keys" start the message.
	hash := sha512.New()
	encryptedSessionKey, err := rsa.EncryptOAEP(hash, rand.Reader, &privateKey.PublicKey, []byte(sessionKey.String()), nil)
	if err != nil {
		panic(err)
	}

	base64EncryptedSessionKey := base64.StdEncoding.EncodeToString(encryptedSessionKey)
	fmt.Println(base64EncryptedSessionKey)

	// 4.  The sending OpenPGP encrypts the message using the session key,
	// which forms the remainder of the message.  Note that the message
	// is also usually compressed.

	// First compress the message
	var compressedMessage bytes.Buffer
	w := zlib.NewWriter(&compressedMessage)
	w.Write([]byte(message))
	w.Close()

	// Start AES
	block, err := aes.NewCipher(sessionKey.Bytes())
	if err != nil {
		panic(err)
	}
	// AES Encrypt message
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(compressedMessage.Bytes()))
	cfb.XORKeyStream(ciphertext, compressedMessage.Bytes())
	base64AESEncryptedMessage := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println(base64AESEncryptedMessage)
}
