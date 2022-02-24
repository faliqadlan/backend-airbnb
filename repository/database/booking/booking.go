package booking

import (
	"be/entities"
	"errors"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type BookingDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookingDb {
	return &BookingDb{
		db: db,
	}
}

func (repo *BookingDb) Create(user_uid string, room_uid string, newBooking entities.Booking) (BookingCreateResp, error) {

	var uid string

	for {
		uid = shortuuid.New()
		find := entities.Room{}
		res := repo.db.Model(&entities.Room{}).Where("booking_uid =  ?", uid).First(&find)
		if res.RowsAffected == 0 {
			break
		}
		if res.Error != nil {
			return BookingCreateResp{}, res.Error
		}
	}

	newBooking.Booking_uid = uid

	// check reservation
	bookingInit := entities.Booking{}

	resRev := repo.db.Model(entities.Booking{}).Where("status = reservation OR status = onGoing AND end_date >= ? AND start_date <= ?", newBooking.Start_date, newBooking.End_date).Find(bookingInit)

	if resRev.RowsAffected != 0 {
		return BookingCreateResp{}, errors.New("the date already picked up")
	}

	res := repo.db.Model(&entities.Booking{}).Create(&newBooking)

	if res.Error != nil {
		return BookingCreateResp{}, res.Error
	}
	var bookingres BookingCreateResp

	resp := repo.db.Model(&entities.Booking{}).Where("booking_uid =?", newBooking.Booking_uid).Select("bookings.booking_uid as Booking_uid, rooms.name as Name, bookings.description as Description,bookings.start_date as Start_date,bookings.end_date as End_date, sum(rooms.price*bookings.day) as Price_total").Joins("inner join rooms on bookings.room_uid = rooms.room_uid").First(&bookingres)
	if resp.Error != nil {
		return BookingCreateResp{}, res.Error
	}
	return bookingres, nil
}
