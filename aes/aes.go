package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type AesEncrypt struct {
}

func (this *AesEncrypt) getKey(strKey string) []byte {
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func (this *AesEncrypt) Encrypt(strKey, strMesg string) ([]byte, error) {
	key := this.getKey(strKey)
	fmt.Println("key",string(key))
	var iv = []byte(key)[:aes.BlockSize]
	fmt.Println("key",string(iv))
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil

}
