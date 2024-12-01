package routes

import (
	"subscription-store/internal/api/handlers"
	repo "subscription-store/internal/repository"
	usecase "subscription-store/internal/usecase"

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

	subscriberRepo := repo.NewInMemorySubscriberRepository()
	subscriberUseCase := usecase.NewSubscriberUseCase(subscriberRepo)
	subscriberHandler := &handlers.SubscriberHandler{
		UseCase: subscriberUseCase,
	}

	r.HandleFunc("/api/v1/subscriptions", subscriptionHandler.CreateSubscription).Methods("POST")
	r.HandleFunc("/api/v1/subscriptions", subscriptionHandler.GetSubscriptions).Methods("GET")
	r.HandleFunc("/api/v1/subscriptions/{subscriptionid}", subscriptionHandler.GetSubscription).Methods("GET")
	r.HandleFunc("/api/v1/subscriptions/{subscriptionid}", subscriptionHandler.UpdateSubscription).Methods("PUT")
	r.HandleFunc("/api/v1/subscriptions/{subscriptionid}", subscriptionHandler.DeleteSubscription).Methods("DELETE")

	// Routes for subscribers
	r.HandleFunc("/api/v1/subscribers/{subscriberid}/{subscriptionid}", subscriberHandler.CreateUserSubscription).Methods("POST")
	r.HandleFunc("/api/v1/subscribers/{subscriberid}", subscriberHandler.GetUserSubscriptions).Methods("GET")
	r.HandleFunc("/api/v1/subscribers/{subscriberid}/{subscriptionid}", subscriberHandler.GetUserSubscription).Methods("GET")
	r.HandleFunc("/api/v1/subscribers/{subscriberid}/{subscriptionid}", subscriberHandler.UpdateUserSubscription).Methods("PUT")
	r.HandleFunc("/api/v1/subscribers/{subscriberid}/{subscriptionid}", subscriberHandler.DeleteUserSubscription).Methods("DELETE")

	return r
}
