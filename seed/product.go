package seed

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/andre-fajar-n/Online-Store/models"
	"gorm.io/gorm"
)

type ProductDummy struct {
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
}

func ProductSeed(db *gorm.DB) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(dir + "/seed/product.json")
	if err != nil {
		log.Println("Error open json file")
		return err
	}
	defer jsonFile.Close()

	jsonString, _ := ioutil.ReadAll(jsonFile)

	var jsonData []ProductDummy

	if err := json.Unmarshal([]byte(string(jsonString)), &jsonData); err != nil {
		log.Println("Error unmarshal")
		return err
	}

	for _, v := range jsonData {
		temp := models.Product{
			Name:     v.Name,
			Quantity: v.Quantity,
			Price:    v.Price,
		}
		if err := db.Create(&temp).Error; err != nil {
			log.Println("Error insert")
			return err
		}
	}

	return nil
}
