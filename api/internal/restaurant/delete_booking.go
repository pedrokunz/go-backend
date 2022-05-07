package restaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrokunz/go_backend/usecase/repository"
	restaurantUseCase "github.com/pedrokunz/go_backend/usecase/restaurant"
)

func HandleDeleteBooking(r *gin.RouterGroup, repo repository.DeleteBooking) {
	deleteBooking := restaurantUseCase.NewDeleteBooking(repo)

	r.DELETE("/booking/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := deleteBooking.Perform(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "booking deleted",
		})
	})
}