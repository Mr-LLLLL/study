package mock

//go:generate mockgen -source=./db.go -destination=db_mock.go -package=mock
type DB interface {
	Get(key string) (dbres, error)
	Print() string
}

type dbres struct {
	i int
	s string
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value.i
	}

	if len(db.Print()) == 0 {
		return 0
	}

	return -1
}
