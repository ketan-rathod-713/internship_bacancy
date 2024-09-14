package main

// import (
// 	"crypto/md5"
// 	"encoding/hex"
// 	"fmt"
// )

// func main() {
// 	fmt.Println("password hashing using md5 and checking if it works for sftp")

// 	password := "12345"
// 	hashedPassword, err := hashPassword(password)
// 	fmt.Println("hashed password", hashedPassword, "error", err)

// 	fmt.Println("check password hashes")
// 	if checkPasswordHash(password, hashedPassword) {
// 		fmt.Println("Password is correct")
// 	} else {
// 		fmt.Println("Password is incorrect")
// 	}
// }

// // hashPassword generates an MD5 hash of the given password and appends a $1$ prefix.
// func hashPassword(password string) (string, error) {
// 	hash := md5.New()
// 	_, err := hash.Write([]byte(password))
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("$1$%s", hex.EncodeToString(hash.Sum(nil))), nil
// }

// // checkPasswordHash compares the MD5 hash of the given password with the provided hash.
// func checkPasswordHash(password, hash string) bool {
// 	if len(hash) < 4 || hash[:4] != "$1$" {
// 		// The hash doesn't start with the $1$ prefix, so it's not a valid MD5 hash
// 		return false
// 	}
// 	hashedPassword := md5.New()
// 	hashedPassword.Write([]byte(password))
// 	return fmt.Sprintf("$1$%s", hex.EncodeToString(hashedPassword.Sum(nil))) == hash
// }
