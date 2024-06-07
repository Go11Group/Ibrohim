package handler

import (
	"http_pg/storage/postgres"
	"net/http"
)

type Handler struct {
	User 		*postgres.UserRepo
	Product 	*postgres.ProductRepo
	UserProduct *postgres.UserProductRepo
}

func NewHandler(h Handler) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/", func(w http.ResponseWriter, r *http.Request) {userGet(h, w, r)})
	mux.HandleFunc("POST /user/", func(w http.ResponseWriter, r *http.Request) {userPost(h, w, r)})
	mux.HandleFunc("PUT /user/", func(w http.ResponseWriter, r *http.Request) {userPut(h, w, r)})
	mux.HandleFunc("DELETE /user/", func(w http.ResponseWriter, r *http.Request) {userDelete(h, w, r)})

	mux.HandleFunc("GET /product/", func(w http.ResponseWriter, r *http.Request) {productGet(h, w, r)})
	mux.HandleFunc("POST /product/", func(w http.ResponseWriter, r *http.Request) {productPost(h, w, r)})
	mux.HandleFunc("PUT /product/", func(w http.ResponseWriter, r *http.Request) {productPut(h, w, r)})
	mux.HandleFunc("DELETE /product/", func(w http.ResponseWriter, r *http.Request) {productDelete(h, w, r)})

	mux.HandleFunc("GET /user-products/", func(w http.ResponseWriter, r *http.Request) {userProductsGet(h,w,r)})
	mux.HandleFunc("POST /user-products/", func(w http.ResponseWriter, r *http.Request) {userProductsPost(h,w,r)})
	mux.HandleFunc("PUT /user-products/", func(w http.ResponseWriter, r *http.Request) {userProductsPut(h,w,r)})
	mux.HandleFunc("DELETE /user-products/", func(w http.ResponseWriter, r *http.Request) {userProductsDelete(h,w,r)})
	return &http.Server{Addr: "localhost:8080", Handler: mux}
}