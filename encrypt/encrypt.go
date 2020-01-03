package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

/**
Encrypt 加密字符串
*/
func Encrypt(b []byte) string {
	mb := md5.Sum(b)
	sb := sha512.Sum384(mb[:])
	return hex.EncodeToString(sb[:])
}

/**
	GenRsaKey 非对称加密生成公钥和私钥
 */
func GenRsaKey(bits int, privateName string, publicName string) ([]byte, []byte, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	x509Private := x509.MarshalPKCS1PrivateKey(private)
	privateBlock := pem.Block{
		Type:  privateName,
		Bytes: x509Private,
	}
	privateBytes := pem.EncodeToMemory(&privateBlock)
	public := private.PublicKey
	x509Public, err := x509.MarshalPKIXPublicKey(&public)
	if err != nil {
		return nil, nil, err
	}
	publicBlock := pem.Block{
		Type:  publicName,
		Bytes: x509Public,
	}
	publicBytes := pem.EncodeToMemory(&publicBlock)
	return privateBytes, publicBytes, nil
}

/**
	RsaEncrypt 使用公钥编码
 */
func RsaEncrypt(key []byte, msg []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), msg)
}

/**
	RsaDecrypt 使用私钥解码
*/
func RsaDecrypt(key []byte, msg []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, msg)
}
