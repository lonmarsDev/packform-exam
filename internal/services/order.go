package services

import (
	"fmt"

	"github.com/lonmarsDev/packform-exam/pkg/constants"
	"github.com/lonmarsDev/packform-exam/pkg/csv"
	"github.com/lonmarsDev/packform-exam/pkg/datastore"
	"github.com/lonmarsDev/packform-exam/pkg/utils"
)

func ImportOrder() (string, error) {
	FileDir := fmt.Sprintf("%s%s", "test_data/", constants.OrdersFilenamePostgres)
	data, err := csv.NewCsvReaderAll(FileDir)
	if err != nil {
		panic(err)
	}
	orderRecords := make([]*datastore.Order, 0)
	for i, v := range data {
		if i > 0 {
			createdAt, err := utils.StringToDateime(v[1])
			if err != nil {
				panic(err)
			}
			order := &datastore.Order{
				ID:         utils.StringToInt(v[0]),
				CreatedAt:  *createdAt,
				OrderName:  v[2],
				CustomerID: v[3],
			}
			orderRecords = append(orderRecords, order)
		}

	}
	datastore.InsertAllOrder(orderRecords)
	return "", nil
}

func ImportOrderItem() (string, error) {
	FileDir := fmt.Sprintf("%s%s", "test_data/", constants.OrderItemsFileNamePostgres)
	data, err := csv.NewCsvReaderAll(FileDir)
	if err != nil {
		panic(err)
	}
	orderItemRecords := make([]*datastore.OrderItem, 0)
	for i, v := range data {
		if i > 0 {

			orderItem := &datastore.OrderItem{
				ID:           utils.StringToInt(v[0]),
				OrderID:      utils.StringToInt(v[1]),
				PricePerUnit: utils.StringToFloat32(v[2]),
				Quantity:     utils.StringToInt(v[3]),
				Product:      v[4],
			}
			orderItemRecords = append(orderItemRecords, orderItem)
		}

	}
	datastore.InsertAllOrderItem(orderItemRecords)
	return "", nil
}

func AllOrder(page, search, createdDateFrom, createdDateTo string) interface{} {
	orderList := make([]map[string]interface{}, 0)
	offset := utils.StringToInt(page)*5 - 5

	orders := datastore.AllOrder(search, &offset, utils.IntToPointer(5))
	//get Company
	customers := make(map[string]*datastore.Customer)
	customerCompanies := make(map[int]*datastore.CustomerCompany)
	var customersIDs []string

	var customer datastore.Customer
	for _, v := range orders {
		customersIDs = append(customersIDs, v.CustomerID)
	}

	customerList, _ := customer.GetByID(customersIDs)
	for _, v := range customerList {
		if v != nil {
			customers[v.UserID] = v
		}
	}
	var customerCompanyIDs []int
	var customerCompany datastore.CustomerCompany
	for _, v := range customerList {
		customerCompanyIDs = append(customerCompanyIDs, v.CompanyID)
	}
	customerCompanyList, _ := customerCompany.GetByID(customerCompanyIDs)
	for _, v := range customerCompanyList {
		if v != nil {
			customerCompanies[v.CompanyID] = v
		}
	}

	for _, value := range orders {
		newsItem := map[string]interface{}{
			"created_at":       value.CreatedAt.Format("Jan-02-2006 3:04 PM"),
			"order_name":       fmt.Sprintf("%s %s %s", value.OrderName, "|", value.Product),
			"customer_company": "",
			"customer_id":      value.CustomerID,
			"delivered_amount": fmt.Sprintf("%.2f", float64(value.DeliveredQuantity)*value.PricePerUnit),
			"totalAmount":      fmt.Sprintf("%.2f", float64(value.DeliveredQuantity)*value.PricePerUnit),
		}
		if customers[value.CustomerID] != nil {
			newsItem["customer_name"] = customers[value.CustomerID].Name

			if customerCompanies[customers[value.CustomerID].CompanyID] != nil {
				newsItem["customer_company"] = customerCompanies[customers[value.CustomerID].CompanyID].CompanyName
			}

		}

		orderList = append(orderList, newsItem)
	}

	recordCount := len(datastore.AllOrder(search, nil, nil))
	data := map[string]interface{}{
		"data":  orderList,
		"count": recordCount,
	}

	return data
}
