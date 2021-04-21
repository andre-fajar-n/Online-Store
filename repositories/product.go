package repositories

import (
	"github.com/andre-fajar-n/Online-Store/config"
	"github.com/andre-fajar-n/Online-Store/models"
)

func GetOneProductByID(productID uint) (*models.Product, error) {
	conn := config.ConnDB
	result := new(models.Product)
	if err := conn.Where("id = ?", productID).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
