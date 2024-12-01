package repository

import (
	"errors"
	"subscription-store/internal/entities"
	"sync"
)

type InMemorySubscriberRepository struct {
	mu            sync.RWMutex
	subscriptions map[string]map[string]entities.Subscription
}

func NewInMemorySubscriberRepository() *InMemorySubscriberRepository {
	return &InMemorySubscriberRepository{
		subscriptions: make(map[string]map[string]entities.Subscription),
	}
}

// Создание подписки для пользователя
func (r *InMemorySubscriberRepository) CreateUserSubscription(subscriberID, subscriptionID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.subscriptions[subscriberID]; !exists {
		r.subscriptions[subscriberID] = make(map[string]entities.Subscription)
	}

	if _, exists := r.subscriptions[subscriberID][subscriptionID]; exists {
		return errors.New("subscription already exists for this user")
	}

	//r.subscriptions[subscriberID][subscriptionID] = sub
	return nil
}

// Получение всех подписок пользователя
func (r *InMemorySubscriberRepository) GetUserSubscriptions(subscriberID string) ([]entities.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	userSubscriptions, exists := r.subscriptions[subscriberID]
	if !exists {
		return nil, errors.New("no subscriptions found for this user")
	}

	var subscriptions []entities.Subscription
	for _, sub := range userSubscriptions {
		subscriptions = append(subscriptions, sub)
	}

	return subscriptions, nil
}

// Получение подписки пользователя по ID
func (r *InMemorySubscriberRepository) GetUserSubscription(subscriberID, subscriptionID string) (entities.Subscription, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	userSubscriptions, exists := r.subscriptions[subscriberID]
	if !exists {
		return entities.Subscription{}, errors.New("no subscriptions found for this user")
	}

	sub, exists := userSubscriptions[subscriptionID]
	if !exists {
		return entities.Subscription{}, errors.New("subscription not found")
	}

	return sub, nil
}

// Обновление подписки пользователя
func (r *InMemorySubscriberRepository) UpdateUserSubscription(subscriberID, subscriptionID string, sub entities.Subscription) (entities.Subscription, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	userSubscriptions, exists := r.subscriptions[subscriberID]
	if !exists {
		return entities.Subscription{}, errors.New("no subscriptions found for this user")
	}

	if _, exists := userSubscriptions[subscriptionID]; !exists {
		return entities.Subscription{}, errors.New("subscription not found")
	}

	r.subscriptions[subscriberID][subscriptionID] = sub
	return sub, nil
}

// Удаление подписки пользователя
func (r *InMemorySubscriberRepository) DeleteUserSubscription(subscriberID, subscriptionID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	userSubscriptions, exists := r.subscriptions[subscriberID]
	if !exists {
		return errors.New("no subscriptions found for this user")
	}

	if _, exists := userSubscriptions[subscriptionID]; !exists {
		return errors.New("subscription not found")
	}

	delete(userSubscriptions, subscriptionID)

	// Если у пользователя больше нет подписок, удаляем его запись
	if len(userSubscriptions) == 0 {
		delete(r.subscriptions, subscriberID)
	}

	return nil
}
