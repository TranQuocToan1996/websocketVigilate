package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pusher/pusher-http-go"
)

func (repo *DBRepo) PusherAuth(w http.ResponseWriter, r *http.Request) {
	userID := repo.App.Session.GetInt(r.Context(), "userID")

	u, err := repo.DB.GetUserById(userID)
	if err != nil {
		log.Println(err)
		return
	}

	params, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	existData := pusher.MemberData{
		UserID: fmt.Sprint(userID),
		UserInfo: map[string]string{
			"name": u.FirstName,
			"id":   fmt.Sprint(userID),
		},
	}

	res, err := app.WsClient.AuthenticatePresenceChannel(params, existData)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}

}

func (repo *DBRepo) TestPusher(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}
	data["message"] = "helloWorld"
	err := repo.App.WsClient.Trigger("public-channel", "test-event", data)
	if err != nil {
		log.Println(err)
		return
	}
}
