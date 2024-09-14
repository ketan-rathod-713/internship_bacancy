package main

import (
	"fmt"
	md5_crypt "password_hashing_example/md5crypt"
)

func main() {
	password := "123456"
	salt := "abcdefghijklmnopqrstuvwxyzABCDEF"

	hashedpassword, err := md5_crypt.MD5Crypt(password, salt)
	fmt.Println("hashedpassword", hashedpassword, "err", err)

	correct, err := md5_crypt.CompareHash(password, hashedpassword)
	fmt.Println("correct", correct, "error", err)
}
