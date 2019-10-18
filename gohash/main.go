package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	pass := "marcos13231989"
	md5HashInBytes := md5.Sum([]byte(pass))
	md5HashInString := hex.EncodeToString(md5HashInBytes[:])
	y := sha256.Sum224([]byte(pass))
	x := sha256.Sum256([]byte(pass))
	fmt.Println(md5HashInString)
	fmt.Println(md5HashInBytes)
	fmt.Println(y)
	fmt.Println(x)
}
