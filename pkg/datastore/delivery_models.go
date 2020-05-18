package datastore

func InsertAllDelivery(deliveries []*Delivery) {
	for _, v := range deliveries {
		PostresDb.Create(v)
	}
	return
}
