package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/masci/flickr"
	"github.com/masci/flickr/photos"
	"github.com/masci/flickr/photosets"
)

func main() {
	var pause = func() {
		var foo string
		fmt.Println("Press a key to continue")
		fmt.Scanln(&foo)
	}

	// retrieve Flickr credentials from env vars
	apik := os.Getenv("FLICKRGO_API_KEY")
	apisec := os.Getenv("FLICKRGO_API_SECRET")
	token := os.Getenv("FLICKRGO_OAUTH_TOKEN")
	tokenSecret := os.Getenv("FLICKRGO_OAUTH_TOKEN_SECRET")

	// do not proceed if credentials were not provided
	if apik == "" || apisec == "" || token == "" || tokenSecret == "" {
		fmt.Fprintln(os.Stderr, "Please set FLICKRGO_API_KEY, FLICKRGO_API_SECRET "+
			", FLICKRGO_OAUTH_TOKEN and FLICKRGO_OAUTH_TOKEN_SECRET env vars")
		os.Exit(1)
	}

	// create an API client with credentials
	client := flickr.NewFlickrClient(apik, apisec)
	client.OAuthToken = token
	client.OAuthTokenSecret = tokenSecret

	// get a list of filepaths
	// upload first filepath (add some sort of progress meter here)
	// add to a set
	// for each remaining filepath (upload and add to set w/progress meter)

	filepaths := getFilePaths()

	// upload a photo
	path, _ := filepath.Abs("gopher.jpg")
	params := flickr.NewUploadParams()
	params.Title = "A Gopher"
	resp, err := flickr.UploadFile(client, path, params)
	if err != nil {
		fmt.Println("Failed uploading:", err)
		if resp != nil {
			fmt.Println(resp.ErrorMsg)
		}
		os.Exit(1)
	} else {
		fmt.Println("Photo uploaded, id:", resp.Id)
		pause()
	}

	// create a photoset using above photo as primary
	respS, err := photosets.Create(client, "GoFlickrSet", "", resp.Id)
	if err != nil {
		fmt.Println("Failed creating set:", respS.ErrorMsg())
		os.Exit(1)
	} else {
		fmt.Println("Set created, id:", respS.Set.Id, "url:", respS.Set.Url)
		pause()
	}

	// upload another photo using default params
	path, _ = filepath.Abs("gophers.jpg")
	resp, err = flickr.UploadFile(client, path, nil)
	if err != nil {
		fmt.Println("Failed uploading:", err)
		if resp != nil {
			fmt.Println(resp.ErrorMsg())
		}
		os.Exit(1)
	} else {
		fmt.Println("Photo uploaded, id:", resp.Id)
		pause()
	}

	// assign above photo to the photoset
	respAdd, err := photosets.AddPhoto(client, respS.Set.Id, resp.Id)
	if err != nil {
		fmt.Println("Failed adding photo to the set:", err, respAdd.ErrorMsg())
		os.Exit(1)
	} else {
		fmt.Println("Added photo", resp.Id, "to set", respS.Set.Id)
		pause()
	}
}
