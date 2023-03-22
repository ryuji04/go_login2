package server

import (
	"fmt"
	"master/model"
	"master/repository"
	"strconv"
)

type User struct {
	Id        int
	Name      string
	Departure string
}

func NewUser() User {
	return User{}
}

func (m User) Create(name string, departure string) {
	var repo repository.UserRepository = repository.NewUserRepository()
	repo.Create(name, departure)
}

// interfaceが不明点
func (m User) GetAll() interface{} {
	var repo repository.UserRepository = repository.NewUserRepository()
	var users model.Users = repo.GetAll()
	return users
}

func (m User) Confirm(id string, name string) bool {
	var repo repository.UserRepository = repository.NewUserRepository()
	var i, _ = strconv.Atoi(id)

	if i != repo.FindById(id).Id {
		panic("idが違います")
	}

	var users model.Users = repo.FindByName(name)
	for _, user := range users {
		fmt.Println("for後:", user, "name:", name)

		// if _ == nil {
		// 	panic("存在しない名前です")
		// }

		if name != user.Name {
			panic("nameが違います")
		}
	}
	return true

}

func (m User) FindById(id string) model.User {
	var repo repository.UserRepository = repository.NewUserRepository()
	return repo.FindById(id)
}

func (m User) Edit(id string, name string, departure string) {
	var repo repository.UserRepository = repository.NewUserRepository()
	repo.Edit(id, name, departure)
}

func (m User) Delete(id string) {
	var repo repository.UserRepository = repository.NewUserRepository()
	repo.Delete(id)
}
