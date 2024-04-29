package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AIResponse struct {
	Response       string `json:"response"`
	SourceResponse string `json:"source_response"`
	Input          string `json:"input"`
	Prompt         string `json:"prompt"`
}

type Response struct {
	Status string     `json:"status"`
	Data   DetailData `json:"data"`
}

type DetailData struct {
	Text   string
	Laptop Laptop
}

type RequestBody struct {
	Budget        int    `json:"budget"`
	Purpose       string `json:"purpose"`
	Ram           string `json:"ram"`
	Cpu           string `json:"cpu"`
	Display_width string `json:"display_width"`
}

type Laptop struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	RAM          string `json:"ram"`
	CPU          string `json:"cpu"`
	DisplayWidth string `json:"display_width"`
}

func GetLaptopRecommendation(c echo.Context) error {
	var request RequestBody
	c.Bind(&request)

	input := fmt.Sprintf("laptop budget %v,purpose %v,ram %v,cpu %v,display %v",
		request.Budget, request.Purpose, request.Ram, request.Cpu, request.Display_width)
	prompt := "kamu adalah seorang customer service spesifikasi komputer yang memberikan data berupa json dengan field name,description,ram,cpu,display_width lalu semua tipe string dan bahasa inggris dan untuk field name adalah nama merk laptop dan seri laptop tersebut,sebagai catatan jangan diapit dengan tanda bactick dan kata json lagi"
	endpoint := fmt.Sprintf("https://chat.ai.cneko.org/?t=%s&p=%s", strings.ReplaceAll(input, " ", "+"), strings.ReplaceAll(prompt, " ", "+"))

	fmt.Println("endpoint: ", endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatalf("Gagal melakukan permintaan: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Gagal membaca body respons: %v", err)
	}
	var aiResponse AIResponse
	err = json.Unmarshal([]byte(body), &aiResponse)
	if err != nil {
		log.Fatalf("Gagal unmarshal JSON: %v", err)
	}

	var laptop Laptop
	err = json.Unmarshal([]byte(aiResponse.Response), &laptop)
	if err != nil {
		log.Fatalf("Gagal unmarshal JSON: %v", err)
	}

	fmt.Println(laptop)

	fmt.Println(aiResponse.Response)
	return c.JSON(http.StatusOK, Response{
		Status: "Success",
		Data: DetailData{
			Text: fmt.Sprintf("With a budget of Rp %v, you can get %v. Here's a recommendation for a %v laptop within your budget.", request.Budget, laptop.Description, request.Purpose),
			Laptop: Laptop{
				Name:         laptop.Name,
				Description:  laptop.Description,
				RAM:          laptop.RAM,
				CPU:          laptop.CPU,
				DisplayWidth: laptop.DisplayWidth,
			},
		},
	})
}

func main() {
	e := echo.New()
	e.POST("/recommend-laptop", GetLaptopRecommendation)

	e.Logger.Fatal(e.Start(":8080"))
}
