package repository

import (
	"github.com/sidmohanty11/go-webstuffs/BB/pkgs/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
