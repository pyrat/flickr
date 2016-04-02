package main

import (
	"fmt"
	"github.com/pyrat/flickr"
	"github.com/pyrat/flickr/photosets"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getFilePaths(base_path string) []os.FileInfo {
	var only_files []os.FileInfo
	files, _ := ioutil.ReadDir(base_path)
	for _, f := range files {
		if f.IsDir() == false {
			only_files = append(only_files, f)
		}
	}
	return only_files
}

func uploadImageAndCreateSet(base_path string, fileinfo os.FileInfo) (*flickr.UploadResponse, error) {
	params := flickr.NewUploadParams()
	path := base_path + "/" + fileinfo.Name()
	resp, err := flickr.UploadFile(client, path, params)
	if err != nil {
		fmt.Println("Failed uploading:", err)
		if resp != nil {
			fmt.Println(resp.ErrorMsg)
		}
	} else {
		fmt.Println("Photo uploaded:", path, resp.Id)
	}

	respS, err := photosets.Create(client, filepath.Base(base_path), resp.Id)
	return respS
}
