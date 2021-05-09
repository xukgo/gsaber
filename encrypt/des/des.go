package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"strings"
)

//const kv = "ae2d321c"
//const iv = "c3d0fd39"


func Encrypt(origData, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted,nil
	//return base64.StdEncoding.EncodeToString(crypted), nil
}

func Decrypt(crypted, key, iv []byte) ([]byte, error) {
	//crypted, err := base64.StdEncoding.DecodeString(cryptedStr)
	//if err != nil {
	//	return nil, err
	//}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/// <summary>
/// 格式化语句
/// </summary>
/// <param name="sql">需要格式化的字符</param>
/// <param name="direction">格式化的方向，false＝格式化 true＝去格式化</param>
/// <returns>返回处理后的字符</returns>
func FormatString(str string, direction bool) string {
	if str == "" {
		return ""
	}

	if !direction {
		str = strings.ReplaceAll(str, "<", "~xyu")
		str = strings.ReplaceAll(str, ">", "~dyu")
		str = strings.ReplaceAll(str, " ", "~kge")
		str = strings.ReplaceAll(str, "'", "~dyi")

		str = strings.ReplaceAll(str, "\"", "~syi")
		str = strings.ReplaceAll(str, "=", "~dey")
		str = strings.ReplaceAll(str, ",", "~dha")
		str = strings.ReplaceAll(str, "*", "~xha")

		str = strings.ReplaceAll(str, ".", "~ddi")
		str = strings.ReplaceAll(str, "/", "~fxg")
		str = strings.ReplaceAll(str, "+", "~jhz")
		str = strings.ReplaceAll(str, "-", "~jhj")
		str = strings.ReplaceAll(str, "_", "~xhx")
	} else {
		str = strings.ReplaceAll(str, "~xyu", "<")
		str = strings.ReplaceAll(str, "~dyu", ">")
		str = strings.ReplaceAll(str, "~kge", " ")
		str = strings.ReplaceAll(str, "~dyi", "'")

		str = strings.ReplaceAll(str, "~syi", "\"")
		str = strings.ReplaceAll(str, "~dey", "=")
		str = strings.ReplaceAll(str, "~dha", ",")
		str = strings.ReplaceAll(str, "~xha", "*")

		str = strings.ReplaceAll(str, "~ddi", ".")
		str = strings.ReplaceAll(str, "~fxg", "/")
		str = strings.ReplaceAll(str, "~jhz", "+")
		str = strings.ReplaceAll(str, "~jhj", "-")
		str = strings.ReplaceAll(str, "~xhx", "_")
	}
	return str
}