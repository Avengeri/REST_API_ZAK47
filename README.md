# Программа для тренировки замены БД.

## 1) Убедитесь, что у вас настроен docker и docker-compose

- Установите Docker на свой компьютер. Инструкцию по установке можно найти [здесь](https://www.docker.com/)
- Проверьте установлен ли `Docker Compose` с помощью команды  `docker compose version`. Если он не установлен, то не
  мои проблемы, решите сами этот вопрос! :thinking:

## 2) Замените .env_example на .env и укажите пароль и имя БД в переменной окружения .env

Например:

```yml
DB_USER=postgres
POSTGRES_PASSWORD=qwerty
```

## 3) Запустите БД через docker-compose
### Выполните команду
```shell
make start 
```

### Если у Вас выдает ошибку, что образ не установлен, то сначала выполните команды:

```shell
docker pull redis:latest  # Скачать образ Redis
```

```shell
docker pull postgres:latest # Скачать образ PostgresQL
```

### Для остановки контейнера используйте команду:

```shell
make stop 
```

### Для перезапуска контейнера с БД используйте команду:

```shell
make restart 
```

## 4) Создайте миграции:

```shell
make migrate_up 
```

### Для отката миграций используйте команду:

```shell
make migrate_down 
```

## 5) Выберите тип БД

- Для работы с БД типа PostgresQL в файле main.go закомментируйте Redis и раскомментируйте Postgres

```Go
    db, err := postgres.NewPostgresDB(postgres.Config{
Host:     viper.GetString("db.host"),
Port:     viper.GetString("db.port"),
Username: viper.GetString("db.username"),
DBName:   viper.GetString("db.dbname"),
SSLMode:  viper.GetString("db.sslmode"),
Password: os.Getenv("POSTGRES_PASSWORD"),
})

if err != nil {
log.Fatalf("Ошибка создания Postgres: %s", err.Error())
}
repo := repository.NewStorageUsersPostgres(db)
```

- Для работы с БД типа Redis в файле main.go закомментируйте Postgres и раскомментируйте Redis

```Go
    rdb, err := redis_storage.NewRedisClient(redis_storage.Config{
Addr: viper.GetString("rdb.address"),
})
repo := repository.NewStorageUsersRedis(rdb)
```

## Еще команды, которые могут пригодиться


```shell
make build # Собирает приложение в exe файл
```

```shell
make run # Собирает приложение в exe файл и запускает его
```

```shell
make swag # Создает swagger документацию
```
## Инструкция как запустить в продакшне:

- Я хз

## Описание Endpoints

### Добавить пользователя

POST http://localhost:8080/user/

```json
{
  "id": 1,
  "name": "ZAK-pipisya"
}
```

### Добавить пользователя

GET http://localhost:8080/user/:id - заменить :id на действующий id

### Проверить пользователя

GET http://localhost:8080/user/check/:id - заменить :id на действующий id

### Удалить пользователя

DELETE http://localhost:8080/user/:id - заменить :id на действующий id

### Получить все id пользователей

GET http://localhost:8080/user/get_all

> Если вы хотите пожаловаться, то обязательно пишите [сюда](https://t.me/zak47) 