package service

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"net/url"
	"os"
	"storage/conf"
	"storage/data"
	"strings"
)

type SyncService struct {
}

func NewSyncService() *SyncService {
	return &SyncService{}
}

func (ss *SyncService) SyncChunk(dstAddress, chunkID string) error {
	chunkData, err := ReadFileByBytes(conf.ChunkFilePath + "/" + chunkID)
	// 已经被删了就不用同步了
	if err != nil {
		return nil
	}
	reqUrl, _ := url.Parse("http://" + dstAddress + "/chunk")
	params := url.Values{}
	params.Set("chunkID", chunkID)
	reqUrl.RawQuery = params.Encode()

	// 错误收敛于单个Storage
	_, _ = http.Post(reqUrl.String(), "text/plain", bytes.NewReader(chunkData))
	return nil
}

func (ss *SyncService) SyncAll(dstAddress string) error {
	// 避免重复同步
	if data.DstStorageAddress[dstAddress] {
		return nil
	}
	data.DstStorageAddress[dstAddress] = true
	// 异步返回
	go func() {
		oplog, err := os.Open(conf.LogFilePath + "/" + conf.OpLogName)
		if err != nil {
			log.Println("Open OpLog Failed: ", err)
			return
		}
		fileScanner := bufio.NewScanner(oplog)
		var rest int64
		data.Lock.RLock()
		rest = data.OperationCount
		data.Lock.RUnlock()
		for fileScanner.Scan() && rest > 0 {
			singleOperation := fileScanner.Text()
			operation := strings.Split(singleOperation, " ")
			if operation[0] == "+" {
				chunkData, err := ReadFileByBytes(conf.ChunkFilePath + "/" + operation[1])
				if err != nil {
					continue
				}
				reqUrl, _ := url.Parse("http://" + dstAddress + "/chunk")
				params := url.Values{}
				params.Set("chunkID", operation[1])
				reqUrl.RawQuery = params.Encode()

				// 错误收敛于单个Storage
				_, _ = http.Post(reqUrl.String(), "text/plain", bytes.NewReader(chunkData))
			} else {
				reqUrl, _ := url.Parse("http://" + dstAddress + "/chunk")
				params := url.Values{}
				params.Set("chunkID", operation[1])
				reqUrl.RawQuery = params.Encode()
				req, _ := http.NewRequest(http.MethodDelete, reqUrl.String(), nil)

				// 错误收敛于单个Storage
				_, _ = http.DefaultClient.Do(req)
			}
			rest--
		}
		delete(data.DstStorageAddress, dstAddress)
	}()
	return nil
}
