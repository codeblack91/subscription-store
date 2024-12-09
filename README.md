# SubscriptionStoreApi
Функции сервиса:
- Cоздание подписок; (done)
- Подключение подписок пользователю (toDo)
- Заполнение личных данных пользователя; (toDo)
- Уведомления; (toDo)
- Тарификация подписок; (toDo)

### Компоненты приложения:
- gorilla/mux (done)
- postgreSql (toDo)
- для конфигов (done)
- стандартный логер (toDo)
- Swaggo для swagger (toDo)
- Запуск в docker-compose (toDo)
- Шаблон проекта согласно принципов чистой архитектуры (done)

### Доступные api-методы:
[Открыть в Swagger UI](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/codeblack91/subscription-store/main/SubscriptionStoreApi.yaml)

## История изменений:
upd: 01.12.2024 г.
1. Создание проекта;
2. Шаблон чистой архитектуры;
3. Хранение данных приложения в map;
4. crud операции для создания подписок;
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
5. crud для добавления подписок пользователям;
6. Подготовка openapi файла.
















