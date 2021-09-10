package mainScreen

import (
	"github.com/Londers/vpuServer/model/accToken"
	"github.com/Londers/vpuServer/model/data"
	u "github.com/Londers/vpuServer/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//HMainScreen обработчик открытия сокета
func HMainScreen(c *gin.Context, hub *HubMainScreen) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		u.SendRespond(c, u.Message(http.StatusBadRequest, err.Error()))
		return
	}
	accInfo := new(accToken.Token)
	tokenInfo := new(jwt.Token)

	cookie, err := c.Cookie("Authorization")
	//Проверка куков получили ли их вообще
	if err != nil {
		cookie = ""
		accInfo.IP = c.ClientIP()
		client := &ClientMS{hub: hub, conn: conn, send: make(chan phoneResponse, 256), cInfo: accInfo, rawToken: tokenInfo.Raw, cookie: cookie, isLogin: false, work: true}
		client.listPhone = make(map[string]data.Phone)
		go client.writePump()
		go client.readPump()
		//time.Sleep(1*time.Second)
		client.hub.register <- client
		return
	}
	cookie = cookie[7:]
	for client, _ := range hub.clients {
		if client.cookie == cookie {
			client.work = true
			client.isLogin = true
			client.conn = conn
			client.send = make(chan phoneResponse, 256)
			client.isLogin = true
			client.listPhone = make(map[string]data.Phone)
			hub.clients[client] = true
			go client.writePump()
			go client.readPump()
			hub.mainPage()
			return
		}
	}
	cookie = ""
	accInfo.IP = c.ClientIP()
	client := &ClientMS{hub: hub, conn: conn, send: make(chan phoneResponse, 256), cInfo: accInfo, rawToken: tokenInfo.Raw, cookie: cookie, isLogin: false, work: true}
	client.listPhone = make(map[string]data.Phone)
	go client.writePump()
	go client.readPump()
	resp := newPhoneMess(typeLogOut, nil)
	status := logOut(client.cInfo.Login)
	if status {
		resp.Data["authorizedFlag"] = false
	}
	client.cInfo = new(accToken.Token)
	client.cookie = ""
	client.send <- resp
	client.isLogin = false

	client.hub.register <- client
	return

}
