package usecases

import (
	"errors"
	"subscription-store/internal/entities"
	"subscription-store/internal/repository"
)

// SubscriptionUseCase — интерфейс для логики подписок
type SubscriptionUseCase interface {
	CreateSubscription(sub entities.Subscription) (entities.Subscription, error)
	GetSubscriptions() ([]entities.Subscription, error)
	GetSubscription(subscriptionID string) (entities.Subscription, error)
	UpdateSubscription(subscriptionID string, sub entities.Subscription) (entities.Subscription, error)
	DeleteSubscription(subscriptionID string) error
}

// subscriptionUseCaseImpl — структура, реализующая интерфейс SubscriptionUseCase
type subscriptionUseCaseImpl struct {
	repo repository.SubscriptionRepository // Интерфейс, а не указатель на интерфейс
}

// NewSubscriptionUseCase — создает новую реализацию интерфейса SubscriptionUseCase
func NewSubscriptionUseCase(repo repository.SubscriptionRepository) SubscriptionUseCase {
	return &subscriptionUseCaseImpl{repo: repo}
}

// Создание подписки
func (uc *subscriptionUseCaseImpl) CreateSubscription(sub entities.Subscription) (entities.Subscription, error) {
	if sub.Name == "" || sub.Price <= 0 || sub.PaymentsCount == "" {
		return entities.Subscription{}, errors.New("invalid subscription data")
	}

	createdSubscription, err := uc.repo.CreateSubscription(sub)
	if err != nil {
		return entities.Subscription{}, err
	}
	return createdSubscription, nil
}

// Получение всех подписок
func (uc *subscriptionUseCaseImpl) GetSubscriptions() ([]entities.Subscription, error) {
	subscriptions, err := uc.repo.GetSubscriptions()
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// Получение подписки по ID
func (uc *subscriptionUseCaseImpl) GetSubscription(subscriptionID string) (entities.Subscription, error) {
	if subscriptionID == "" {
		return entities.Subscription{}, errors.New("subscriptionID is required")
	}

	subscription, err := uc.repo.GetSubscription(subscriptionID)
	if err != nil {
		return entities.Subscription{}, err
	}
	return subscription, nil
}

// Обновление подписки
func (uc *subscriptionUseCaseImpl) UpdateSubscription(subscriptionID string, sub entities.Subscription) (entities.Subscription, error) {
	if subscriptionID == "" || sub.Name == "" || sub.Price <= 0 || sub.PaymentsCount == "" {
		return entities.Subscription{}, errors.New("invalid data for updating subscription")
	}

	updatedSubscription, err := uc.repo.UpdateSubscription(subscriptionID, sub)
	if err != nil {
		return entities.Subscription{}, err
	}
	return updatedSubscription, nil
}

// Удаление подписки
func (uc *subscriptionUseCaseImpl) DeleteSubscription(subscriptionID string) error {
	if subscriptionID == "" {
		return errors.New("subscriptionID is required")
	}

	err := uc.repo.DeleteSubscription(subscriptionID)
	if err != nil {
		return err
	}
	return nil
}
