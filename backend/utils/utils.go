package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateClient() {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	key := hex.EncodeToString(bytes)
	fmt.Println(key)
}
