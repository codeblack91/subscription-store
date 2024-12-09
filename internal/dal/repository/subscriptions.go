package repository

import (
	"errors"
	entities "subscription-store/internal/dal/models"
	"sync"
)

type SubscriptionRepository interface {
	CreateSubscription(sub entities.Subscription) (entities.Subscription, error)
	GetSubscriptions() ([]entities.Subscription, error)
	GetSubscription(subscriptionID string) (entities.Subscription, error)
	UpdateSubscription(subscriptionID string, sub entities.Subscription) (entities.Subscription, error)
	DeleteSubscription(subscriptionID string) error
}

type InMemorySubscriptionRepository struct {
	mu            sync.RWMutex
	subscriptions map[string]entities.Subscription
}

func NewInMemorySubscriptionRepository() *InMemorySubscriptionRepository {
	return &InMemorySubscriptionRepository{
		subscriptions: make(map[string]entities.Subscription),
	}
}

// Создание подписки
func (r *InMemorySubscriptionRepository) CreateSubscription(sub entities.Subscription) (entities.Subscription, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Проверим, существует ли уже подписка с таким ID
	if _, exists := r.subscriptions[sub.ID]; exists {
		return entities.Subscription{}, errors.New("subscription already exists")
	}

	r.subscriptions[sub.ID] = sub
	return sub, nil
}

// Получение всех подписок
func (r *InMemorySubscriptionRepository) GetSubscriptions() ([]entities.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var subs []entities.Subscription
	for _, sub := range r.subscriptions {

		subs = append(subs, sub)
	}

	return subs, nil
}

// Получение подписки по ID
func (r *InMemorySubscriptionRepository) GetSubscription(subscriptionID string) (entities.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	subscription, exists := r.subscriptions[subscriptionID]
	if !exists {
		return entities.Subscription{}, errors.New("subscription not found")
	}
	return subscription, nil
}

// Обновление подписки
func (r *InMemorySubscriptionRepository) UpdateSubscription(subscriptionID string, sub entities.Subscription) (entities.Subscription, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.subscriptions[subscriptionID]
	if !exists {
		return entities.Subscription{}, errors.New("subscription not found")
	}

	r.subscriptions[subscriptionID] = sub
	return sub, nil
}

// Удаление подписки
func (r *InMemorySubscriptionRepository) DeleteSubscription(subscriptionID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.subscriptions[subscriptionID]
	if !exists {
		return errors.New("subscription not found")
	}

	delete(r.subscriptions, subscriptionID)
	return nil
}
