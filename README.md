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

upd: 09.12.2024 г.
1. Поменял структуру проекта, появилась папка data access layer
2. usecase переименовал в repository
3. В cmd сделал разделение на server и client
4. Реализовал запросы в DaData, полючаю "RUB" что бы выводить рядом с ценой
5. Тестово вызывают запрос DaData.Currency из main.go
6. Создал папку infrastructure с реализацией клиента для DaData
7. Поправил нейминги файлов
8. Переименовал entities на models
9. Добавил viper для работы с конфигами
10. Выпилил ручки про subscribers (Users)
11. Добавил логгер в utils, но нев недрил в проекте















