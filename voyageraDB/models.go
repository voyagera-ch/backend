// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package voyageraDB

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Activity struct {
	ID        int64
	TripID    pgtype.Int8
	Address   string
	Name      pgtype.Text
	Bio       pgtype.Text
	StartAt   pgtype.Timestamp
	EndAt     pgtype.Timestamp
	Latitude  pgtype.Float8
	Longitude pgtype.Float8
	IsHotel   pgtype.Bool
	IsDeleted pgtype.Bool
}

type Image struct {
	ID           int64
	ActivitiesID pgtype.Int8
	Url          string
	Latitude     pgtype.Float8
	Longitude    pgtype.Float8
	IsActivity   pgtype.Bool
	IsDeleted    pgtype.Bool
}

type Trip struct {
	ID        int64
	StartAt   pgtype.Timestamp
	EndAt     pgtype.Timestamp
	Name      string
	Bio       pgtype.Text
	IsDeleted pgtype.Bool
}

type User struct {
	ID          int64
	Password    string
	Phonenumber pgtype.Text
	Email       pgtype.Text
	IsDeleted   pgtype.Bool
}

type UserTrip struct {
	ID        int64
	UserID    pgtype.Int8
	TripID    pgtype.Int8
	IsAdmin   pgtype.Bool
	IsDeleted pgtype.Bool
}
