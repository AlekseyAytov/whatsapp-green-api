package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	greenapi "github.com/green-api/whatsapp-api-client-golang-v2"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)

		var auth AuthData
		err := json.Unmarshal(body, &auth)
		if err != nil {
			fmt.Println(err)
			return
		}

		green := greenapi.GreenAPI{
			APIURL:           "https://api.green-api.com/",
			MediaURL:         "https://media.green-api.com/",
			IDInstance:       auth.ID,
			APITokenInstance: auth.Token,
		}

		resp, err := green.Account().GetSettings()
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := resp.Body.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write(data)
	}
}

func StateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)

		var auth AuthData
		err := json.Unmarshal(body, &auth)
		if err != nil {
			fmt.Println(err)
			return
		}

		green := greenapi.GreenAPI{
			APIURL:           "https://api.green-api.com/",
			MediaURL:         "https://media.green-api.com/",
			IDInstance:       auth.ID,
			APITokenInstance: auth.Token,
		}

		resp, err := green.Account().GetStateInstance()
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := resp.Body.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write(data)
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)

		var auth AuthData
		err := json.Unmarshal(body, &auth)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(auth)

		green := greenapi.GreenAPI{
			APIURL:           "https://api.green-api.com/",
			MediaURL:         "https://media.green-api.com/",
			IDInstance:       auth.ID,
			APITokenInstance: auth.Token,
		}

		resp, err := green.Sending().SendMessage(auth.Number+"@c.us", auth.Text)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(resp.Body))

		data, err := resp.Body.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write(data)
	}
}

func SendFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)

		var auth AuthData
		err := json.Unmarshal(body, &auth)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(auth)

		green := greenapi.GreenAPI{
			APIURL:           "https://api.green-api.com/",
			MediaURL:         "https://media.green-api.com/",
			IDInstance:       auth.ID,
			APITokenInstance: auth.Token,
		}

		resp, err := green.Sending().SendFileByUrl(auth.Number+"@c.us", auth.File, "test", greenapi.OptionalCaptionSendUrl("Caption"))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(resp.Body))

		data, err := resp.Body.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write(data)
	}
}