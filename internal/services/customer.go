package services

import (
	"fmt"
	"strings"

	"github.com/lonmarsDev/packform-exam/pkg/constants"
	"github.com/lonmarsDev/packform-exam/pkg/csv"
	"github.com/lonmarsDev/packform-exam/pkg/datastore"
	"github.com/lonmarsDev/packform-exam/pkg/utils"
)

func ImportCustomer() (string, error) {
	FileDir := fmt.Sprintf("%s%s", "test_data/", constants.CustomersFileNameMongoDb)
	data, err := csv.NewCsvReaderAll(FileDir)
	if err != nil {
		panic(err)
	}
	customerRecords := make([]*datastore.Customer, 0)
	for i, v := range data {
		if i > 0 {
			creditCards := strings.Split(v[5], ",")
			customer := &datastore.Customer{
				UserID:      v[0],
				Login:       v[1],
				Password:    v[2],
				Name:        v[3],
				CompanyID:   utils.StringToInt(v[4]),
				CreditCards: creditCards,
			}
			customerRecords = append(customerRecords, customer)
		}

	}
	err = datastore.InsertAllCustomer(customerRecords)
	if err != nil {
		panic(err)
	}
	return "", nil

}

func ImportCustomerCompany() (string, error) {
	FileDir := fmt.Sprintf("%s%s", "test_data/", constants.CustomerCompaniesFileNameMongoDb)
	data, err := csv.NewCsvReaderAll(FileDir)
	if err != nil {
		panic(err)
	}
	customerCompanyRecords := make([]*datastore.CustomerCompany, 0)
	for i, v := range data {
		if i > 0 {
			customerCompany := &datastore.CustomerCompany{
				CompanyID:   utils.StringToInt(v[0]),
				CompanyName: v[1],
			}
			customerCompanyRecords = append(customerCompanyRecords, customerCompany)
		}

	}
	err = datastore.InsertAllCustomerCompany(customerCompanyRecords)
	if err != nil {
		panic(err)
	}
	return "", nil

}
