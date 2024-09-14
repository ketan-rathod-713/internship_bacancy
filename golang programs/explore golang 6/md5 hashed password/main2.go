package main

// import (
// 	"crypto/md5"
// 	"fmt"
// 	"strings"
// )

// const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// // b64From24Bit converts a 24-bit integer to a Base64 string
// func b64From24Bit(b2, b1, b0 byte, n int) string {
// 	w := uint32(b2)<<16 | uint32(b1)<<8 | uint32(b0)
// 	var result strings.Builder
// 	for n > 0 {
// 		n--
// 		result.WriteByte(itoa64[w&0x3f])
// 		w >>= 6
// 	}
// 	return result.String()
// }

// // md5Crypt is the main function that performs the MD5 crypt algorithm
// func md5Crypt(password, salt string) string {
// 	// take initial prefix for md5 crypt algorithm
// 	const magic = "$1$"

// 	// Truncate salt to 8 characters
// 	if len(salt) > 8 {
// 		salt = salt[:8]
// 	}

// 	// Initial MD5
// 	ctx := md5.New()
// 	ctx.Write([]byte(password))
// 	ctx.Write([]byte(magic))
// 	ctx.Write([]byte(salt))

// 	// Another MD5
// 	altCtx := md5.New()
// 	altCtx.Write([]byte(password))
// 	altCtx.Write([]byte(salt))
// 	altCtx.Write([]byte(password))
// 	final := altCtx.Sum(nil)

// 	// what are we doing here ?
// 	for i := len(password); i > 0; i -= 16 {
// 		if i > 16 {
// 			ctx.Write(final[:16])
// 		} else {
// 			ctx.Write(final[:i])
// 		}
// 	}

// 	for i := len(password); i > 0; i >>= 1 {
// 		if i&1 != 0 {
// 			ctx.Write([]byte{0})
// 		} else {
// 			ctx.Write([]byte{password[0]})
// 		}
// 	}

// 	final = ctx.Sum(nil)

// 	for i := 0; i < 1000; i++ {
// 		altCtx.Reset()

// 		if i&1 != 0 {
// 			altCtx.Write([]byte(password))
// 		} else {
// 			altCtx.Write(final[:])
// 		}

// 		if i%3 != 0 {
// 			altCtx.Write([]byte(salt))
// 		}

// 		if i%7 != 0 {
// 			altCtx.Write([]byte(password))
// 		}

// 		if i&1 != 0 {
// 			altCtx.Write(final[:])
// 		} else {
// 			altCtx.Write([]byte(password))
// 		}

// 		final = altCtx.Sum(nil)
// 	}

// 	// Generate the output string
// 	hash := magic + salt + "$"
// 	hash += b64From24Bit(final[0], final[6], final[12], 4)
// 	hash += b64From24Bit(final[1], final[7], final[13], 4)
// 	hash += b64From24Bit(final[2], final[8], final[14], 4)
// 	hash += b64From24Bit(final[3], final[9], final[15], 4)
// 	hash += b64From24Bit(final[4], final[10], final[5], 4)
// 	hash += b64From24Bit(0, 0, final[11], 2)

// 	return hash
// }

// func main() {
// 	// Example usage
// 	password := "tridip123@123"
// 	salt := "MPzeqLP8"

// 	hashedPassword := md5Crypt(password, salt)
// 	fmt.Println("Hashed Password:", hashedPassword)

// 	// To check the password
// 	isValid := (hashedPassword == md5Crypt(password, salt))
// 	fmt.Println("Password Valid:", isValid)
// }
