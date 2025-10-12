package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"net/url"
)

var (
	//IV 偏移量
	IV = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

// SHA256 sha256 编码
func SHA256(data []byte) []byte {
	ret := sha256.Sum256(data)
	return ret[:]
}

// Base64Decode Base64 解码
func Base64Decode(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

// LoadKey 读取解密密钥
func LoadKey(EncryptKey string) []byte {
	Key := SHA256([]byte(EncryptKey))
	return Key[:32]
}

// AESDecrypt AES 解密
func AESDecrypt(EncryptKey string, ciphertext []byte) ([]byte, error) {
	key := LoadKey(EncryptKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, IV)
	plainText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plainText, ciphertext)
	plainText = PKCS7UnPadding(plainText)
	return plainText, nil
}

// PKCS7UnPadding 对齐
func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])
	return plainText[:(length - unpadding)]
}

// Decode 入口函数
func Decode(str string, EncryptKey string) ([]byte, error) {
	var err error
	var decoded, raw []byte
	decoded, err = Base64Decode(str)
	if err != nil {
		panic(err)
	}
	raw, err = AESDecrypt(EncryptKey, decoded)
	if err != nil {
		panic(err)
	}
	return raw, err
}

// 产生随机字符串
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	lettersLength := big.NewInt(int64(len(letters)))
	ret := make([]byte, n)
	for i := range n {
		num, err := rand.Int(rand.Reader, lettersLength)
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}

// HmacSha256 计算HmacSha256
// key 是加密所使用的key
// data 是加密的内容
func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

func Base64Encode(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}

func encodeURI(uri string) string {
	//js encodeURI(uri)
	return url.PathEscape(uri)
}
