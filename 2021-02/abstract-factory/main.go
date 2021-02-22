package main

import (
	"fmt"
	
)

/**
Suppose we want to build a global car factory. If it was factory design pattern, then it was suitable for a single location. But for this pattern, we need multiple locations and some critical design changes.

We need car factories in each location like IndiaCarFactory, USACarFactory and DefaultCarFactory. Now, our application should be smart enough to identify the location where it is being used, so we should be able to use appropriate car factory without even knowing which car factory implementation will be used internally. This also saves us from someone calling wrong factory for a particular location
 */

const (
	CarTypeMini = "MINI"
	CarTypeMicro = "MICRO"
	CarTypeLuxury = "Luxury"
	USA = "USA"
	INDIA = "INDIA"
	EU = "EU"
)

//Car ..
type Car struct {
	Model string
	Location string
}

type ICar interface {
	Paint()
}

type MicroCar struct {
	car Car
}
func (m *MicroCar) Paint() {
	fmt.Printf("painting my micro car\n")
}

type MiniCar struct {
	Car
}
func (m *MiniCar) Paint() {
	fmt.Printf("painting my mini car\n")
}

type LuxuryCar struct {
	Car
}
func (m *LuxuryCar) Paint() {
	fmt.Printf("painting my luxury car")
}

//ICarFactory ...
type ICarFactory interface {
	BuildCar(typ string) ICar
}

//UsaFactory ...
type UsaFactory struct  {
}
func (f *UsaFactory)  BuildCar(typ string) ICar {
	switch typ {
	case CarTypeMicro : return &MicroCar{}
	case CarTypeLuxury : return &LuxuryCar{}
	case CarTypeMini : return &MiniCar{}
	}
	return nil
}


func GetFactoryCar(loc string) ICarFactory {
	switch loc {
	case USA :return  &UsaFactory{}
	//we here should make for other countries 
	}
	return nil
}


func main() {
	carMini  := GetFactoryCar(USA).BuildCar(CarTypeMini)
	carLuxury := GetFactoryCar(USA).BuildCar(CarTypeLuxury)
	carMini.Paint()
	carLuxury.Paint()
}
