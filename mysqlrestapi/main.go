package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6
func main() {
	cipher := encrypt([]byte("umasenhae"), "marcos")
	fmt.Println("encrypted", cipher)
	playnText := decrypt(cipher, "marcos")
	fmt.Println("decrypted", playnText)

	var a = App{}
	a.Initialize("root", "", "rest_api_example")
	a.Run("localhost:8081")

	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/rest_api_example")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	// // insert, err := db.Query("INSERT INTO users VALUES ( 200, 'Marcos', 30 )")
	// del, err := db.Query("DELETE FROM users")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// // defer insert.Close()
	// defer del.Close()

}

// key
func createHash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	cipherBlock, _ := aes.NewCipher([]byte(createHash("migormiguxo")))
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize()) // slice
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return plainText
}
