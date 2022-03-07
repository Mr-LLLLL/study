package spec

type ISpecification interface {
	IsSatisfiedBy(o interface{}) (bool, error)
}

type Operator int8

const (
	ORand Operator = iota
	ORor
	ORnot
)

type CompositeSpecification struct {
	cols [][]ISpecification
	flag []Operator
}

func (cs CompositeSpecification) And(s ...ISpecification) CompositeSpecification {
	cols := append([]ISpecification{}, s...)
	cs.cols = append(cs.cols, cols)
	cs.flag = append(cs.flag, ORand)
	return cs
}

func (cs CompositeSpecification) Or(s ...ISpecification) CompositeSpecification {
	cols := append([]ISpecification{}, s...)
	cs.cols = append(cs.cols, cols)
	cs.flag = append(cs.flag, ORor)
	return cs
}

func (cs CompositeSpecification) Not() CompositeSpecification {
	cs.cols = append(cs.cols, []ISpecification{})
	cs.flag = append(cs.flag, ORnot)
	return cs
}

func (cs CompositeSpecification) IsSatisfiedBy(o interface{}) (bool, error) {
	var (
		res = true
		err error
	)
	for i, v := range cs.cols {
		switch cs.flag[i] {
		case ORand:
			if !res {
				break
			}
			res, err = and(v, o)
			if err != nil {
				return res, err
			}
		case ORor:
			if res && i != 0 {
				break
			}
			res, err = or(v, o)
			if err != nil {
				return res, err
			}
		case ORnot:
			res = !res
		}
	}

	return res, nil
}

func NewCompSpec() CompositeSpecification {
	return CompositeSpecification{}
}

func and(s []ISpecification, o interface{}) (bool, error) {
	for _, v := range s {
		satisfied, err := v.IsSatisfiedBy(o)
		if err != nil {
			return false, err
		}
		if !satisfied {
			return false, nil
		}
	}

	return true, nil
}

func or(s []ISpecification, o interface{}) (bool, error) {
	if len(s) == 0 {
		return true, nil
	}

	for _, v := range s {
		satisfied, err := v.IsSatisfiedBy(o)
		if err != nil {
			return false, err
		}

		if satisfied {
			return satisfied, nil
		}
	}

	return false, nil
}
