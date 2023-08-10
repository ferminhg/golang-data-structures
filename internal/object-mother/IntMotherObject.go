package object_mother

import (
	"github.com/go-faker/faker/v4"
	"log"
)

type IntMotherObject struct{}

func NewIntMotherObject() *IntMotherObject {
	log.Println("IntMotherObject created with seed")
	return &IntMotherObject{}
}

func (i IntMotherObject) Get() int {
	value, _ := faker.RandomInt(0, 100)
	return value[0]
}
