# Имя бинарного файла
APP_NAME = shortLink

# Путь к исходному коду (если нужен)
SRC = ./cmd

# Цель по умолчанию — сборка для Windows
default: build-windows

# Проверка форматирования и статического анализа
check:
	go fmt ./...
	go vet ./...

# Сборка под Windows
build-windows: check
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe $(SRC)

# Сборка под Linux
build-linux: check
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) $(SRC)

# Очистка скомпилированных файлов
clean:
	rm -f $(APP_NAME) $(APP_NAME).exe

# --- Docker команды ---

# Собрать докер-образ
docker-build:
	docker-compose build

# Запустить контейнеры (в фоне)
docker-up:
	docker-compose up -d

# Собрать образы и сразу запустить контейнеры
docker-build-up: docker-build docker-up

# Остановить контейнеры
docker-down:
	docker-compose down

# Пересобрать и запустить контейнеры заново
docker-restart: docker-build docker-down docker-up

# Очистить остановленные контейнеры и неиспользуемые образы (опционально)
docker-clean:
	docker system prune -f