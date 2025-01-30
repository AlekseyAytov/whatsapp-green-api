package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	greenapi "github.com/green-api/whatsapp-api-client-golang-v2"
)

const (
	apiURL   = "https://api.green-api.com/"
	mediaURL = "https://media.green-api.com/"
)

type Adapter struct {
	router   *mux.Router
	pageData PageData
	greenSDK greenapi.GreenAPI
}

func NewAdapter() *Adapter {
	api := Adapter{
		router:   mux.NewRouter(),
		pageData: PageData{},
		greenSDK: greenapi.GreenAPI{
			APIURL:           apiURL,
			MediaURL:         mediaURL,
			IDInstance:       "",
			APITokenInstance: "",
		},
	}
	api.endpoints()
	return &api
}

func (a *Adapter) Router() *mux.Router {
	return a.router
}

func (a *Adapter) endpoints() {
	a.router.Use(logMidlleWare)
	// a.router.Handle(`/`, http.FileServer(http.Dir("web/template/")))
	a.router.HandleFunc("/settings", a.getSettings).Methods(http.MethodPost)
	a.router.HandleFunc(`/state`, a.getState).Methods(http.MethodPost)
	a.router.HandleFunc(`/send`, a.sendMessage).Methods(http.MethodPost)
	a.router.HandleFunc(`/file`, a.sendFile).Methods(http.MethodPost)
	a.router.PathPrefix("/").Handler(http.FileServer(http.Dir("web/template/")))
}

func (a *Adapter) updatePageData(r io.ReadCloser) error {
	newData, _ := io.ReadAll(r)
	err := json.Unmarshal(newData, &a.pageData)
	if err != nil {
		return err
	}
	a.updateGreenAPICred()
	return nil
}

func (a *Adapter) updateGreenAPICred() {
	a.greenSDK.IDInstance = a.pageData.ID
	a.greenSDK.APITokenInstance = a.pageData.Token
}

func (a *Adapter) getSettings(w http.ResponseWriter, r *http.Request) {
	err := a.updatePageData(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := a.greenSDK.Account().GetSettings()
	if err != nil {
		log.Println(err)
		return
	}

	data, err := resp.Body.MarshalJSON()
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(data)
}

func (a *Adapter) getState(w http.ResponseWriter, r *http.Request) {
	err := a.updatePageData(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := a.greenSDK.Account().GetStateInstance()
	if err != nil {
		log.Println(err)
		return
	}

	data, err := resp.Body.MarshalJSON()
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(data)
}

func (a *Adapter) sendMessage(w http.ResponseWriter, r *http.Request) {
	err := a.updatePageData(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := a.greenSDK.Sending().SendMessage(a.pageData.Number+"@c.us", a.pageData.Text)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := resp.Body.MarshalJSON()
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(data)
}

func (a *Adapter) sendFile(w http.ResponseWriter, r *http.Request) {
	err := a.updatePageData(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := a.greenSDK.Sending().SendFileByUrl(a.pageData.Number+"@c.us", a.pageData.File, "test", greenapi.OptionalCaptionSendUrl("Caption"))
	if err != nil {
		log.Println(err)
		return
	}

	data, err := resp.Body.MarshalJSON()
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(data)
}
