package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AesEncrypt AES 加密
func AesEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	data = PKCS5Padding(data, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	encrypter := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(data))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	encrypter.CryptBlocks(encrypted, data)
	return encrypted, nil
}

// AesDecrypt AES 解密
func AesDecrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	encrypter := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(data))
	// origData := crypted
	encrypter.CryptBlocks(decrypted, data)
	decrypted = PKCS5UnPadding(decrypted)
	// origData = ZeroUnPadding(origData)
	return decrypted, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
