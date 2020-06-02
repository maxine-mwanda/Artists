package controllers

import (
	"api/middleware"
	"api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	//	controller
	fmt.Println(request.Header.Get("user-agent"))
	log.Println("we are home")
	response := map[string]interface{}{
		"name": "maxine",
		"age":  23,
	}
	middleware.JsonResponse(writer, 200, response)
}

func Singers(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		artists, err := models.Fetchsingers()
		if err != nil {
			fmt.Println("unable to fetch artists")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		middleware.JsonResponse(writer, 200, artists)
	} else {
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println("unable to read artists")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		var artist models.NewArtist
		err = json.Unmarshal(data, &artist)
		if err != nil {
			fmt.Println("unable to read artists to struct")
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		id, err := models.Create_artist(artist)
		if err != nil {
			fmt.Println(err)
			middleware.JsonResponse(writer, http.StatusBadRequest, "error occured")
			return
		}
		middleware.JsonResponse(writer, 200, id)
	}

}

func Singer(writer http.ResponseWriter, request *http.Request) {
	values := mux.Vars(request)
	singer := values["singer"]
	artist := models.Fetchsinger(singer)
	if artist.Id == 0 {
		middleware.JsonResponse(writer, 404, nil)
		return
	}
	middleware.JsonResponse(writer, 200, artist)
}
