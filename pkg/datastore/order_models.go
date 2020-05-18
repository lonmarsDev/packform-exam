package datastore

func InsertAllOrder(orders []*Order) {
	for _, v := range orders {
		PostresDb.Create(v)
	}
	return
}

func InsertAllOrderItem(orderitems []*OrderItem) {
	for _, v := range orderitems {
		PostresDb.Create(v)
	}
	return
}

func AllOrder(searchKey string, offset, limit *int) []*Result {
	var results []*Result

	tx := PostresDb
	tx = tx.Table("orders").Select("orders.order_name, orders.created_at, orders.customer_id, order_items.price_per_unit, order_items.quantity, order_items.product, deliveries.delivered_quantity").Joins("left join order_items on order_items.order_id = orders.id").Joins("left join deliveries on deliveries.order_item_id = order_items.id")
	if searchKey != "" {
		search := "%" + searchKey + "%"
		tx = tx.Where(" order_items.product LIKE ? OR orders.order_name LIKE ?", search, search)
	}
	if offset != nil && limit != nil {
		tx = tx.Offset(*offset).Limit(*limit)
	} else {

	}
	tx.Scan(&results)
	return results
}
