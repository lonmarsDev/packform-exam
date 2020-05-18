package services

import (
	"fmt"

	"github.com/lonmarsDev/packform-exam/pkg/constants"
	"github.com/lonmarsDev/packform-exam/pkg/csv"
	"github.com/lonmarsDev/packform-exam/pkg/datastore"
	"github.com/lonmarsDev/packform-exam/pkg/utils"
)

func ImportDelivery() (string, error) {
	FileDir := fmt.Sprintf("%s%s", "test_data/", constants.DeliveriesFileNamePostgres)
	data, err := csv.NewCsvReaderAll(FileDir)
	if err != nil {
		panic(err)
	}
	deliveryRecords := make([]*datastore.Delivery, 0)
	for i, v := range data {
		if i > 0 {
			delivery := &datastore.Delivery{
				ID:                utils.StringToInt(v[0]),
				OrderItemID:       utils.StringToInt(v[1]),
				DeliveredQuantity: utils.StringToInt(v[2]),
			}
			deliveryRecords = append(deliveryRecords, delivery)
		}
	}
	datastore.InsertAllDelivery(deliveryRecords)
	return "", nil
}
