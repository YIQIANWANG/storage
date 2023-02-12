package service

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"storage/conf"
	"storage/data"
	"strings"
)

type ChunkService struct {
}

func NewChunkService() *ChunkService {
	return &ChunkService{}
}

func (cs *ChunkService) PutChunk(chunkID string, chunkData []byte) error {
	if !PathExists(conf.ChunkFilePath) {
		err := os.MkdirAll(conf.ChunkFilePath, os.ModePerm)
		if err != nil {
			log.Println("Create Chunk Directory Failed: ", err)
			return err
		}
	}
	if PathExists(conf.ChunkFilePath + "/" + chunkID) {
		return nil
	}

	err := WriteFileByBytes(conf.ChunkFilePath+"/"+chunkID, chunkData)
	if err != nil {
		log.Println("Write Chunk Failed: ", err)
		return err
	}

	data.Lock.Lock()
	data.OperationCount++
	data.Lock.Unlock()
	_, _ = data.OpLogFile.WriteString("+ " + chunkID + "\n")

	return nil
}

func (cs *ChunkService) GetChunk(chunkID string) ([]byte, error) {
	chunkData, err := ReadFileByBytes(conf.ChunkFilePath + "/" + chunkID)
	if err != nil {
		log.Println("Read Chunk Failed: ", err)
		return nil, err
	}

	return chunkData, nil
}

func (cs *ChunkService) DelChunk(chunkID string) (int, error) {
	if !PathExists(conf.ChunkFilePath + "/" + chunkID) {
		return 0, nil
	}

	size, err := ComputeFileSize(conf.ChunkFilePath + "/" + chunkID)
	if err != nil {
		log.Println("Delete Chunk Failed: ", err)
		return -1, err
	}
	err = os.Remove(conf.ChunkFilePath + "/" + chunkID)
	if err != nil {
		log.Println("Delete Chunk Failed: ", err)
		return -1, err
	}

	data.Lock.Lock()
	data.OperationCount++
	data.Lock.Unlock()
	_, _ = data.OpLogFile.WriteString("- " + chunkID + "\n")

	return size, nil
}

func (cs *ChunkService) GetChunkIDs() ([]byte, error) {
	oplog, err := os.Open(conf.LogFilePath + "/" + conf.OpLogName)
	if err != nil {
		log.Println("Open OpLog Failed: ", err)
		return nil, err
	}

	chunkIDSet := make(map[string]bool)
	fileScanner := bufio.NewScanner(oplog)
	var rest int64
	data.Lock.RLock()
	rest = data.OperationCount
	data.Lock.RUnlock()
	for fileScanner.Scan() && rest > 0 {
		singleOperation := fileScanner.Text()
		operation := strings.Split(singleOperation, " ")
		if operation[0] == "+" {
			chunkIDSet[operation[1]] = true
		} else {
			chunkIDSet[operation[1]] = false
		}
		rest--
	}
	chunkIDs := make([]string, 0)
	for k, v := range chunkIDSet {
		if v {
			chunkIDs = append(chunkIDs, k)
		}
	}
	respData, _ := json.Marshal(chunkIDs)
	
	return respData, nil
}
