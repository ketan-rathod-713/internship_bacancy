package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Println("ERROR", err)
		return // Exit if there's an error
	}

	defer conn.Close()

	ans, err := redis.String(conn.Do("GET", "name")) // Use redis.String to convert the result to string
	if err != nil {
		log.Println("ERROR", err)
		return // Exit if there's an error
	}

	value := `{"Id":"1710147856178495379","SalesPersonName":"Ketan Rathod","EmailId":"ketanrtd1@gmail.com","Password":"ijUqABuu_qh8V8bp99Xjgpbo9g==","PhoneNo":"909-972-6655","DealershipName":"Tridip Motors","ConnectedDevice":null,"AppleIdentifier":"","GoogleIdentifier":"","FacebookIdentifier":"","IsDeleted":false,"IsVerified":false,"CreatedDate":"03/11/2024 02:04:16 AM","UpdatedDate":"03/11/2024 02:04:16 AM","IsTFA":false,"TFACode":"","IsXMLADF":false,"ADFEmail":""}`
	hash := "salespersoninfo"
	key := 1234
	_, err = conn.Do("hset",hash, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
