package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type StructData struct {
	Data []Response `json:"Data"`
}

type Response struct {
	ConfirmDate    string      `json:"ConfirmDate"`
	No             interface{} `json:"No"`
	Age            interface{} `json:"Age"`
	Gender         string      `json:"Gender"`
	GenderEn       string      `json:"GenderEn"`
	Nation         interface{} `json:"Nation"`
	NationEn       string      `json:"NationEn"`
	Province       interface{} `json:"Province"`
	ProvinceId     int         `json:"ProvinceId"`
	District       interface{} `json:"District"`
	ProvinceEn     string      `json:"ProvinceEn"`
	StatQuarantine int         `json:"StatQuarantine"`
}

func FetchData(url string) (StructData, error) {

	resp, err := http.Get(url)
	if err != nil {
		return StructData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return StructData{}, err
	}

	bodyJSON := StructData{}

	if err := json.Unmarshal(body, &bodyJSON); err != nil {
		return StructData{}, err
	}

	return bodyJSON, nil
}

func CountAge(covidData StructData) (ageCount map[string]int) {
	ageMap := map[string]int{}

	data := covidData.Data

	for _, value := range data {
		if value.Age == nil {
			ageMap["N/A"] += 1
		} else if value.Age.(float64) >= 0 && value.Age.(float64) <= 30 {
			ageMap["0-30"] += 1
		} else if value.Age.(float64) <= 60 {
			ageMap["31-60"] += 1
		} else if value.Age.(float64) >= 61 {
			ageMap["61+"] += 1
		}
	}

	return ageMap
}

func CountProvince(covidData StructData) (provinceCount map[string]int) {
	provinceMap := map[string]int{}
	data := covidData.Data

	for _, value := range data {
		if value.Province == nil {
			provinceMap["N/A"] += 1
		} else {
			provinceMap[value.Province.(string)] += 1
		}
	}

	return provinceMap
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		covidData, _ := FetchData("http://static.wongnai.com/devinterview/covid-cases.json")

		age := CountAge(covidData)

		province := CountProvince(covidData)

		c.JSON(http.StatusOK, gin.H{
			"age":      age,
			"province": province,
		})
	})
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
