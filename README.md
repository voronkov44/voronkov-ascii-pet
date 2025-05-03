# ASCII-PET API Service

REST API сервис для управления картинками питомцев и их описанием

## 🛠 Технологии
- Go 1.24 (чистая архитектура)
- Docker (для локального развертывания)
- Swagger (документация API)

## Установка и запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/voronkov44/voronkov-ascii-pet.git
```

### 2. Переход в корневую директорию проекта
```bash
cd voronkov-ascii-pet
```

### 3. Развертывание и запуск API через Docker
*Требуется установка [docker](https://www.docker.com/products/docker-desktop/), если не установлен, смотрите [зависимости](https://github.com/voronkov44/voronkov-ascii-pet/tree/main#%D0%B7%D0%B0%D0%B2%D0%B8%D1%81%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B8)*
#### Сборка Docker-образа
```bash
docker build -t ascii-pet .
```
где:
- `-t ascii-pet` - задает тег образу(имя)
- `.` - значит, что контекст сборки в текущей директории

#### Запуск контейнера
```bash
docker run -p 8080:8080 ascii-pet
```
где:
- `-p 8080:8080` - проброс портов с контейнера на хост
- `ascii-pet` - имя образа

## Адреса
- Backend (API) - ```http://localhost:8080/```
- Swagger документация - ```http://localhost:8080/swagger/index.html```


## **Зависимости**
Установка пакета [Docker Engine](https://docs.docker.com/engine/install/)
