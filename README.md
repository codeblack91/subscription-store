# subscription_shop
[Компоненты приложения](#Компоненты приложения)

#Компоненты приложения:
- gorilla/mux
- postgreSql (toDo)
- для конфигов (toDo)
- стандартный логер
- Swaggo для swagger (toDo) openApi, uapicodegen
- Запуск в docker-compose (toDo)
- Шаблон проекта согласно принципов чистой архитектуры

#Схема вызовов сервисов:
- Вызов каталога: SubscriptionShopBackend GET api/v1/catalog -> CatalogApi GET api/v1/catalog
- Вызов карточки подписки: SubscriptionShopBackend GET api/v1/catalog/subscription/{id} -> CatalogApi GET api/v1/catalog/subscription/{id}

- Создание подписки: SubscriptionShopBackend api/v1/subscriptions -> SubscriptionApi POST api/v1/subscriptions -> NotificationApi POST api/v1/email -> POST api/v1/customer
- Получение подписок подписчика: SubscriptionShopBackend api/v1/subscriptions/{customerid} -> SubscriptionApi POST api/v1/subscriptions/{customerid} -> CusctomerApi GET api/v1/catalog/subscription
- Получение конкретной подписки подписчика: SubscriptionShopBackend api/v1/subscriptions/{customerid}/{subscriptionid} -> SubscriptionApi POST api/v1/subscriptions/{customerid}/{subscriptionid} -> CusctomerApi GET api/v1/catalog/subscription

  ![Приложение подписок](https://github.com/user-attachments/assets/ca18379f-8dec-45ba-aebd-1988f75e7f45)



