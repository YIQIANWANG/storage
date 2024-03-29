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
	if PathExists(conf.ChunkFilePath + "/" + chunkID) {
		return nil
	}

	err := WriteFileByBytes(conf.ChunkFilePath+"/"+chunkID, chunkData)
	if err != nil {
		log.Println("PutChunk Failed: ", err)
		return err
	}
	data.AvailableCapLock.Lock()
	data.AvailableCap -= len(chunkData)
	data.AvailableCapLock.Unlock()
	data.OpCountLock.Lock()
	data.OpCount++
	data.OpCountLock.Unlock()
	_, _ = data.OpLog.WriteString("+ " + chunkID + "\n")

	return nil
}

func (cs *ChunkService) GetChunk(chunkID string) ([]byte, error) {
	chunkData, err := ReadFileByBytes(conf.ChunkFilePath + "/" + chunkID)
	if err != nil {
		log.Println("GetChunk Failed: ", err)
		return nil, err
	}

	return chunkData, nil
}

func (cs *ChunkService) DelChunk(chunkID string) error {
	if !PathExists(conf.ChunkFilePath + "/" + chunkID) {
		return nil
	}

	size, err := ComputeFileSize(conf.ChunkFilePath + "/" + chunkID)
	if err != nil {
		log.Println("DelChunk Failed: ", err)
		return err
	}
	err = os.Remove(conf.ChunkFilePath + "/" + chunkID)
	if err != nil {
		log.Println("DelChunk Failed: ", err)
		return err
	}
	data.AvailableCapLock.Lock()
	data.AvailableCap += size
	data.AvailableCapLock.Unlock()
	data.OpCountLock.Lock()
	data.OpCount++
	data.OpCountLock.Unlock()
	_, _ = data.OpLog.WriteString("- " + chunkID + "\n")

	return nil
}

func (cs *ChunkService) GetChunkIDs() ([]byte, error) {
	oplog, err := os.Open(conf.LogFilePath + "/" + conf.OpLogName)
	if err != nil {
		log.Println("GetChunkIDs Failed: ", err)
		return nil, err
	}

	chunkIDSet := make(map[string]bool)
	fileScanner := bufio.NewScanner(oplog)
	var rest int64
	data.OpCountLock.Lock()
	rest = data.OpCount
	data.OpCountLock.Unlock()
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
