package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/internal/core/interfaces"
)

func Transfer(c *gin.Context) {
	repository := *c.MustGet("repository").(*interfaces.TransferRepository)
	senderID, err := strconv.Atoi(c.PostForm("senderID"))
	addresseeID, err := strconv.Atoi(c.PostForm("addresseeID"))
	money, err := strconv.Atoi(c.PostForm("money"))

	if err != nil && senderID == 0 || addresseeID == 0 || money == 0 || money < 0 {
		c.JSON(http.StatusBadRequest, "invalid arguments")
		return
	}

	sender, err := repository.FindUser(senderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	addressee, err := repository.FindUser(addresseeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = repository.Transfer(sender, addressee, money)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "transaction was successful")
}
