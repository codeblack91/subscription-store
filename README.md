# SubscriptionStoreApi


### Компоненты приложения:
- gorilla/mux
- postgreSql (toDo)
- для конфигов (toDo)
- стандартный логер (toDo)
- Swaggo для swagger (toDo)
- Запуск в docker-compose (toDo)
- Шаблон проекта согласно принципов чистой архитектуры

### Доступные api-методы:
[Открыть в Swagger UI](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/codeblack91/subscription-store/main/SubscriptionStoreApi.yaml)


Основные возможности:
1. crud операции для создания подписок;
   - для инициализации подписки в памяти необходимо выполнить:
     curl -X POST http://localhost:8080/api/v1/subscriptions \
  -H "Content-Type: application/json" \
  -d '{
        "id": "50f92801-85e9-43a0-a89c-14733d5b5748",
        "name": "Premium",
        "price": 1000,
        "paymentsCount": "12",
        "isActive": true
      }'
3. crud для добавления подписок пользователям.
















