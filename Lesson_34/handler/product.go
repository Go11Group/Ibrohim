package handler

import (
	"encoding/json"
	"fmt"
	"http_pg/model"
	"net/http"
	"strconv"
	"strings"
)

func productGet(h Handler, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/product/"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	products, err := h.Product.GetProduct(model.Product{ID: id})
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		fmt.Println("Error: ", err)
		return
	}
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
}

func productPost(h Handler, w http.ResponseWriter, r *http.Request) {
	p := model.Product{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.Product.CreateProduct(p)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	w.Write([]byte("New product inserted to database"))
}

func productPut(h Handler, w http.ResponseWriter, r *http.Request) {
	p := model.Product{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.Product.UpdateProduct(p)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "Product with ID %d updated", p.ID)
}

func productDelete(h Handler, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/product/"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.Product.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "Product with ID %d deleted", id)
}