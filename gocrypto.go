package gocrypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	mathrand "math/rand"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltMinLen = 8
	saltMaxLen = 32
	iter       = 1000
	keyLen     = 32
)

// 生成盐值
func randSalt() ([]byte, error) {
	// 生成8-32之间的随机数字
	salt := make([]byte, mathrand.Intn(saltMaxLen-saltMinLen)+saltMinLen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// 通过盐值参数进行加密
func encryptPwdWithSalt(pwd, salt []byte) (pwdEn []byte) {
	pwd = append(pwd, salt...)
	pwdEn = pbkdf2.Key(pwd, salt, iter, keyLen, sha256.New)
	return
}

// EncryptPwd 密码加密
func EncryptPwd(pwd string) (encrypt string, err error) {
	if len(pwd) < 1 {
		return "", errors.New("pwd string length can not be 0")
	}

	// 1. 生成随机长度盐值
	salt, err := randSalt()
	if err != nil {
		return "", err
	}

	// 2. 生成加密串
	en := encryptPwdWithSalt([]byte(pwd), salt)
	en = append(en, salt...)

	// 3. 合并盐值
	encrypt = base64.StdEncoding.EncodeToString(en)

	return
}

// 验证密码是否与加密串匹配
func CheckEncryptPwdMatch(pwd, encrypt string) (ok bool) {
	// 1 参数校验
	if len(encrypt) == 0 {
		fmt.Println("Error: encrypt can not be 0 lenght!")
		return
	}

	enDecode, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println("encrypt decode string err:", err)
		return
	}

	// 2. 获取加密串固定长度
	salt := enDecode[keyLen:]

	// 3. 比对
	enBase64 := base64.StdEncoding.EncodeToString(enDecode[0:keyLen])
	pwdEnBase64 := base64.StdEncoding.EncodeToString(encryptPwdWithSalt([]byte(pwd), salt))
	ok = enBase64 == pwdEnBase64

	return
}
