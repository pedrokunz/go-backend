package main

import (
	"github.com/pedrokunz/go_backend/api"
	bookingMongo "github.com/pedrokunz/go_backend/infra/repository/mongo/booking"
)

func main() {
	mongoRepo, err := bookingMongo.New()
	if err != nil {
		panic(err)
	}

	api.ListenAndServe(api.ApiDependencies{
		ListBookingRepo:   mongoRepo,
		CreateBookingRepo: mongoRepo,
	})
}
