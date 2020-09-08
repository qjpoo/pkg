package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {

	fmt.Println(md5Str("woGo"))
	fmt.Println(sha512Str("woGo"))
	fmt.Println(base64DecodeStr(base64EncodeStr("fd")))
}

//md5验证
func md5Str(src string) string {

	h := md5.New()
	h.Write([]byte(src)) //
	//fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

//sha512验证
func sha512Str(src string) string {
	h := sha512.New()
	h.Write([]byte(src)) //
	//fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

//base编码
func base64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

//base解码
func base64DecodeStr(src string) string {
	a, err := (base64.StdEncoding.DecodeString(src))
	if err != nil {
		return "error"
	}
	return string(a)
}
