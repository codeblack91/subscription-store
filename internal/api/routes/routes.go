package routes

import (
	"subscription-store/internal/api/handlers"
	repo "subscription-store/internal/dal/repository"
	usecase "subscription-store/internal/service"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Initialize UseCases
	subscriptionRepo := repo.NewInMemorySubscriptionRepository()
	subscriptionUseCase := usecase.NewSubscriptionUseCase(subscriptionRepo)
	//handler := handlers.NewSubscriptionHandler(subscriptionUseCase)
	subscriptionHandler := &handlers.SubscriptionHandler{
		UseCase: subscriptionUseCase,
	}

	r.HandleFunc("/api/v1/subscriptions", subscriptionHandler.CreateSubscription).Methods("POST")
	r.HandleFunc("/api/v1/subscriptions", subscriptionHandler.GetSubscriptions).Methods("GET")
	r.HandleFunc("/api/v1/subscriptions/{subscriptionid}", subscriptionHandler.GetSubscription).Methods("GET")
	r.HandleFunc("/api/v1/subscriptions/{subscriptionid}", subscriptionHandler.UpdateSubscription).Methods("PUT")
	r.HandleFunc("/api/v1/subscriptions/{subscriptionid}", subscriptionHandler.DeleteSubscription).Methods("DELETE")

	return r
}
