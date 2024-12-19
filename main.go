package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("user=postgres password=123456 dbname=test port=5432 sslmode=disable"), nil)
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&User{}, &Order{}, &Bill{})
	// uid1, _ := uuid.Parse("14738322-6d7f-4e0a-935e-0e19f17c8757")
	uid2, _ := uuid.Parse("31e8f5fe-9848-4088-b786-ff75f1ab7563")
	// db.Create(&User{ID: uid1, Name: "user-1"})
	// db.Create(&Order{ID: uuid.New(), UserID: uid1, Name: "order-1"})
	// db.Create(&User{ID: uid2, Name: "user-2"})

	user, order := &User{}, &Order{}
	bill := &Bill{}
	userTable, orderTable := user.TableName(), order.TableName()
	billTable := bill.TableName()

	var uo struct {
		User
		*Order
		*Bill
	}
	err = db.Model(user).Select(
		userTable+".*, "+orderTable+".*", billTable+".*",
	).Joins(
		"LEFT JOIN "+orderTable+" ON "+userTable+".id = "+orderTable+".user_id",
	).Joins(
		"LEFT JOIN "+billTable+" ON "+userTable+".id = "+billTable+".user_id",
	).Where(userTable+".id = ?", uid2).Scan(&uo).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("user: ", uo.User, "order: ", uo.Order, "bill: ", uo.Bill)
	// I want get:  {31e8f5fe-9848-4088-b786-ff75f1ab7563 user 1} order:  <nil> bill:  <nil>
	// But I got:  {31e8f5fe-9848-4088-b786-ff75f1ab7563 user 1} order:  &{00000000-0000-0000-0000-000000000000 00000000-0000-0000-0000-000000000000  <nil>} bill:  &{00000000-0000-0000-0000-000000000000 <nil> <nil>}
}
