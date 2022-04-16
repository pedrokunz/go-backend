package restaurant

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrokunz/go_backend/usecase/repository"
	restaurantUseCase "github.com/pedrokunz/go_backend/usecase/restaurant"
)

func HandleCreateBooking(r *gin.RouterGroup, repo repository.CreateBooking) {
	createBooking := restaurantUseCase.NewCreateBooking(repo)

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
