package main

import "log"

type IModel interface {
	Validate()
}


type User struct {
	IModel
}

func (u User)Validate() {
	log.Println("call form User")
}

type Mapper struct {
	User IModel
}
func NewMapper() *Mapper {
	m := new(Mapper)
	m.User = User{}
	return m
}

func (m Mapper)GetUser() IModel {
	return m.User
}


func main() {
	mapper := NewMapper()

	User := mapper.GetUser()

	User.Validate()

}

