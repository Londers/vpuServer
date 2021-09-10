package data

import (
	"encoding/json"
	"fmt"
	u "github.com/Londers/vpuServer/utils"
	"net/http"
)

//Phone Структура данных об устройстве для работы на СО
type Phone struct {
	Areas    []int64 `json:"areas"`
	Login    string  `json:"login"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Status   struct {
		Cfaze   int64  `json:"cfaze"`
		Connect bool   `json:"connect"`
		DateDB  string `json:"dateDB"`
		Device  string `json:"device"`
		LastOps string `json:"last_ops"`
		Ltime   string `json:"ltime"`
		Nfaze   int64  `json:"nfaze"`
	} `json:"status"`
}

func (phone *Phone) Update() u.Response{

	//db, id := GetDB()
	//defer FreeDB(id)

	var jsonPhone, _ = json.Marshal(&phone)
	//_, err := db.Exec(`UPDATE public.phones SET phone = $1`, jsonPhone)
	//if err != nil {
	//	resp := u.Message(http.StatusInternalServerError, fmt.Sprintf("Phone update error: %s", err.Error()))
	//	return resp
	//}1

	fmt.Printf(string(jsonPhone))

	resp := u.Message(http.StatusOK, "телефон обновлен")
	return resp
}

