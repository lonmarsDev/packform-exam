package datastore

func Migrate() {
	PostresDb.AutoMigrate(&Order{}, &OrderItem{}, &Delivery{})
	return
}
