package api

import (
	"encoding/json"
	"groupie-tracker/structures"
	"groupie-tracker/utils"
	"net/http"
)

const api = "https://groupietrackers.herokuapp.com/api/artists"

func GetArtists() ([]structures.Artist, error) {
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body) ///
	defer resp.Body.Close()               ///
	artists := []structures.Artist{}
	err = decoder.Decode(&artists) ///
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func GetArtistDetails(id string) (structures.Artist, error) {
	artistEndpoint := api + "/" + id
	resp, err := http.Get(artistEndpoint)
	if err != nil {
		return structures.Artist{}, err
	}
	decoder := json.NewDecoder(resp.Body)
	artist := structures.Artist{}
	err = decoder.Decode(&artist)
	if err != nil {
		return structures.Artist{}, err
	}
	populatedArtist, err := utils.ExtractEvents(artist)
	if err != nil {
		return structures.Artist{}, err
	}
	return populatedArtist, nil
}
