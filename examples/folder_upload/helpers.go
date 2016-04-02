package main

import (
	"fmt"
	"github.com/masci/flickr"
	"github.com/masci/flickr/photosets"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var image_extensions = map[string]bool{"jpg": true, "gif": true, "png": true, "jpeg": true}

func getFilePaths(base_path string) []os.FileInfo {
	var only_files []os.FileInfo
	files, _ := ioutil.ReadDir(base_path)
	for _, f := range files {
		fmt.Println(f.Name())
		if f.IsDir() == false && isImage(f) {
			only_files = append(only_files, f)
		}
	}
	return only_files
}

func isImage(fileinfo os.FileInfo) bool {
	ext := filepath.Ext(fileinfo.Name())
	fmt.Println("Extension is", ext)
	lowerStr := strings.ToLower(ext)

	if image_extensions[lowerStr[1:len(lowerStr)]] == true {
		return true
	} else {
		return false
	}
}

func uploadImageAndCreateSet(base_path string, fileinfo os.FileInfo, client *flickr.FlickrClient) (*photosets.PhotosetResponse, error) {
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

	return photosets.Create(client, filepath.Base(base_path), "", resp.Id)
}

func uploadImageToSet(base_path string, fileinfo os.FileInfo, client *flickr.FlickrClient, setResponse *photosets.PhotosetResponse) (*flickr.BasicResponse, error) {
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

	return photosets.AddPhoto(client, setResponse.Set.Id, resp.Id)
}
