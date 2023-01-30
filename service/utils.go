package service

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
)

func ComputeHash(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	hashStr := hex.EncodeToString(hash.Sum(nil))
	return hashStr
}

func ReadFileByBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFileByBytes(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, os.ModePerm)
	return err
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil || !os.IsNotExist(err) {
		return true
	}
	return false
}
