package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Info struct {
	Id         int      `json:"id"`
	Synopsis   string   `json:"synopsis"`
	Creators   []string `json:"creators"`
	Genres     []string `json:"genres"`
	YearsAired string   `json:"yearsAired"`
}

type Character struct {
	Id                       int                      `json:"id"`
	Name                     string                   `json:"name"`
	Image                    string                   `json:"image"`
	Bio                      string                   `json:"bio"`
	Nationality              string                   `json:"nationality"`
	Ethnicity                string                   `json:"ethnicity"`
	Ages                     []int                    `json:"ages"`
	Born                     string                   `json:"born"`
	Died                     []int                    `json:"died"`
	PhysicalDescription      PhysicalDescription      `json:"physicalDescription"`
	PersonalInformation      PersonalInformation      `json:"personalInformation"`
	PoliticalInformation     PoliticalInformation     `json:"politicalInformation"`
	ChronologicalInformation ChronologicalInformation `json:"chronologicalInformation"`
}

type PhysicalDescription struct {
	Gender    string `json:"gender"`
	EyeColor  string `json:"eyeColor"`
	HairColor string `json:"hairColor"`
	SkinColor string `json:"skinColor"`
}

type PersonalInformation struct {
	LoveInterest   string   `json:"loveInterst"`
	Allies         []string `json:"allies"`
	Enemies        []string `json:"enemies"`
	Weapons        []string `json:"weaponsOfChoice"`
	FightingStyles []string `json:"fightingStyles"`
}

type PoliticalInformation struct {
	Profession   []string `json:"profession"`
	Position     []string `json:"position"`
	Predecessor  string   `json:"predecessor"`
	Successor    string   `json:"successor"`
	Affiliations []string `json:"affiliations"`
}

type ChronologicalInformation struct {
	FirstAppearance string   `json:"firstAppearance"`
	LastAppearance  []string `json:"lastAppearance"`
	VoicedBy        []string `json:"voicedBy"`
}

type Episode struct {
	Id             int    `json:"id"`
	Season         string `json:"Season"`
	Episode        string `json:"NumInSeason"`
	Title          string `json:"Title"`
	AnimatedBy     string `json:"AnimatedBy"`
	TirectedBy     string `json:"DirectedBy"`
	WrittenBy      string `json:"WrittenBy"`
	AirDate        string `json:"OriginalAirDate"`
	ProductionCode int    `json:"ProductionCode"`
}

type Questions struct {
	Id              int      `json:"id"`
	Question        string   `json:"question"`
	PossibleAnswers []string `json:"possibleAnsers"`
	CorrectAnswer   string   `json:"correctAnswer"`
}

func ApiData() []Character {
	charURL := "https://api.sampleapis.com/avatar/characters"
	charResponse, err := http.Get(charURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	charResponseData, err := ioutil.ReadAll(charResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var charObject []Character
	json.Unmarshal(charResponseData, &charObject)

	// for key, _ := range charObject {
	// 	fmt.Println(charObject[key].Name)
	// 	fmt.Println("")
	// 	for _, value := range charObject[key].PoliticalInformation.Position {
	// 		fmt.Println(value)
	// 	}
	// 	fmt.Println("")
	// }

	return charObject
}
