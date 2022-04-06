package api

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"net/http"

	bookingMongo "github.com/pedrokunz/go_backend/infra/repository/mongo/create_booking"
	restaurantUseCase "github.com/pedrokunz/go_backend/usecase/restaurant"

	"github.com/gin-gonic/gin"
)

func NewHandler() {
	h := gin.Default()

	// mockRepo := bookingMock.New()
	mongoRepo, err := bookingMongo.New()
	if err != nil {
		panic(err)
	}

	createBooking := restaurantUseCase.NewCreateBooking(mongoRepo)

	h.POST("/booking", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
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

	h.Run(":" + os.Getenv("HTTP_PORT"))
}
