package mock

type GoodsDao interface {
	Get(ID int) string
}

func NewDao() *dao {
	return &dao{}
}
