package room

import (
	"be/configs"
	"be/entities"
	"be/repository/database/image"
	"be/repository/database/user"
	"be/utils"
	"fmt"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run create", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}

		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}

		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 100, Name: "room1 name", Price: 100, Description: "room1 detail"}
		res, err := repo.Create(mock2)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})

	t.Run("success run Update", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail"}
		res2, err2 := repo.Create(mock2)
		if err2 != nil {
			t.Fatal()
		}
		mock3 := entities.Room{Name: "room3 name", Price: 300, Description: "room3 detail"}
		res, err := repo.Update(res1.User_uid, res2.Room_uid, mock3)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// log.Info(res)
	})
}
func TestGetAll(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})

	t.Run("success run get all", func(t *testing.T) {

		//mock User
		mockUser1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		resu1, err1 := user.New(db).Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockUser2 := entities.User{Name: "user2 name", Email: "user2 email", Password: "user1 password"}
		resu2, err2 := user.New(db).Create(mockUser2)
		if err2 != nil {
			t.Fatal()
		}
		mockUser3 := entities.User{Name: "user3 name", Email: "user3 email", Password: "user1 password"}
		_, err3 := user.New(db).Create(mockUser3)
		if err3 != nil {
			t.Fatal()
		}
		//==================

		city := "1"
		var category string = ""
		var name string = ""
		var length string = "1"
		var s string = "room"

		var status string = ""
		mockroom1 := entities.Room{User_uid: resu1.User_uid, City_id: 1, Address: "JL.Dramaga", Name: "room1 name", Price: 100, Description: "room1 detail", Status: "open"}
		_, errroom1 := repo.Create(mockroom1)
		if errroom1 != nil {
			t.Fatal()
		}
		mockroom2 := entities.Room{User_uid: resu2.User_uid, City_id: 1, Address: "JL.Dramaga", Name: "roxoom2 name", Price: 100, Description: "room2 detail", Status: "open"}
		_, errroom2 := repo.Create(mockroom2)
		if errroom2 != nil {
			t.Fatal()
		}

		mockroom3 := entities.Room{User_uid: resu2.User_uid, City_id: 2, Address: "JL.Dramaga", Name: "roxoxom3 name", Price: 100, Description: "room1 detail", Status: "open"}
		_, errroom3 := repo.Create(mockroom3)
		if errroom3 != nil {
			t.Fatal()
		}
		mockroom4 := entities.Room{User_uid: resu2.User_uid, City_id: 2, Address: "JL.Dramaga", Name: "room3 name", Price: 100, Description: "room1 detail", Status: "open"}
		_, errroom4 := repo.Create(mockroom4)
		if errroom4 != nil {
			t.Fatal()
		}
		mockroom5 := entities.Room{User_uid: resu2.User_uid, City_id: 3, Name: "room3 name", Price: 100, Description: "room1 detail", Status: "open"}
		_, errroom5 := repo.Create(mockroom5)
		if errroom5 != nil {
			t.Fatal()
		}

		res, _ := repo.GetAll(s, city, category, name, length, status)
		log.Info(res)

		// assert.Equal(t, "0", res[0].ID)
		assert.Equal(t, "roxoom2 name", res[0].Name)
		// assert.Equal(t, res)
		log.Info(res)
	})
}

func TestGetByID(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run GetById", func(t *testing.T) {
		mock1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		res1, err1 := user.New(db).Create(mock1)
		if err1 != nil {
			t.Fatal()
		}
		mock2 := entities.Room{User_uid: res1.User_uid, City_id: 1, Name: "room1 name", Price: 100, Description: "room1 detail", Category: "superior"}
		res2, err2 := repo.Create(mock2)
		if err2 != nil {
			t.Fatal()
		}

		mock3 := image.ImageReq{}

		for i := 0; i < 3; i++ {
			mock3.Array = append(mock3.Array, image.ImageInput{Url: (fmt.Sprintf("url%d", i+1))})
		}

		if err := image.New(db).Create(res2.Room_uid, mock3); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(res2.Room_uid)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res)
	})
}

func TestGetAllRoom(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Image{})
	db.Migrator().DropTable(&entities.Booking{})
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Image{})

	t.Run("success run GetAllRoom all", func(t *testing.T) {
		mockUser1 := entities.User{Name: "user1 name", Email: "user1 email", Password: "user1 password"}
		resu1, err1 := user.New(db).Create(mockUser1)
		if err1 != nil {
			t.Fatal()
		}
		mockroom1 := entities.Room{User_uid: resu1.User_uid, City_id: 257, Address: "JL.Dramaga", Name: "biasa name", Price: 100, Description: "room1 detail", Status: "open", Category: "standart"}
		_, errroom1 := repo.Create(mockroom1)
		if errroom1 != nil {
			t.Fatal()
		}
		mockroom2 := entities.Room{User_uid: resu1.User_uid, City_id: 232, Address: "JL.Dramaga", Name: "mewah name", Price: 100, Description: "room2 detail", Status: "open", Category: "superior"}
		_, errroom2 := repo.Create(mockroom2)
		if errroom2 != nil {
			t.Fatal()
		}
		mockroom3 := entities.Room{User_uid: resu1.User_uid, City_id: 212, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom3 := repo.Create(mockroom3)
		if errroom3 != nil {
			t.Fatal()
		}
		mockroom4 := entities.Room{User_uid: resu1.User_uid, City_id: 200, Address: "JL.Dramaga", Name: "sederhana name", Price: 100, Description: "room1 detail", Status: "open", Category: "luxury"}
		_, errroom4 := repo.Create(mockroom4)
		if errroom4 != nil {
			t.Fatal()
		}

		res, err := repo.GetAllRoom(0, "malang", "", "", "")
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Info(res, len(res))

	})
}
