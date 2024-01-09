package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"applicationDesignTest/api/v1/orders/controllers"
	"applicationDesignTest/internals/domain/entities"
	"applicationDesignTest/internals/domain/usecases"
	"github.com/gorilla/mux"
)

type OrdersControllerInterface interface {
	GetByUser(ctx context.Context, id int64) ([]entities.Order, error)
	OrderPost(ctx context.Context, order entities.Order) (entities.Order, error)
}

type OrderHandlers struct {
	controllers OrdersControllerInterface
	logger      *log.Logger
}

func NewOrderHandlers(
	ordersUseCases usecases.OrdersUseCases,
	logger *log.Logger,
) *OrderHandlers {
	return &OrderHandlers{
		controllers: controllers.NewOrdersController(ordersUseCases),
		logger:      logger,
	}
}

func (h *OrderHandlers) GetByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var orders []entities.Order
	orders, err = h.controllers.GetByUser(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var b []byte
	b, err = json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (h *OrderHandlers) Post(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		http.Error(w, "nil request", http.StatusBadRequest)
		return
	}

	var order entities.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		msg := fmt.Sprintf("%s %v: %s", r.Method, r.URL, err.Error())
		h.logger.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	order, err = h.controllers.OrderPost(r.Context(), order)
	if err != nil {
		msg := fmt.Sprintf("%s %v: %s", r.Method, r.URL, err.Error())
		h.logger.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	var b []byte
	b, err = json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
