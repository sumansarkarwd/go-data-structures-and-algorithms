package main

import "fmt"

type Wind struct {
	direction float64
	speed     float64
}

type CityWeatherData struct {
	cityName    string
	temparature float64
	Wind
}

func main() {
	wind1 := Wind{}
	wind1.set(55.5, 26.7)
	wind1.printWindData()

	cityWeaData1 := CityWeatherData{}
	cityWeaData1.Set("Kolkata", 32.2, 60.0, 40.5)
	cityWeaData1.printCityWeatherData()

}

func (w *Wind) set(direction, s float64) {
	w.direction = direction
	w.speed = s
}

func (w Wind) printWindData() {
	fmt.Println(w)
}

func (c *CityWeatherData) Set(cityName string, t, d, s float64) {
	c.cityName = cityName
	c.temparature = t
	c.direction = d
	c.speed = s
}

func (c CityWeatherData) printCityWeatherData() {
	fmt.Println(c)
}
