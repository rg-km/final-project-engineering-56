package api

import (
	"encoding/json"
	"final-project-engineering-56/Backend/model"
	"net/http"
)

type BukulistResponse struct {
	Error string `json:"error"`
}

type Buku struct {
	Buku model.Books `json:"buku"`
}

func (api *API) Getallbuku(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var bukuList []model.Books
	bukuList, err := api.bukuRepo.GetAllBuku()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(BukulistResponse{Error: "Failed to get all buku"})
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(bukuList)

}

func (api *API) InputBuku(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var Buku model.Books
	err := json.NewDecoder(req.Body).Decode(&Buku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.bukuRepo.INSERTBuku(Buku.ISBN, Buku.Judul, Buku.Penerbit, Buku.Pengarang, Buku.Tahun, Buku.Gambar, Buku.Deskripsi)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Buku berhasil ditambahkan"))
}

func (api *API) GetBukuByID(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var Buku model.Books
	err := json.NewDecoder(req.Body).Decode(&Buku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Buku, err = api.bukuRepo.GetBukuByID(Buku.IDbuku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(Buku)
}

func (api *API) DeleteBukuByID(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var Buku model.Books
	err := json.NewDecoder(req.Body).Decode(&Buku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.bukuRepo.DELETEBuku(Buku.IDbuku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Buku berhasil dihapus"))
}

func (api *API) UpdateBukuByID(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var Buku model.Books
	err := json.NewDecoder(req.Body).Decode(&Buku)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = api.bukuRepo.UPDATEBuku(Buku.IDbuku, Buku.ISBN, Buku.Judul, Buku.Penerbit, Buku.Pengarang, Buku.Tahun, Buku.Gambar, Buku.Deskripsi)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Buku berhasil diupdate"))
}
