package api

import (
	"encoding/json"
	"github.com/pedrokunz/go_backend/infra/repository/mock/booking"
	"github.com/pedrokunz/go_backend/usecase/restaurant"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler() {
	h := gin.Default()

	createBooking := restaurant.NewCreateBooking(booking.New())

	h.POST("/booking", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		createBookingInput := restaurant.CreateBookingInput{}
		err = json.Unmarshal(body, &createBookingInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = createBooking.Perform(createBookingInput)
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

	h.Run()
}
