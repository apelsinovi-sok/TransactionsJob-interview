package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/internal/core/interfaces"
)

func AddMoney(c *gin.Context) {
	repository := *c.MustGet("repository").(*interfaces.TransferRepository)
	userID, err := strconv.Atoi(c.PostForm("id"))
	money, err := strconv.Atoi(c.PostForm("money"))
	if err != nil || userID == 0 || money == 0 || money < 0 {
		c.JSON(http.StatusBadRequest, "invalid arguments")
		return
	}

	user, err := repository.FindUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	repository.AddMoney(user, money)
	c.JSON(http.StatusOK, "money is credited to the account")
}
