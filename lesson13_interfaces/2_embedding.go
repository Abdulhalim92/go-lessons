package main

import "fmt"

// Starter Интерфейс для запуска
type Starter interface {
	Start() string
}

// Mover Интерфейс для движения
type Mover interface {
	Move() string
}

// Stopper Интерфейс для остановки
type Stopper interface {
	Stop() string
}

// Vehicle Интерфейс для транспортного средства
type Vehicle interface {
	Starter
	Mover
	Stopper
}

// Реализация автомобиля
type Car struct {
	Brand string
}

func (c Car) Start() string {
	return c.Brand + " started"
}

func (c Car) Move() string {
	return c.Brand + " is moving"
}

func (c Car) Stop() string {
	return c.Brand + " stopped"
}

// Bicycle Реализация велосипеда
type Bicycle struct {
	Brand string
}

func (b Bicycle) Start() string {
	return b.Brand + " started"
}

func (b Bicycle) Move() string {
	return b.Brand + " is moving"
}

func (b Bicycle) Stop() string {
	return b.Brand + " stopped"
}

// OperateVehicle Функция для работы с любым транспортным средством
func OperateVehicle(v Vehicle) {
	fmt.Println(v.Start())
	fmt.Println(v.Move())
	fmt.Println(v.Stop())
}

func embedding() {
	car := Car{Brand: "Toyota"}
	bike := Bicycle{Brand: "Giant"}

	OperateVehicle(car)
	OperateVehicle(bike)
}
