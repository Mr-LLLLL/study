package main

func GetPrice(db DB, id int) int {
	name := db.GetNameById(id)
	price := db.GetGoodsPriceById(id)

	if name == "mac" {
		price <<= 1
	}

	return price
}
