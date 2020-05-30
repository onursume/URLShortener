package shorten

import (
	"github.com/matoous/go-nanoid"
)

const encodeStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const defaultByte = 8

func GenerateId() (string, error) {
	id, err := gonanoid.Generate(encodeStr, defaultByte)
	if err != nil {
		return "", err
	}
	return id, nil
}