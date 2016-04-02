package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func uploadImageAndCreateSet(fileinfo os.FileInfo) flickr.photoset, Error {
	params := flickr.NewUploadParams()
	resp, err := flickr.UploadFile(client, path, params)
	if err != nil {
		fmt.Println("Failed uploading:", err)
		if resp != nil {
			fmt.Println(resp.ErrorMsg)
		}
	}
	else {
		fmt.Println("Photo uploaded:", filepath, resp.Id)
	}

	respS, err := photosets.Create(client, "")

} 
