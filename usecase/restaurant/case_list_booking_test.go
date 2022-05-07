package restaurant_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
	restaurantMockRepository "github.com/pedrokunz/go_backend/infra/repository/mock/booking"
	restaurantUsecase "github.com/pedrokunz/go_backend/usecase/restaurant"
)

func TestListBooking(t *testing.T) {
	bookingMock := restaurantMockRepository.New()
	ctx := context.Background()
	now := time.Now()
	createValidBookings(t, ctx, bookingMock, now)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "SUCCESS",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := restaurantUsecase.NewListBooking(bookingMock)
			results, err := u.Perform(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecaseListBooking.Perform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, output := range results {
				if output.ID == "" {
					t.Errorf("empty id")
				}
			}
		})
	}

	t.Run("has to be on the date, hour and table order", func(t *testing.T) {
		expectedOutput := []restaurantUsecase.ListBookingOutput{
			{
				Username:     "user_test",
				BookingDate:  now.Format(time.RFC3339),
				CustomerName: "customer_test",
				TableID:      1,
			},
			{
				Username:     "user_test",
				BookingDate:  now.Add(2 * time.Hour).Format(time.RFC3339),
				CustomerName: "customer_test",
				TableID:      1,
			},
			{
				Username:     "user_test",
				BookingDate:  now.Add(4 * time.Hour).Format(time.RFC3339),
				CustomerName: "customer_test",
				TableID:      1,
			},
			{
				Username:     "user_test",
				BookingDate:  now.Format(time.RFC3339),
				CustomerName: "customer_test",
				TableID:      2,
			},
			{
				Username:     "user_test",
				BookingDate:  now.Add(2 * time.Hour).Format(time.RFC3339),
				CustomerName: "customer_test",
				TableID:      2,
			},
		}

		u := restaurantUsecase.NewListBooking(bookingMock)
		output, err := u.Perform(ctx)
		if err != nil {
			t.Fatalf("usecaseListBooking.Perform() error = %v, wantErr %v", err, false)
		}

		for i := range output {
			isEqual := false
			for j := range expectedOutput {
				if i != j {
					continue
				}

				if output[i].BookingDate == expectedOutput[j].BookingDate &&
					output[i].CustomerName == expectedOutput[j].CustomerName &&
					output[i].Username == expectedOutput[j].Username {
					isEqual = true
					break
				}
			}

			if !isEqual {
				t.Fatalf("expected booking is not equal to received booking")
			}
		}
	})

	t.Run("has only booking from today on", func(t *testing.T) {
		err := bookingMock.Create(ctx, &restaurant.Booking{
			Username: "booking",
			CustomerName: "from yesterday",
			Date: time.Now().Add(-24 * time.Hour),
			TableID: 1,
		})
		if err != nil {
			t.Fatalf("failed to create a booking for test")
		}

		u := restaurantUsecase.NewListBooking(bookingMock)
		output, err := u.Perform(ctx)
		if err != nil {
			t.Fatalf("usecaseListBooking.Perform() error = %v, wantErr %v", err, false)
		}

		nowDateString := now.String()
		for _, booking := range output {
			if booking.BookingDate < nowDateString {
				t.Fatalf("invalid booking listed")
			}
		}
	})

	t.Run("has the expected output info", func(t *testing.T) {
		u := restaurantUsecase.NewListBooking(bookingMock)
		output, err := u.Perform(ctx)
		if err != nil {
			t.Fatalf("usecaseListBooking.Perform() error = %v, wantErr %v", err, false)
		}

		firstOutput := output[0]
		if firstOutput.Username != "user_test" ||
			firstOutput.BookingDate != now.Format(time.RFC3339) ||
			firstOutput.CustomerName != "customer_test" ||
			firstOutput.TableID != 1 {
			t.Fatalf("")
		}
	})

	t.Run("ERROR - GetBookingsFromDay", func(t *testing.T) {
		bookingMock.GetBookingsFromDayFunc = func(ctx context.Context, date time.Time) ([]*restaurant.Booking, error) {
			return nil, errors.New(t.Name())
		}

		u := restaurantUsecase.NewListBooking(bookingMock)
		_, err := u.Perform(ctx)
		if err == nil {
			t.Errorf("usecaseListBooking.Perform() error = %v, wantErr %v", err, true)
		}
	})
}
