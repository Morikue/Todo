# указываем образ, на основе которого будем собирать приложение
FROM golang:1.21

# указываем рабочую папку внутри образа
WORKDIR /build

# установим переменные среды
ENV APP_PORT=3001
ENV LOG_INDEX=todo
ENV LOG_IS_DEBUG=false
ENV LOG_TO_FILE=false

ENV POSTGRES_HOST=postgres
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DBNAME=postgres

ENV USERS_HOST=users-service
ENV USERS_REST_PORT=3000
ENV USERS_GRPC_PORT=50000

ENV RABBITMQ_USER=notifications
ENV RABBITMQ_PASSWORD=notifications
ENV RABBITMQ_HOST=rabbitmq
ENV RABBITMQ_PORT=5672

ENV JAEGER_HOST=jaeger:6831
ENV JAEGER_LOG_SPANS=true

ENV GRPC_HOST=0.0.0.0
ENV GRPC_PORT=50001

# копируем файлы приложения в рабочую папку образа
COPY app /build/app
COPY cmd /build/cmd
COPY config /build/config
COPY internal /build/internal
COPY pkg /build/pkg

COPY go.mod /build/
COPY go.sum /build/

# запускаем в образе команды
# скачиваем зависимости приложения
# собирвем приложение
# делаем бинарный файл исполняемым
RUN go mod download && \
    go build -o cmd/todo/todo ./cmd/todo/main.go && \
    chmod +x cmd/todo/todo

# указываем бинарный файл
ENTRYPOINT ["cmd/todo/todo"]

# открываем порты из образа
EXPOSE 3001/tcp
