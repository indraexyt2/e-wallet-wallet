package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"e-wallet-wallet/constants"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	strPayload := ``

	endpoint := `/wallet/v1/ex/1/unlink`
	timestamp := `2025-01-18T21:47:02+07:00`
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	strPayload = re.ReplaceAllString(strPayload, "")
	strPayload = strings.ToLower(strPayload) + timestamp + endpoint

	fmt.Println("strPayload: ", strPayload)
	fmt.Println("endpoint: ", endpoint)
	fmt.Println("timestamp: ", timestamp)

	secretKey := constants.MappingClient["e-commerce"]
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(strPayload))
	generateSignature := hex.EncodeToString(h.Sum(nil))

	fmt.Println(generateSignature)
}
