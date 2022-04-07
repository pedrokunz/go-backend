package restaurant

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	bookingMongo "github.com/pedrokunz/go_backend/infra/repository/mongo/booking"
	restaurantUseCase "github.com/pedrokunz/go_backend/usecase/restaurant"
)

func HandleCreateBooking(r *gin.RouterGroup) {
	mongoRepo, err := bookingMongo.New()
	if err != nil {
		panic(err)
	}

	createBooking := restaurantUseCase.NewCreateBooking(mongoRepo)

	r.POST("/booking", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		createBookingInput := restaurantUseCase.CreateBookingInput{}
		err = json.Unmarshal(body, &createBookingInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = createBooking.Perform(c, createBookingInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "booking created",
		})
	})
}
