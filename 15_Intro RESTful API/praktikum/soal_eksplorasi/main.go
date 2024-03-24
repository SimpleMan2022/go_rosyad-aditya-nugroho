package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FruitResponse struct {
	Name       string `json:"name"`
	Family     string `json:"family"`
	Nutritions struct {
		Calories      float64 `json:"calories"`
		Fat           float64 `json:"fat"`
		Sugar         float64 `json:"sugar"`
		Carbohydrates float64 `json:"carbohydrates"`
		Protein       float64 `json:"protein"`
	} `json:"nutritions"`
}

func main() {
	fmt.Print("enter fruit name: ")
	var fruitName string
	fmt.Scan(&fruitName)

	url := fmt.Sprintf("https://www.fruityvice.com/api/fruit/%s", fruitName)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting fruits: ", err)
	}
	defer res.Body.Close()

	var fruitResp FruitResponse

	if err := json.NewDecoder(res.Body).Decode(&fruitResp); err != nil {
		log.Fatal("Error decoding fruits: ", err)
	}
	fmt.Println("== FRUIT DATA ==")
	fmt.Println("Fruit Name:", fruitResp.Name)
	fmt.Println("Fruit Family:", fruitResp.Family)
	fmt.Println("Calories:", fruitResp.Nutritions.Calories)
	fmt.Println("Fat:", fruitResp.Nutritions.Fat)
	fmt.Println("Sugar:", fruitResp.Nutritions.Sugar)
	fmt.Println("Carbohydrates:", fruitResp.Nutritions.Carbohydrates)
	fmt.Println("Protein:", fruitResp.Nutritions.Protein)

}
