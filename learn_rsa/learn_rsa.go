package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)  
  
func main() {  
    // 生成一对 RSA 密钥对  
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)  
    if err != nil {  
        panic(err)  
    }  
    publicKey := &privateKey.PublicKey  
  
    // 要签名的消息  
    message := []byte("Hello, world!")  
  
    // 使用 SHA256 作为哈希算法，对消息进行哈希  
    hash := sha256.Sum256(message)  
  
    // 对哈希值进行签名  
    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])  
    if err != nil {  
        panic(err)  
    }  
  
    // 验证签名是否正确  
    if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature); err == nil {  
        fmt.Println("Signature is valid")  
    } else {  
        fmt.Println("Signature is invalid")  
    }  
}