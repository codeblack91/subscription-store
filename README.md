# subscription-store


# Компоненты приложения:
- gorilla/mux
- postgreSql (toDo)
- для конфигов (toDo)
- стандартный логер
- Swaggo для swagger (toDo) openApi, oapi-codegen
- Запуск в docker-compose (toDo)
- Шаблон проекта согласно принципов чистой архитектуры

# Доступные api-методы:
- POST api/v1/subscriptions - создает подписку
- GET api/v1/subscriptions - возвращает доступные подписки
- GET api/v1/subscriptions/{subscriptionid} - возвращает доступную подписку
- PUT api/v1/subscriptions/{subscriptionid} - изменяет подписку
- DELETE api/v1/subscriptions/{subscriptionid} - удаляет подписку


- POST api/v1/subscriptions/{subscriberid} - создает подписку пользователю
- GET api/v1/subscriptions/{subscriberid} - возвращает принадлежащие пользователю подписки
- GET api/v1/subscriptions/{subscriberid}/{subscriptionid} - возвращает конкретную подписку пользователя
- PUT api/v1/subscriptions/{subscriberid}/{subscriptionid} - изменяет принадлежащую пользователю подписку
- DELETE api/v1/subscriptions/{subscriberid}/{subscriptionid} - удаляет принадлежащую пользователю подписку



# subscription_shop


# Компоненты приложения:
- gorilla/mux
- postgreSql (toDo)
- для конфигов (toDo)
- стандартный логер
- Swaggo для swagger (toDo) openApi, oapi-codegen
- Запуск в docker-compose (toDo)
- Шаблон проекта согласно принципов чистой архитектуры

# Доступные api-методы:
- POST api/v1/subscriptions - создает подписку
- GET api/v1/subscriptions - возвращает доступные подписки
- GET api/v1/subscriptions/{subscriptionid} - возвращает доступную подписку
- PUT api/v1/subscriptions/{subscriptionid} - изменяет подписку
- DELETE api/v1/subscriptions/{subscriptionid} - удаляет подписку


- POST api/v1/subscriptions/{subscriberid} - создает подписку пользователю
- GET api/v1/subscriptions/{subscriberid} - возвращает принадлежащие пользователю подписки
- GET api/v1/subscriptions/{subscriberid}/{subscriptionid} - возвращает конкретную подписку пользователя
- PUT api/v1/subscriptions/{subscriberid}/{subscriptionid} - изменяет принадлежащую пользователю подписку
- DELETE api/v1/subscriptions/{subscriberid}/{subscriptionid} - удаляет принадлежащую пользователю подписку

# API Documentation

Для просмотра спецификации воспользуйтесь Swagger UI:

<iframe src="https://petstore.swagger.io/?url=https://example.com/openapi.yaml" width="100%" height="500px"></iframe>













