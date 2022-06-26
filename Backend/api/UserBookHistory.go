package api

import (
	"encoding/json"
	"final-project-engineering-56/Backend/model"
	"io/ioutil"
	"net/http"
	"strings"
)

func (api *API) InsertHistory(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var history model.HistoryBuku
	err := json.NewDecoder(req.Body).Decode(&history)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.historyRepo.HistoryBuku(history.ID_user, history.ID_books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Berhasil Ditambahkan"))

}

func (api *API) Updatehistorybuku(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var Update model.HistoryBuku
	err := json.NewDecoder(req.Body).Decode(&Update)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.historyRepo.UpdateHistoryBuku(Update.ID, Update.ID_user, Update.ID_books, Update.Favorite)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Berhasil Diupdate"))

}

func (api *API) searchBooks(w http.ResponseWriter, req *http.Request) {
	keys := req.URL.Query()["Param"]
	searchTerm := ""

	// if !ok || len(keys[0]) < 1 {
	searchTerm = strings.Replace(keys[0], " ", "+", -1)
	// }

	var client = &http.Client{}
	var baseURL = "https://www.googleapis.com/books/v1/volumes?q=" + searchTerm + "&maxResults=5"

	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := client.Do(request)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	data := model.TempGoogleBooks{}

	_ = json.Unmarshal([]byte(responseData), &data)

	// out, _ := json.Marshal(responseData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseData))
}

//buka isi bukunya
func (api *API) fetchBooks(w http.ResponseWriter, req *http.Request) {
	// var data []model.TempGoogleBooks
	// var apiKey = "AIzaSyBoUiDVozCCFOQUF0oLYGydXMV6Otf9wG0"

	// var baseURL = "https://www.googleapis.com/books/v1/volumes?q=" + searchTerm + "&key=" + apiKey
}

func (api *API) setFavorit(w http.ResponseWriter, req *http.Request ){
	api.AllowOrigin(w, req)
	var UpdateFav model.HistoryBuku
	err := json.NewDecoder(req.Body).Decode(&UpdateFav)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.historyRepo.SetFavorit(UpdateFav.ID, UpdateFav.ID_user, UpdateFav.ID_books, UpdateFav.Favorite)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Berhasil Menjadi Favorit"))
}

func (api *API) GetFav(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var favlist []model.HistoryBuku
	favlist, err := api.historyRepo.GetALLFav()
		if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthError{Error: "Failed to get all Favorit Books"})
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(favlist)
}