package main

import (
	"encoding/base64"
	"log"
)

const encryptKey = "your-encryption-key"

func encryptDecrypt(input, key string) string {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return string(output)
}

func encrypt(input string) string {
	encrypted := encryptDecrypt(input, encryptKey)
	return base64.StdEncoding.EncodeToString([]byte(encrypted))
}

func decrypt(input string) string {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	return encryptDecrypt(string(decoded), encryptKey)
}
