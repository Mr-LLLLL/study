package interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type IExpression interface {
	interpret(stats map[string]float64) bool
}

type AlertRule struct {
	expresstion IExpression
}

func NewAlertRule(rule string) (*AlertRule, error) {
	exp, err := NewAndExpression(rule)
	return &AlertRule{
		expresstion: exp,
	}, err
}

func (r AlertRule) interpret(stats map[string]float64) bool {
	return r.expresstion.interpret(stats)
}

type GreaterExpression struct {
	key   string
	value float64
}

func (g GreaterExpression) interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v > g.value
}

func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if nil != err {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &GreaterExpression{
		key:   data[0],
		value: val,
	}, nil
}

type LessExpression struct {
	key   string
	value float64
}

func (g LessExpression) interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v < g.value
}

func NewLessExpression(exp string) (*LessExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != "<" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if nil != err {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &LessExpression{
		key:   data[0],
		value: val,
	}, nil
}

type AndExpression struct {
	expressions []IExpression
}

func (e AndExpression) interpret(stats map[string]float64) bool {
	for _, exp := range e.expressions {
		if !exp.interpret(stats) {
			return false
		}
	}
	return true
}

func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	expressions := make([]IExpression, len(exps))

	for i, e := range exps {
		var expresstion IExpression
		var err error

		switch {
		case strings.Contains(e, ">"):
			expresstion, err = NewGreaterExpression(e)
		case strings.Contains(e, "<"):
			expresstion, err = NewLessExpression(e)
		default:
			err = fmt.Errorf("exp is invalid: %s", exp)
		}

		if nil != err {
			return nil, err
		}

		expressions[i] = expresstion
	}

	return &AndExpression{
		expressions: expressions,
	}, nil
}
