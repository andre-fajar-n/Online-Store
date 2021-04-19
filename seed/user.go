package seed

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/andre-fajar-n/Online-Store/models"
	"gorm.io/gorm"
)

type UserDummy struct {
	Name     string `json:"name"`
	Username string `json:"Username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func UserSeed(db *gorm.DB) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(dir + "/seed/user.json")
	if err != nil {
		log.Println("Error open json file")
		return err
	}
	defer jsonFile.Close()

	jsonString, _ := ioutil.ReadAll(jsonFile)

	var jsonData []UserDummy

	if err := json.Unmarshal([]byte(string(jsonString)), &jsonData); err != nil {
		log.Println("Error unmarshal")
		return err
	}

	for _, v := range jsonData {
		temp := models.User{
			Name:     v.Name,
			Username: v.Username,
			Password: v.Password,
			Role:     v.Role,
		}

		if err := db.Create(&temp).Error; err != nil {
			log.Println("Error insert")
			return err
		}
	}

	return nil
}
