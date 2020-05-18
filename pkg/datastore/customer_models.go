package datastore

import (
	"context"

	"github.com/lonmarsDev/packform-exam/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"userId,omitempty"`
	Login       string             `bson:"login,omitempty"`
	Password    string             `bson:"password,omitempty"`
	Name        string             `bson:"name,omitempty"`
	CompanyID   int                `bson:"companyId,omitempty"`
	CreditCards []string           `bson:"creditCards,omitempty"`
}

func (u *Customer) GetByID(userIDs []string) ([]*Customer, error) {
	ctx := context.TODO()
	filter := bson.M{"userId": bson.M{"$in": userIDs}}
	cur, err := DbCollections.Customers.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query customer: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	list := make([]*Customer, 0)
	for cur.Next(ctx) {
		a := &Customer{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode customer : %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil

}
func InsertAllCustomer(customers []*Customer) error {
	ctx := context.TODO()
	docs := make([]interface{}, 0)
	for _, customer := range customers {
		doc := bson.D{
			{"userId", customer.UserID},
			{"login", customer.Login},
			{"password", customer.Password},
			{"name", customer.Name},
			{"companyId", customer.CompanyID},
			{"creditCards", customer.CreditCards},
		}
		docs = append(docs, doc)
	}
	_, err := DbCollections.Customers.InsertMany(ctx, docs)
	return err
}
