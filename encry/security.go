package encry

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

//md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cp := h.Sum(nil)
	return hex.EncodeToString(cp)
}

//sha256加密
func Sha256Byte(str string) []byte {
	h := sha256.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

//sha256加密
func Sha256(str string) string {
	return fmt.Sprintf("%x", Sha256Byte(str))
}

//sha1加密
func Sha1Byte(str string) []byte {
	h := sha1.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

//sha1加密
func Sha1(str string) string {
	return fmt.Sprintf("%x", Sha1Byte(str))
}

//sha512加密
func Sha512Byte(str string) []byte {
	h := sha512.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

//sha512加密
func Sha512(str string) string {
	return fmt.Sprintf("%x", Sha512Byte(str))
}

//hmac加密
func Hmac(str string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, str)
	r := fmt.Sprintf("%x", h.Sum(nil))
	return r

}

//HmacSHA1 签名采用HmacSHA1算法 + Base64，编码采用UTF-8
func HmacSHA1(str string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	io.WriteString(h, str)
	r := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return r
}

//HMACSHA256
func HMACSHA256(data []byte, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

//hmac+hex加密
func HmacHex(str string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

//base64编码
func Base64Encode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

//base64解码
func Base64Decode(str string) string {
	st, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		st, err = base64.StdEncoding.DecodeString(str)
		if err != nil {
			return ""
		}

	}
	return string(st)
}

//base64STD编码
func Base64StdEncode(str string) string {
	f := []byte(str)
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(f)))
	base64.StdEncoding.Encode(payload, f)
	return string(payload)
}

//base64STD解码
func Base64StdDecode(str string) string {
	st, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(st)
}

//base64Raw编码
func Base64RawEncode(str string) string {
	f := []byte(str)
	payload := make([]byte, base64.RawStdEncoding.EncodedLen(len(f)))
	base64.RawStdEncoding.Encode(payload, f)
	return string(payload)
}

//base64STD解码
func Base64RawDecode(str string) string {
	st, err := base64.RawStdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return ""
	}
	return string(st)
}

/**
RSA加密解密实例:
m := security.RsaEncrypt("xfdsfsd", `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)
	fmt.Println(security.RsaDecrypt(m, `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`))
**/
//Rsa加密
func RsaEncryptByte(data []byte, publicKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

//Rsa加密
//参数data:要加密的数据
//参数publicKey:公钥
//返回值:加密后数据
func RsaEncrypt(data string, publicKey string) string {
	ndata, err := RsaEncryptByte([]byte(data), publicKey)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return hex.EncodeToString(ndata)
}

//Rsa解密
func RsaDecryptByte(ciphertext []byte, privateKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//Rsa解密
//参数data:要解密的数据
//privateKey:私钥
//返回值:解密后数据
func RsaDecrypt(data string, privateKey string) string {
	by, err1 := hex.DecodeString(data)
	if err1 != nil {
		fmt.Println("加密报错RsaDecrypt:", err1.Error())
		return ""
	}
	ndata, err2 := RsaDecryptByte(by, privateKey)
	if err2 != nil {
		fmt.Println(err2.Error())
		return ""
	}
	return string(ndata)
}

/***
Des加密解密算法
***/

//CBC加密
//参数src:要加密的数据
//参数key:密钥,长度必须是8位数不能超过
//返回值:加密后的数据
func DesEncryptCBC(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	data = PKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

//CBC解密
//参数src:要解密的数据
//参数key:密钥,长度必须是8位数不能超过
//返回值:解密后的数据
func DesDecryptCBC(src, key string) string {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		//fmt.Println("加密报错DesDecryptCBC:", err.Error())
		return ""
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext)
}

//ECB加密
//参数src:要加密的数据
//参数key:密钥,长度必须是8位数不能超过
//返回值:加密后的数据
func DesEncryptECB(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		fmt.Println("Need a multiple of the blocksize")
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}


//ECB解密
//参数src:要解密的数据
//参数key:密钥,长度必须是8位数不能超过
//返回值:解密后的数据
func DesDecryptECB(src, key string) string {
	data, err := hex.DecodeString(src)
	if err != nil {
		fmt.Println("加密报错DesDecryptECB:", err.Error())
		return ""
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		fmt.Println("crypto/cipher: input not full blocks")
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

//明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length >= unpadding {
		return origData[:(length - unpadding)]
	}
	return nil

}

//AES-128加密
func Aes128Encrypt(origData, key []byte, IV []byte) ([]byte, error) {

	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES-128解密
func Aes128Decrypt(crypted, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

//加密
//参数src：原始文本
//参数sKey：加密密钥，非base64的
//参数ivParameter：加密向量，非base64的
//返回值：base64的加密结果
func AES128CBCEncrypt(src string, sKey string, ivParameter string) string {
	key := []byte(sKey)
	iv := []byte(ivParameter)

	result, err := Aes128Encrypt([]byte(src), key, iv)
	if err != nil {
		panic(err)
	}
	return Base64StdEncode(string(result))
}

//解密
//参数src：加密结果base64的
//参数sKey：加密密钥，非base64的
//参数ivParameter：加密向量，非base64的
//返回值：原始文本内容
func AES128CBCDecrypt(src string, sKey string, ivParameter string) string {

	key := []byte(sKey)
	iv := []byte(ivParameter)

	var result []byte
	var err error

	result = ([]byte)(Base64StdDecode(src))

	if len(result) == 0 {
		return ""
	}
	origData, err := Aes128Decrypt(result, key, iv)

	if err != nil {
		panic(err)
	}
	//fmt.Println(origData)
	return string(origData)

}

//计算token 虾皮api需要用到
func CalToken(redirectURL, partnerKey string) (result string) {
	baseStr := partnerKey + redirectURL
	h := sha256.New()
	h.Write([]byte(baseStr))
	result = hex.EncodeToString(h.Sum(nil))
	return result
}

//Rsa2 加密
func Rsa2(origData string, block []byte) (sign string) {
	blocks, _ := pem.Decode(block)
	privateKey, _ := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	h := sha256.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)
	return
}
