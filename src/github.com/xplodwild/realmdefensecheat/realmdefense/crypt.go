package realmdefense

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var kAesIV = []byte{0xE9, 0x4D, 0x7A, 0x51, 0x8D, 0x8E, 0xA8, 0x39, 0x51, 0x8A, 0x61, 0x72, 0xB4, 0x87, 0x2A, 0x10}

// kAesKey can be obtained with: dk := pbkdf2.Key([]byte("A4Q182GGoeYBzOAJ"), []byte{2, 0, 1, 9, 0, 1, 2, 5}, 1000, 32, sha1.New)
var kAesKey = []byte{
	0x16, 0x69, 0x11, 0x86, 0x26, 0x7E, 0x77, 0xF3, 0x11, 0xAE, 0xD5, 0xE8, 0xA1, 0x7F, 0x6B, 0xFC,
	0xB2, 0x52, 0x21, 0xC6, 0xB6, 0xD9, 0xF6, 0xA7, 0x06, 0x83, 0x3E, 0xC6, 0x0E, 0xB6, 0x46, 0x5A,
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
