package utils

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func getCos() *cos.Client {
	u, _ := url.Parse("https://1037buqieryu-1317065983.cos.ap-shanghai.myqcloud.com") //获取客户端
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDa1HK8hnrjPsFTUI1cmjTCoDLMAk5C7qK",
			SecretKey: "PbGjy1qUxFMI1WnpeKrGVYN0Slsqf7Ra",
		},
	})
	return c
}

func Upload(prefix string, fileName string, file io.Reader) error {
	c := getCos()
	_, err := c.Object.Put(context.Background(), prefix+fileName, file, nil)
	if err != nil {
		return err
	}
	return nil
}

func UploadMutipy(prefix string, files []*multipart.FileHeader) {
	c := getCos()
	// 多线程批量上传文件
	filesCh := make(chan *multipart.FileHeader, 2)
	//filePaths := []string{"test1", "test2", "test3"}
	var wg sync.WaitGroup
	threadpool := 2
	for i := 0; i < threadpool; i++ {
		wg.Add(1)
		go Mutipy(prefix, &wg, c, filesCh)
	}
	for _, filePath := range files {
		filesCh <- filePath
	}
	close(filesCh)
	wg.Wait()
}

func Mutipy(prefix string, wg *sync.WaitGroup, c *cos.Client, fileCh <-chan *multipart.FileHeader) {
	defer wg.Done()
	for file := range fileCh {
		fd, err := file.Open()
		if err != nil {
			//ERROR
			continue
		}
		_, err = c.Object.Put(context.Background(), prefix+file.Filename, fd, nil)
		if err != nil {
			//ERROR
		}
		fd.Close()
	}
}
