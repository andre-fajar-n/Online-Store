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

func CountProductAvailable(productID, customerID uint) (int64, error) {
	conn := config.ConnDB
	var model models.Product
	var total int64

	query := conn.Model(&model)
	query = query.Joins(`
	LEFT JOIN (
		SELECT od.* FROM order_details od
		INNER JOIN orders o ON (o.id = od.order_id AND o.status_id = 'CHECKOUT')
	) lod ON lod.product_id = products.id
	`)
	query = query.Where("products.id = ?", productID, customerID)
	if err := query.Select("SUM(products.quantity - COALESCE(lod.quantity,0))").Scan(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}
