package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CompleteJSON struct {
	Color string      `json:"color"`
	Name  string      `json:"name"`
	Types []TypesJSON `json:"types"`
	Image string      `json:"image"`
}

type TypesJSON struct {
	Slot int      `json:"slot"`
	Type TypeJSON `json:"type"`
}

type TypeJSON struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type DetailPokemonJSON struct {
	Types   []TypesJSON `json:"types"`
	Sprites Sprites     `json:"sprites"`
	Species Species     `json:"species"`
}

type Species struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonSpecies struct {
	Color Color `json:"color"`
}

type Color struct {
	Name string `json:"name"`
}

type Sprites struct {
	Other Other `json:"other"`
}
type Other struct {
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
}
type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
}

type GetPokemonsJSON struct {
	Results []ResultsGetPokemonsJSON `json:"results"`
}

type ResultsGetPokemonsJSON struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func RespBodyToString(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func GetPokemon(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=10")
	if err != nil {
		return
	}

	GetPokemonsString := RespBodyToString(resp)
	// fmt.Println(GetPokemonsString)
	var getpokemonjson GetPokemonsJSON

	json.Unmarshal([]byte(GetPokemonsString), &getpokemonjson)
	// fmt.Println(getpokemonjson)

	ResponseJSON := []CompleteJSON{}

	for _, v := range getpokemonjson.Results {
		cJSON := CompleteJSON{
			Name: v.Name,
		}

		resp, err := http.Get(v.Url)
		if err != nil {

		}
		DetailPokemonString := RespBodyToString(resp)
		var detailpokemonjson DetailPokemonJSON
		json.Unmarshal([]byte(DetailPokemonString), &detailpokemonjson)
		cJSON.Types = detailpokemonjson.Types
		cJSON.Image = detailpokemonjson.Sprites.Other.OfficialArtwork.FrontDefault
		// fmt.Printf("cJSON.Image: %v\n", cJSON.Image)
		speciesUrl := detailpokemonjson.Species.Url
		// fmt.Println(speciesUrl)

		resp, err = http.Get(speciesUrl)
		SpeciesPokemonString := RespBodyToString(resp)
		var speciespokemonjson PokemonSpecies
		json.Unmarshal([]byte(SpeciesPokemonString), &speciespokemonjson)
		// fmt.Println(speciespokemonjson)
		cJSON.Color = speciespokemonjson.Color.Name

		ResponseJSON = append(ResponseJSON, cJSON)

		// fmt.Println(ResponseJSON)
	}
	// fmt.Println(ResponseJSON)
	json.NewEncoder(rw).Encode(ResponseJSON)
	// rw.Write([]byte("ok"))
}
