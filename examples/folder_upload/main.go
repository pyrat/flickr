package main

import (
	"fmt"
	"github.com/masci/flickr"
	"os"
)

func main() {
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
	base_path := "/Users/alastairbrunton/Pictures/flickr/testupload"

	filepaths := getFilePaths(base_path)

	// Run a shift on the filepaths slice
	first_image, filepaths := filepaths[0], filepaths[1:]
	photoset, err := uploadImageAndCreateSet(base_path, first_image, client)
	if err != nil {
		fmt.Println("Error uploading first photo and creating photoset", err)
	}

	for _, f := range filepaths {
		_, err := uploadImageToSet(base_path, f, client, photoset)
		if err != nil {
			fmt.Println("Error uploading image", f.Name())
		}
	}
}
