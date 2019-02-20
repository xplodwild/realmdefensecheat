package realmdefense

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// v3: var kAesIV = []byte{0xE9, 0x4D, 0x7A, 0x51, 0x8D, 0x8E, 0xA8, 0x39, 0x51, 0x8A, 0x61, 0x72, 0xB4, 0x87, 0x2A, 0x10}
var kAesIV = []byte{0x52, 0xAA, 0x1F, 0xCE, 0x81, 0xFC, 0x3E, 0xE7, 0xA4, 0x28, 0x86, 0x19, 0xE5, 0x8D, 0x0C, 0x1E}

// v3 kAesKey can be obtained with: dk := pbkdf2.Key([]byte("A4Q182GGoeYBzOAJ"), []byte{2, 0, 1, 9, 0, 1, 2, 5}, 1000, 32, sha1.New)
// v4 kAesKey can be obtained with: dk := pbkdf2.Key([]byte("oEYbzOAJa4q182GG"), []byte{0, 2, 0, 9, 2, 0, 1, 9}, 1000, 32, sha1.New)
var kAesKey = []byte{
	0x44, 0x4F, 0x69, 0xED, 0x5A, 0x19, 0xD3, 0xEF, 0x14, 0x63, 0xEB, 0x03, 0x6C, 0x99, 0xDE, 0x46,
	0xf0, 0x43, 0xCE, 0x7B, 0x28, 0x45, 0xB1, 0x75, 0x41, 0x5A, 0x8D, 0xFE, 0xF0, 0x9C, 0x30, 0x60,
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		fmt.Printf("unpad error. This could happen when incorrect encryption key is used\n")
		return nil
	}

	return src[:(length - unpadding)]
}

func DecryptCFB(msg []byte) []byte {
	block, err := aes.NewCipher(kAesKey)
	if err != nil {
		panic(err)
	}

	stream := cipher.NewCFBDecrypter(block, kAesIV)
	stream.XORKeyStream(msg, msg)
	return Unpad(msg)
}

func EncryptCFB(msg []byte) []byte {
	block, err := aes.NewCipher(kAesKey)
	if err != nil {
		panic(err)
	}

	msg = Pad(msg)
	stream := cipher.NewCFBEncrypter(block, kAesIV)
	stream.XORKeyStream(msg, msg)
	return msg
}
