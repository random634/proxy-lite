package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
)

type CryptoDES struct {
	PasswordMd5 []byte
}

func NewCryptoDES(password string) CryptoMethod {
	c := new(CryptoDES)
	sum := md5.Sum([]byte(password))
	c.PasswordMd5 = sum[0:8]
	return c
}

func (c *CryptoDES) Encrypt(origData []byte) ([]byte, error) {
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(c.PasswordMd5)
	//对明文先进行补码操作
	origData = c.getPKCS5Padding(origData, block.BlockSize())
	//设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, c.PasswordMd5)
	//创建明文长度的字节数组
	cryptedData := make([]byte, len(origData))
	//加密明文,加密后的数据放到数组中
	blockMode.CryptBlocks(cryptedData, origData)
	return cryptedData, nil
}

func (c *CryptoDES) Decrypt(cryptedData []byte) ([]byte, error) {
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(c.PasswordMd5)
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block, c.PasswordMd5)
	//创建密文大小的数组变量
	origData := make([]byte, len(cryptedData))
	//解密密文到数组origData中
	blockMode.CryptBlocks(origData, cryptedData)
	//去补码
	origData = c.getPKCS5UnPadding(origData)

	return origData, nil
}

//实现明文的补码
func (c *CryptoDES) getPKCS5Padding(origData []byte, blockSize int) []byte {
	//计算出需要补多少位
	padding := blockSize - len(origData)%blockSize
	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
	// 需要补padding位的padding值
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//把补充的内容拼接到明文后面
	return append(origData, padtext...)
}

//去除补码
func (c *CryptoDES) getPKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	//解密去补码时需取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文
	return origData[:(length - unpadding)]
}
