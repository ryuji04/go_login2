package model

// var db *gorm.DB

type User struct {
	Id        int    `gorm:"primary_key"`
	Name      string `gorm:"size:100"`
	Departure string `gorm:size:100`
}

type Users []User

// type UserRepository struct {
// }

// func NewUserRepository() UserRepository {
// 	return UserRepository{}
// }

// func (m UserRepository) Create(name string, departure string) {
// 	var user User = User{Name: name, Departure: departure}
// 	db.Create(&user)
// 	db.Save(&user)
// }

// func (m UserRepository) GetAll() Users {
// 	var users Users = Users{}
// 	db.Find(&users)
// 	return users
// }

// func init() {
// 	var err error
// 	db, err = gorm.Open("sqlite3", "user.db")
// 	db.DropTableIfExists(&User{})
// 	db.CreateTable(&User{})

// 	if err != nil {
// 		panic("DB接続失敗")
// 	}
// }
