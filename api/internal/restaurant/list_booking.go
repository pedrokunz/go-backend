package restaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bookingMongo "github.com/pedrokunz/go_backend/infra/repository/mongo/booking"
	restaurantUseCase "github.com/pedrokunz/go_backend/usecase/restaurant"
)

func HandleListBooking(r *gin.RouterGroup) {
	mongoRepo, err := bookingMongo.New()
	if err != nil {
		panic(err)
	}

	listBooking := restaurantUseCase.NewListBooking(mongoRepo)
	r.GET("/booking", func(c *gin.Context) {
		listBookingInput := restaurantUseCase.ListBookingInput{
			Date: c.Query("date"),
		}

		results, err := listBooking.Perform(c, listBookingInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"bookings": results,
		})
	})
}
