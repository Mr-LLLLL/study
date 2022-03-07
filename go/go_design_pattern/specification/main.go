package main

import (
	"fmt"
	"log"
	"spec/spec"
)

type Person struct {
	Age       int
	Sex       bool
	President bool
}

type AdultSpecification struct {
	MinAge int
}

func (s AdultSpecification) IsSatisfiedBy(obj interface{}) (bool, error) {
	o := obj.(Person)
	if o.Age >= s.MinAge {
		return true, nil
	}

	return false, nil
}

func NewAdultSpecification() AdultSpecification {
	return AdultSpecification{
		MinAge: 18,
	}
}

type WomenSpecification struct {
	Sex bool
}

func (s WomenSpecification) IsSatisfiedBy(obj interface{}) (bool, error) {
	o := obj.(Person)
	if o.Sex == s.Sex {
		return true, nil
	}

	return false, nil
}

func NewWomenSpecification() WomenSpecification {
	return WomenSpecification{
		Sex: true,
	}
}

type PresidentSpecification struct {
	President bool
}

func NewPresidentSpecification() PresidentSpecification {
	return PresidentSpecification{
		President: true,
	}
}

func (s PresidentSpecification) IsSatisfiedBy(obj interface{}) (bool, error) {
	o := obj.(Person)
	if o.President == s.President {
		return true, nil
	}

	return false, nil
}

func main() {
	fmt.Println(testCompSpecAnd_resultWantTrue())
	fmt.Println(testCompSpecOr_resultWantTrue())
	fmt.Println(testCompSpecNot_resultWantTrue())
	fmt.Println(testCompSpec_resultWantTrue())
}

func testCompSpecAnd_resultWantTrue() bool {
	p := Person{
		Age:       18,
		Sex:       true,
		President: true,
	}

	b, err := spec.NewCompSpec().And(NewAdultSpecification(), NewWomenSpecification(), NewPresidentSpecification()).IsSatisfiedBy(p)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func testCompSpecOr_resultWantTrue() bool {
	p := Person{
		Age:       18,
		Sex:       false,
		President: false,
	}
	b, err := spec.NewCompSpec().Or(NewAdultSpecification(), NewWomenSpecification(), NewPresidentSpecification()).IsSatisfiedBy(p)
	if err != nil {
		log.Fatal(err)
	}

	p1 := Person{
		Age:       17,
		Sex:       false,
		President: true,
	}
	b1, err := spec.NewCompSpec().Or(NewAdultSpecification(), NewWomenSpecification(), NewPresidentSpecification()).IsSatisfiedBy(p1)
	if err != nil {
		log.Fatal(err)
	}

	p2 := Person{
		Age:       17,
		Sex:       true,
		President: false,
	}
	b2, err := spec.NewCompSpec().Or(NewAdultSpecification(), NewWomenSpecification(), NewPresidentSpecification()).IsSatisfiedBy(p2)
	if err != nil {
		log.Fatal((err))
	}

	return b && b1 && b2
}

func testCompSpecNot_resultWantTrue() bool {
	p := Person{
		Age:       17,
		Sex:       false,
		President: true,
	}
	b, err := spec.NewCompSpec().And(NewAdultSpecification(), NewWomenSpecification(), NewPresidentSpecification()).Not().IsSatisfiedBy(p)
	if err != nil {
		log.Fatal(err)
	}

	p1 := Person{
		Age:       17,
		Sex:       false,
		President: true,
	}
	b1, err := spec.NewCompSpec().And(NewAdultSpecification(), NewWomenSpecification(), NewPresidentSpecification()).Not().IsSatisfiedBy(p1)
	if err != nil {
		log.Fatal(err)
	}

	return b && b1
}

func testCompSpec_resultWantTrue() bool {
	p := Person{
		Age:       17,
		Sex:       true,
		President: true,
	}
	b, err := spec.NewCompSpec().And(NewAdultSpecification(), NewWomenSpecification()).Or(NewPresidentSpecification()).IsSatisfiedBy(p)
	if err != nil {
		log.Fatal(err)
	}

	p1 := Person{
		Age:       18,
		Sex:       true,
		President: true,
	}
	b1, err := spec.NewCompSpec().And(NewAdultSpecification(), NewWomenSpecification()).Not().Or(NewPresidentSpecification()).IsSatisfiedBy(p1)
	if err != nil {
		log.Fatal(err)
	}

	p2 := Person{
		Age:       18,
		Sex:       false,
		President: false,
	}
	b2, err := spec.NewCompSpec().Not().And(NewAdultSpecification(), NewWomenSpecification()).Or(NewPresidentSpecification()).Not().IsSatisfiedBy(p2)
	if err != nil {
		log.Fatal(err)
	}

	p3 := Person{
		Age:       17,
		Sex:       true,
		President: true,
	}
	b3, err := spec.NewCompSpec().Or(NewAdultSpecification(), NewWomenSpecification()).Not().And(NewPresidentSpecification()).Or(NewPresidentSpecification()).And(NewAdultSpecification()).Not().IsSatisfiedBy(p3)
	if err != nil {
		log.Fatal(err)
	}

	return b && b1 && b2 && b3
}
