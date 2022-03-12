package restaurant_test

import (
	bookingMockRepository "github.com/pedrokunz/go_backend/infra/repository/mock/booking"
	tableMockRepository "github.com/pedrokunz/go_backend/infra/repository/mock/table"
	"github.com/pedrokunz/go_backend/usecase/restaurant"
	"testing"
	"time"
)

func TestCreateBooking(t *testing.T) {
	type args struct {
		input restaurant.CreateBookingInput
	}

	saturdayNoon := time.Date(2032, time.March, 21, 12, 0, 0, 0, time.UTC).Format(time.RFC3339)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS",
			args: args{
				input: restaurant.CreateBookingInput{
					Username:     "user_test",
					BookingDate:  saturdayNoon,
					CustomerName: "customer_test",
					TableID:      1,
				},
			},
			wantErr: false,
		},
	}

	tableMock := tableMockRepository.New()
	bookingMock := bookingMockRepository.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := restaurant.NewCreate(tableMock, bookingMock)
			err := u.Perform(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Perform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
