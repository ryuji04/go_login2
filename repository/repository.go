package repository

import (
	"fmt"
	"master/model"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (m UserRepository) Create(name string, departure string) {
	var user model.User = model.User{Name: name, Departure: departure}
	db.Create(&user)
	db.Save(&user)
}

func (m UserRepository) GetAll() model.Users {
	var users model.Users = model.Users{}
	db.Find(&users)
	return users
}

func (m UserRepository) FindById(id string) model.User {
	var user model.User = model.User{}
	if id == "" {
		panic("idが未記入")
	}
	var i, _ = strconv.Atoi(id)
	db.Find(&user, i)
	fmt.Println("findbyId:", user)
	return user
}

func (m UserRepository) FindByName(name string) model.Users {
	var users model.Users = model.Users{}
	if name == "" {
		panic("nameが未記入")
	}
	db.Where("name = ?", name).Find(&users)
	fmt.Println("findbyName:", users)
	fmt.Println("len(users):", len(users))
	if len(users) == 0 {
		panic("nilです")
	}
	return users
}

func (m UserRepository) Edit(id string, name string, departure string) {
	var user model.User = model.User{}
	fmt.Println("update前", user.Id)
	var i, _ = strconv.Atoi(id)
	db.Model(&user).Where("Id = ?", i).Updates(model.User{Name: name, Departure: departure})
	fmt.Println("update後")
}

func (m UserRepository) Delete(id string) {
	var user model.User = model.User{}
	fmt.Println("delete前")
	var i, _ = strconv.Atoi(id)
	db.Where("id = ?", i).Delete(&user)
	fmt.Println("update後")
}

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "user.db")
	db.DropTableIfExists(&model.User{})
	db.CreateTable(&model.User{})

	if err != nil {
		panic("DB接続失敗")
	}
}
