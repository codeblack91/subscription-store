package usecases

import (
	"errors"
	"subscription-store/internal/entities"
	"subscription-store/internal/repository"
)

type SubscriberUseCase interface {
	CreateUserSubscription(subscriberID, subscriptionID string) error
	GetUserSubscriptions(subscriberID string) ([]entities.Subscription, error)
	GetUserSubscription(subscriberID, subscriptionID string) (entities.Subscription, error)
	UpdateUserSubscription(subscriberID, subscriptionID string, sub entities.Subscription) (entities.Subscription, error)
	DeleteUserSubscription(subscriberID, subscriptionID string) error
}

type subscriberUseCase struct {
	repo *repository.InMemorySubscriberRepository
}

func NewSubscriberUseCase(repo *repository.InMemorySubscriberRepository) SubscriberUseCase {
	return &subscriberUseCase{repo: repo}
}

// Создание подписки пользователя
func (uc *subscriberUseCase) CreateUserSubscription(subscriberID, subscriptionID string) error {
	if subscriberID == "" || subscriptionID == "" {
		return errors.New("subscriberID and subscriptionID are required")
	}

	/*
		if sub.Name == "" || sub.Price <= 0 || sub.PaymentsCount == "" {
			return errors.New("invalid subscription data")
		}
	*/
	return uc.repo.CreateUserSubscription(subscriberID, subscriptionID)
}

// Получение всех подписок пользователя
func (uc *subscriberUseCase) GetUserSubscriptions(subscriberID string) ([]entities.Subscription, error) {
	if subscriberID == "" {
		return nil, errors.New("subscriberID is required")
	}

	return uc.repo.GetUserSubscriptions(subscriberID)
}

// Получение подписки пользователя по ID
func (uc *subscriberUseCase) GetUserSubscription(subscriberID, subscriptionID string) (entities.Subscription, error) {
	if subscriberID == "" || subscriptionID == "" {
		return entities.Subscription{}, errors.New("subscriberID and subscriptionID are required")
	}

	return uc.repo.GetUserSubscription(subscriberID, subscriptionID)
}

// Обновление подписки пользователя
func (uc *subscriberUseCase) UpdateUserSubscription(subscriberID, subscriptionID string, sub entities.Subscription) (entities.Subscription, error) {
	if subscriberID == "" || subscriptionID == "" {
		return entities.Subscription{}, errors.New("subscriberID and subscriptionID are required")
	}

	if sub.Name == "" || sub.Price <= 0 || sub.PaymentsCount == "" {
		return entities.Subscription{}, errors.New("invalid subscription data")
	}

	return uc.repo.UpdateUserSubscription(subscriberID, subscriptionID, sub)
}

// Удаление подписки пользователя
func (uc *subscriberUseCase) DeleteUserSubscription(subscriberID, subscriptionID string) error {
	if subscriberID == "" || subscriptionID == "" {
		return errors.New("subscriberID and subscriptionID are required")
	}

	return uc.repo.DeleteUserSubscription(subscriberID, subscriptionID)
}
