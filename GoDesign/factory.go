package main

import "fmt"

//工厂模式
//接口
type Restaurant interface {
	GetFood()
	SellFood()
}

//结构体对应restaurant
type Chibanting struct {
	name  string
	price string
}

type Wandao struct {
	name  string
	price string
}

type Rongxinguan struct {
	name  string
	price string
}

//方法
func (c *Chibanting) GetFood() {
	c.price = "80"
	c.name = "鳌虾"
	fmt.Println("赤坂亭食物已完成" + c.name + c.price)
}

func (w *Wandao) GetFood() {
	w.price = "90"
	w.name = "鳌虾"
	fmt.Println("万岛食物已完成" + w.name + w.price)
}

func (r *Rongxinguan) GetFood() {
	r.price = "100"
	r.name = "鳌虾"
	fmt.Println("荣新馆食物已完成" + r.name + r.price)
}

func (c *Chibanting) SellFood() {
	c.price = "80"
	c.name = "鳌虾"
	fmt.Println("赤坂亭食物已卖出" + c.name + c.price)
}

func (w *Wandao) SellFood() {
	w.price = "90"
	w.name = "鳌虾"
	fmt.Println("万岛食物已卖出" + w.name + w.price)
}

func (r *Rongxinguan) SellFood() {
	r.price = "100"
	r.name = "鳌虾"
	fmt.Println("荣新馆食物已卖出" + r.name + r.price)
}

//函数新建餐厅

func NewRestaurant(name string) Restaurant {
	if name == "c" {
		return &Chibanting{}
	} else if name == "w" {
		return &Wandao{}
	} else {
		return nil
	}
	return nil
}

func main() {
	NewRestaurant("c").SellFood()
	NewRestaurant("c").GetFood()
}
