package restaurant

import (
	"context"

	"github.com/pedrokunz/go_backend/usecase/repository"
)

type usecaseDeleteBooking struct {
	bookingRepository repository.DeleteBooking
}

func NewDeleteBooking(bookingRepository repository.DeleteBooking) *usecaseDeleteBooking {
	return &usecaseDeleteBooking{
		bookingRepository: bookingRepository,
	}
}

func (u *usecaseDeleteBooking) Perform(ctx context.Context, id string) error {
	return u.bookingRepository.Delete(ctx, id)
}
