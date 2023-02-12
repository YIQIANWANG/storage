package service

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func ComputeHash(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	hashStr := hex.EncodeToString(hash.Sum(nil))

	return hashStr
}

func ComputeFileSize(fileName string) (int, error) {
	file, err := os.Stat(fileName)
	if err != nil {
		return -1, err
	}

	return int(file.Size()), nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil || !os.IsNotExist(err) {
		return true
	}

	return false
}

func ReadFileByBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFileByBytes(fileName string, data []byte) error {
	err := os.WriteFile(fileName, data, os.ModePerm)
	return err
}
