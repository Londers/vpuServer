package handlers

import (
	//"fmt"
	//"github.com/ruraomsk/TLServer/internal/model/accToken"
	"net/http"

	"github.com/gin-gonic/gin"

	"../model/data"

	u "github.com/ruraomsk/VPUserver/utils"
)

//ActUpdatePhone обработчик запроса обновления устройства
var ActUpdatePhone = func(c *gin.Context) {
	var phone = &data.Phone{}

	err := c.ShouldBindJSON(&phone)
	if err != nil {
		u.SendRespond(c, u.Message(http.StatusBadRequest, "incorrectly filled data"))
		return
	}

	resp := phone.Update()
	u.SendRespond(c, resp)
}

