package datastore

import (
	"context"

	"github.com/lonmarsDev/packform-exam/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerCompany struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CompanyID   int                `bson:"companyId,omitempty"`
	CompanyName string             `bson:"companyName,omitempty"`
}

func (u *CustomerCompany) GetByID(CompanyIDs []int) ([]*CustomerCompany, error) {
	ctx := context.TODO()
	filter := bson.M{"companyId": bson.M{"$in": CompanyIDs}}
	cur, err := DbCollections.CustomerCompanies.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query customer company: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	list := make([]*CustomerCompany, 0)
	for cur.Next(ctx) {
		a := &CustomerCompany{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode customer company: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil

}

func InsertAllCustomerCompany(CustomerCompanies []*CustomerCompany) error {
	ctx := context.TODO()
	docs := make([]interface{}, 0)
	for _, customerCompany := range CustomerCompanies {
		doc := bson.D{
			{"companyId", customerCompany.CompanyID},
			{"companyName", customerCompany.CompanyName},
		}
		docs = append(docs, doc)
	}
	_, err := DbCollections.CustomerCompanies.InsertMany(ctx, docs)
	return err
}
