package spotify_service

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"net/http"

	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	"github.com/zmb3/spotify"
)

func GetSearchType(req_search_type spotifyv1.SearchType) (spotify.SearchType, error) {
	switch req_search_type {
	case spotifyv1.SearchType_SEARCH_TYPE_ALBUM:
		return spotify.SearchTypeAlbum, nil
	case spotifyv1.SearchType_SEARCH_TYPE_ARTIST:
		return spotify.SearchTypeArtist, nil
	}
	return spotify.SearchTypeAlbum, errors.New("bad search type")
}

func GetImageUrl(req_search_type spotifyv1.SearchType, result spotify.SearchResult) (string, error) {
	switch req_search_type {
	case spotifyv1.SearchType_SEARCH_TYPE_ALBUM:
		results := result.Albums
		if len(results.Albums) == 0 || len(results.Albums[0].Images) == 0 {
			return "", errors.New("bad search query")
		}
		return results.Albums[0].Images[0].URL, nil
	case spotifyv1.SearchType_SEARCH_TYPE_ARTIST:
		results := result.Artists
		if len(results.Artists) == 0 || len(results.Artists[0].Images) == 0 {
			return "", errors.New("bad search query")
		}
		return results.Artists[0].Images[0].URL, nil
	}
	return "", errors.New("bad search query")
}
func ImageURLToImage(image_url string) (image.Image, error) {
	resp, err := http.Get(image_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("bad http status code for image retrieval")
	}
	image, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return image, nil
}
func ImageToBytes(image image.Image) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := jpeg.Encode(buff, image, nil)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
