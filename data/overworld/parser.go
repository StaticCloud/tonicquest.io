package data

import (
	"encoding/json"
	"io"
	"os"
)

type Data struct {
	Locations []Location `json:"locations"`
}

type Location struct {
	Name    string     `json:"name"`
	Map     [][]string `json:"map"`
	Enemies []Enemy    `json:"enemies"`
}

type Enemy struct {
	Name        string      `json:"name"`
	Sprite      string      `json:"sprite"`
	Health      int         `json:"health"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

func Unmarshall() (*Data, error) {
	jsonFile, err := os.Open("./data/overworld/overworld.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, byteError := io.ReadAll(jsonFile)
	if byteError != nil {
		return nil, byteError
	}

	var data *Data

	unmarshallErr := json.Unmarshal(byteValue, &data)
	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	return data, nil
}

func LoadWorld() (*Data, error) {
	data, err := Unmarshall()
	if err != nil {
		return nil, err
	}

	return data, nil
}
