# Todo на Golang
Используется микросервисная архитектура: микросервисы Notifications, Users и Todos, а также общий шлюз Gateway.
Межсервисное взаимодействие на gRPC, асинхронное на RabbitMq, Gateway на REST. Также для удобства тестирования есть REST ендпоинты для Users и Todos.
База данных ProstgreSQL. Для работы с БД были использованы библиотеки Sqlx и Pgx.
Также использование Zerolog и Gorilla, Jwt-go для авторизации.
Были созданы Dockerfiles, docker-compose и make файл.
