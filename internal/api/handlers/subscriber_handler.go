package handlers

import (
	"encoding/json"
	"net/http"
	entities "subscription-store/internal/entities"
	usecases "subscription-store/internal/usecase"

	"github.com/gorilla/mux"
)

type SubscriberHandler struct {
	UseCase usecases.SubscriberUseCase
}

func (h *SubscriberHandler) CreateUserSubscription(w http.ResponseWriter, r *http.Request) {
	subscriberID := mux.Vars(r)["subscriberid"]
	subscriptionID := mux.Vars(r)["subscriptionid"]
	err := h.UseCase.CreateUserSubscription(subscriberID, subscriptionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *SubscriberHandler) GetUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriberID := mux.Vars(r)["subscriberid"]
	subscriptions, err := h.UseCase.GetUserSubscriptions(subscriberID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subscriptions)
}

func (h *SubscriberHandler) GetUserSubscription(w http.ResponseWriter, r *http.Request) {
	subscriberID := mux.Vars(r)["subscriberid"]
	subscriptionID := mux.Vars(r)["subscriptionid"]
	subscription, err := h.UseCase.GetUserSubscription(subscriberID, subscriptionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subscription)
}

func (h *SubscriberHandler) UpdateUserSubscription(w http.ResponseWriter, r *http.Request) {
	subscriberID := mux.Vars(r)["subscriberid"]
	subscriptionID := mux.Vars(r)["subscriptionid"]
	var sub entities.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updated, err := h.UseCase.UpdateUserSubscription(subscriberID, subscriptionID, sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (h *SubscriberHandler) DeleteUserSubscription(w http.ResponseWriter, r *http.Request) {
	subscriberID := mux.Vars(r)["subscriberid"]
	subscriptionID := mux.Vars(r)["subscriptionid"]
	if err := h.UseCase.DeleteUserSubscription(subscriberID, subscriptionID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
