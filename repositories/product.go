package repositories

import (
	"github.com/andre-fajar-n/Online-Store/config"
	"github.com/andre-fajar-n/Online-Store/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
	conn *gorm.DB
}

func (r *ProductRepo) GetByID(productID uint) (*models.Product, error) {
	r.conn = config.ConnDB
	result := new(models.Product)
	if err := r.conn.Where("id = ?", productID).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
