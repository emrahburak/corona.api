package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	. "./models"
)

func main() {
	MakeModel()
	return
}

// var CoronaGlobalStore = make(map[string]CoronaGlobal)

// func getCoronaGlobal(body []byte) (*CoronaGlobal, error) {
// 	virus := new(CoronaGlobal)
// 	err := json.Unmarshal(body, &virus)
// 	if err != nil {
// 		fmt.Println("Unmarshal işleminde bir hata oluştu ", err)
// 	}

// 	return virus, err
// }

func MakeModel() {
	db, err := gorm.Open("sqlite3", "db/aboutCorona.sdb")
	db.LogMode(true)
	defer db.Close()

	if err == nil {
		db.AutoMigrate(&CoronaGlobal{})
	}

	coronaModel, err2 := MakeRequest()
	if err != nil {

		log.Fatal(err2)
	}

	// fmt.Println("------------Model---------")
	// fmt.Println("Confirmed : " + strconv.Itoa(coronaModel.Confirmed))
	// fmt.Println("Deaths : " + strconv.Itoa(coronaModel.Deaths))
	// fmt.Println("Recovered : " + strconv.Itoa(coronaModel.Recovered))

	resultModel := CoronaGlobal{Confirmed: coronaModel.Confirmed, Deaths: coronaModel.Deaths, Recovered: coronaModel.Recovered}

	db.Create(&resultModel)

	// WriteToScreen(resultModel)

}

func WriteToScreen(e CoronaGlobal) {
	fmt.Printf("%d\t,%d,%d,%d", e.ID, e.Confirmed, e.Deaths, e.Recovered)

}

func MakeRequest() (*CoronaGlobal, error) {
	resp, err := http.Get("https://wuhan-coronavirus-api.laeyoung.endpoint.ainize.ai/jhu-edu/brief")
	if err != nil {
		log.Fatalln(err)
	}

	decoder := json.NewDecoder(resp.Body)

	val := &CoronaGlobal{}
	err = decoder.Decode(val)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Confirmed : " + strconv.Itoa(val.Confirmed))
	// fmt.Println("Deaths : " + strconv.Itoa(val.Deaths))
	// fmt.Println("Recovered : " + strconv.Itoa(val.Recovered))

	return val, err

}
