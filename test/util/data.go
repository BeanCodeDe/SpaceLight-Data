package util

import "crypto/rand"

const (
	Url             = "http://localhost:1203"
	CorrelationId   = "X-Correlation-ID"
	ContentTypValue = "application/json; charset=utf-8"
	ContentTyp      = "Content-Type"
)

const alphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandomString(length int) string {
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphaNum[b%byte(len(alphaNum))]
	}
	return string(bytes)
}
