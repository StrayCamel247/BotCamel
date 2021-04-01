package utils

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	// url2 "net/url"
	"io"
	"os"
	"time"
)

// ReadFile 读取文件
// 读取失败返回 nil
func ReadFile(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.WithError(err).WithField("util", "ReadFile").Errorf("unable to read '%s'", path)
		return nil
	}
	return bytes
}

// 文件下载
func DownloadImg(filename, url string) error {
	// 记录下载时间
	_nowTime := time.Now()
	_logHandler := func(start time.Time) {
		tc := time.Since(start)
		log.Info(fmt.Sprintf("time cost = %v\n", tc))
	}
	defer _logHandler(_nowTime)
	// 构造请求头
	spaceClient := http.Client{
		// 请求时间
		Timeout: time.Minute * 10,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Warn(err)
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Warn(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		fmt.Println("图片下载失败；url")
		log.WithError(err)
		return err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(filename)
	if err != nil {
		log.WithError(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
	return nil
}

// 检查文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("File reading error", err)
	return false
}
