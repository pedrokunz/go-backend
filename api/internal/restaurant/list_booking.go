package restaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrokunz/go_backend/usecase/repository"
	restaurantUseCase "github.com/pedrokunz/go_backend/usecase/restaurant"
)

func HandleListBooking(r *gin.RouterGroup, repo repository.ListBooking) {
	listBooking := restaurantUseCase.NewListBooking(repo)
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
