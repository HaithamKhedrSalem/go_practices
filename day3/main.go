package main

import "fmt"

type Car struct {
	wheels int
 	doors int
}

type User interface {
	set_name()
	get_name()
}

type Client struct{
	name string
}

type Admin struct{
	name string
}

type Manager struct {
	age int
	salary float64
}

type Employee struct {
	Manager
}

func (managerObject *Manager) set_age(age int) {
	managerObject.age = age
}

func (managerObject *Manager) set_salary(salary float64){
	managerObject.salary = salary
}

func (employeeObject *Employee) get_age() int {
	return employeeObject.age
}

func (employeeObject *Employee) get_salary() float64{
	return employeeObject.salary
}

func (clientObject *Client) set_name(name string){
	clientObject.name = name
}

func (clientObject *Client) get_name() string{
	return clientObject.name
}

func (adminObject *Admin) set_name(name string){
	adminObject.name = name
}

func (adminObject *Admin) get_name() string{
	return adminObject.name
}

func (carObject Car) get_car_wheels() int {
	return carObject.wheels
}

func (carObject Car) get_car_doors() int {
	return carObject.doors
}

func (carObject Car) set_car_wheels(wheels int) Car {
	carObject.wheels = wheels
	return carObject
}

func (carObject Car) set_car_doors(doors int) Car {
	carObject.doors = doors
	return carObject
}

func (carObject *Car) ptr_set_car_wheels(wheels int){
	carObject.wheels = wheels
}

func (carObject *Car) ptr_set_car_doors(doors int){
	carObject.doors = doors
}


func main(){
	var carObject Car
	carObject = carObject.set_car_doors(2)
	carObject = carObject.set_car_wheels(4)
	carObject2 := &Car{0,0}
	carObject2.ptr_set_car_wheels(6)
	carObject2.ptr_set_car_doors(2)
	var clientObject Client
	var adminObject Admin
	(&clientObject).set_name("haitham mohammed khedr")
	(&adminObject).set_name("haitham khedr salem")
	var employeeObject Employee
	employeeObject.set_age(35)
	employeeObject.set_salary(1200.35)
	fmt.Println(carObject.get_car_wheels())
	fmt.Println(carObject.get_car_doors())
	fmt.Println(carObject2.get_car_wheels())
	fmt.Println(carObject2.get_car_doors())
	fmt.Println(clientObject.get_name())
	fmt.Println(adminObject.get_name())
	fmt.Println(employeeObject.get_age())
	fmt.Println(employeeObject.get_salary())
}