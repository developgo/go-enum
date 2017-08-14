// Code generated by go-enum
// DO NOT EDIT!

package example

import (
	"fmt"
)

const (
	// AnimalCat is a Animal of type Cat
	AnimalCat Animal = iota
	// AnimalDog is a Animal of type Dog
	AnimalDog
	// AnimalFish is a Animal of type Fish
	AnimalFish
)

const _AnimalName = "CatDogFish"

var _AnimalMap = map[Animal]string{
	0: _AnimalName[0:3],
	1: _AnimalName[3:6],
	2: _AnimalName[6:10],
}

func (i Animal) String() string {
	if str, ok := _AnimalMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Animal(%d)", i)
}

var _AnimalValue = map[string]Animal{
	_AnimalName[0:3]:  0,
	_AnimalName[3:6]:  1,
	_AnimalName[6:10]: 2,
}

// ParseAnimal attempts to convert a string to a Animal
func ParseAnimal(name string) (Animal, error) {
	if x, ok := _AnimalValue[name]; ok {
		return Animal(x), nil
	}
	return Animal(0), fmt.Errorf("%s is not a valid Animal", name)
}