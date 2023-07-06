package main

type DB interface {
	GetNameById(int) string
	GetGoodsPriceById(int) int
}

type Mysql struct {
}

func (m *Mysql) GetNameById(id int) string {
	return ""
}

func (m *Mysql) GetGoodsPriceById(id int) int {
	return 0
}
