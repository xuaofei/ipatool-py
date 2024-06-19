package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func requestTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("requestTaskHandler in")
	defer log.Printf("requestTaskHandler out")

	request := taskInfoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	response := taskInfoResponse{}
	response.AppleID = "fsi2clsfiz8@163.com"
	response.ApplePwd = "Ls112211"
	response.AppBundleID = "jp.eure.pairs"
	response.Country = "JP"
	response.TaskID = fmt.Sprintf("%v", time.Now())

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("requestTaskHandler failed, Marshal taskInfo err:%v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("requestTaskHandler failed, Write err:%v", err)
		return
	}
}

func request2FAHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request2FAHandler in")
	defer log.Printf("request2FAHandler out")

	request := twoFAInfoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	log.Printf("request2FAHandler AppleID:%v ,TaskID:%v", request.AppleID, request.TaskID)

	filePath := "C:\\Users\\xuaofei\\Desktop\\AppStoreHistoryIpa\\db.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("request2FAHandler failed ,读取文件失败：%v", err)
		return
	}

	log.Printf("request2FAHandler read file content:%v", string(content))

	response := twoFAInfoResponse{}
	response.TwoFACode = string(content)

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("request2FAHandler failed, Marshal taskInfo err:%v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("request2FAHandler failed, Write err:%v", err)
		return
	}
}

// 上传ipa版本信息
func uploadVersionsInfoHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	request := ipaVersionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	log.Printf("uploadVersionsInfoHandler AppleID:%v, TaskID:%v", request.AppleID, request.TaskID)

	filePath := "C:\\Users\\xuaofei\\Desktop\\AppStoreOldVersion\\db.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("request2FAHandler failed ,读取文件失败：%v", err)
		return
	}

	fmt.Println(string(content))

	_, err = w.Write([]byte(content))
	if err != nil {
		log.Printf("request2FAHandler failed, Write err:%v", err)
		return
	}
}

// 查询需要下载的版本信息
func requestDownloadVersionsInfo(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	request := downloadIpaVersionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	log.Printf("uploadVersionsInfoHandler AppleID:%v, TaskID:%v", request.AppleID, request.TaskID)

	filePath := "C:\\Users\\xuaofei\\Desktop\\AppStoreOldVersion\\db.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("request2FAHandler failed ,读取文件失败：%v", err)
		return
	}

	fmt.Println(string(content))

	_, err = w.Write([]byte(content))
	if err != nil {
		log.Printf("request2FAHandler failed, Write err:%v", err)
		return
	}

}
