package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
)

func CalculateMD5HashFormFile(file multipart.File) (string, error) {
	hash := md5.New()

	var resultMD5String string
	if _, err := io.Copy(hash, file); err != nil {
		return resultMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	resultMD5String = hex.EncodeToString(hashInBytes)
	return resultMD5String, nil
}
