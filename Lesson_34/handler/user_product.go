package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func userProductsGet(h Handler, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/user-products/"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	products, err := h.UserProduct.GetUserProducts(id)
	if err != nil {
		http.Error(w, "User_Products not found", http.StatusNotFound)
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

func userProductsPost(h Handler, w http.ResponseWriter, r *http.Request) {
	userIDs, productIDs, quantities := r.URL.Query()["user-id"], r.URL.Query()["product-id"], r.URL.Query()["quantity"]
	userID, err := strconv.Atoi(userIDs[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	productID, err := strconv.Atoi(productIDs[0])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	quantity, err := strconv.Atoi(quantities[0])
	if err != nil {
		http.Error(w, "Invalid stock quantity", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.UserProduct.AddProductToUser(userID, productID, quantity)
	if err != nil {
		http.Error(w, "Failed to add product to user", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "Product with ID %d added to user with ID %d", productID, userID)
}

func userProductsPut(h Handler, w http.ResponseWriter, r *http.Request) {
	userIDs, productIDs, quantities := r.URL.Query()["user-id"], r.URL.Query()["product-id"], r.URL.Query()["quantity"]
	userID, err := strconv.Atoi(userIDs[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	productID, err := strconv.Atoi(productIDs[0])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	quantity, err := strconv.Atoi(quantities[0])
	if err != nil {
		http.Error(w, "Invalid stock quantity", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.UserProduct.UpdateProductQuantityForUser(userID, productID, quantity)
	if err != nil {
		http.Error(w, "Failed to update product quantity", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "Product quantity for user with ID %d updated", userID)
}

func userProductsDelete(h Handler, w http.ResponseWriter, r *http.Request) {
	userIDs, productIDs := r.URL.Query()["user-id"], r.URL.Query()["product-id"]
	userID, err := strconv.Atoi(userIDs[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	productID, err := strconv.Atoi(productIDs[0])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.UserProduct.RemoveProductFromUser(userID, productID)
	if err != nil {
		http.Error(w, "Failed to delete product from user", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "Product with ID %d removed from user with ID %d", productID, userID)
}