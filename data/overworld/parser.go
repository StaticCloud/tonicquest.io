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
	Name       string      `json:"name"`
	Map        [][]string  `json:"map"`
	Savepoints []Savepoint `json:"savepoints"`
	Enemies    []Enemy     `json:"enemies"`
	Allies     []Ally      `json:"allies"`
}

type Enemy struct {
	ID          int         `json:"_id"`
	Name        string      `json:"name"`
	Sprite      string      `json:"sprite"`
	Health      int         `json:"health"`
	Moves       []string    `json:"moves"`
	Defeated    bool        `json:"defeated"`
	Coordinates Coordinates `json:"coordinates"`
}

type Ally struct {
	Name        string      `json:"name"`
	Sprite      string      `json:"sprite"`
	Coordinates Coordinates `json:"coordinates"`
	Dialogue    []Dialogue  `json:"dialogue"`
}

type Savepoint struct {
	ID          int         `json:"_id"`
	Coordinates Coordinates `json:"coordinates"`
}

type Dialogue struct {
	Speaker int    `json:"speaker"`
	Text    string `json:"text"`
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

func SaveWorld(data *Data) (*Data, error) {
	file, err := os.Create("./data/overworld/overworld.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}

	newData, err := Unmarshall()
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func ResetWorld(data *Data) (*Data, error) {
	for i := range data.Locations[0].Enemies {
		data.Locations[0].Enemies[i].Defeated = false
	}

	return SaveWorld(data)
}
