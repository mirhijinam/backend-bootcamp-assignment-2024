# Тестовое задание для отбора на Avito Backend Bootcamp
## Сервис домов

## Стек
- Golang
- PostgreSQL

### Инструкция по запуску
Для запуска сервиса необходимо:

1. Склонировать репозиторий с проектом: ```git clone https://github.com/mirhijinam/backend-bootcamp-assignment-2024```
2. Войти с директорию с проектом: ```cd backend-bootcamp-assignment-2024```
3. Сбилдить окружение: ```make up```
4. Запустить приложение: ```go run main.go```

Для запуска тестов необходимо:

1. Сбилдить тестовое окружение: ```make up_test```
2. Запустить тесты с покрытием: ```make cover-html```

Для остановки и удаления запущенных контейнеров использовать ```make down``` и ```make down_test``` соответственно.

С дополнительными командами можно ознакомиться в ```Makefile```.

## Функционал сервиса

1. **Фиктивная авторизация. _/dummyLogin_:**
	```
 	curl --location 'localhost:8080/dummyLogin?user_type=moderator'
 	```
	```
	{
	    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2MzA5MzYsImlhdCI6MTcyNDA5NDkzNiwiVXNlciI6eyJJRCI6ImUzOTI4NmU3LWIxOTQtNDEyMy05NGZiLWM2OTI2MDM0NWY0MCIsIlJvbGUiOiJtb2RlcmF0b3IifX0.SV9iBTids72XLluboHhvpvZ5vEF0G1TiY9FkYChVdsw"
	}
 	```
	- Передать тип пользователя (client, moderator)
	- В ответ вернет токен с соответствующим уровнем доступа: обычного пользователя или модератора

2. **Регистрация пользователей по почте и паролю. _/register_:**
	```
 	curl --location 'localhost:8080/register' \
	     --header 'Content-Type: application/json' \
	     --data-raw '{
		"email": "email1@mail.ru",
		"password": "password",
		"user_type": "client"
	     }'
 	```
	```
 	{
 	"user_id": "0e8311bf-d71a-4f9a-bf95-ba5ac53ec88a"
	}
	```
	- В базе создаётся и сохраняется новый пользователь желаемого типа: обычный пользователь (client) или модератор (moderator)
	
3. **Авторизация пользователей по почте и паролю. _/login_:**
   	```
	curl --location 'localhost:8080/login' \
	     --header 'Content-Type: application/json' \
	     --data '{
		"id": "aae32c62-c470-473e-bf3a-7fe6b02b55ca",
		"password": "password"
	     }'
	```
	```
 	{
 	    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwOTYyNTMsImlhdCI6MTcyNDA5NTM1MywiVXNlciI6eyJJRCI6IjBlODMxMWJmLWQ3MWEtNGY5YS1iZjk1LWJhNWFjNTNlYzg4YSIsIlJvbGUiOiJjbGllbnQifX0.ft0DB1gP7-JrBj3RIm6FHuRY6swB-q64aXn5eiJ5Y3E"
	}
	```
	- При успешной авторизации возвращается токен для пользователя с соответствующим уровнем доступа

4. **Создание дома. _/house/create_:**
	```
	curl --location 'localhost:8080/house/create' \
	     --header 'Content-Type: application/json' \
	     --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2MzA5MzYsImlhdCI6MTcyNDA5NDkzNiwiVXNlciI6eyJJRCI6ImUzOTI4NmU3LWIxOTQtNDEyMy05NGZiLWM2OTI2MDM0NWY0MCIsIlJvbGUiOiJtb2RlcmF0b3IifX0.SV9iBTids72XLluboHhvpvZ5vEF0G1TiY9FkYChVdsw' \
	     --data '{
  	     	"address": "Лесная улица, 6, Москва, 125196",
  		"year": 2000,
  		"developer": null
	     }'
	```
	```
	{
	    "id": 1,
	    "address": "Лесная улица, 6, Москва, 125196",
	    "year": 2000,
	    "created_at": "2024-08-19T19:23:24Z"
	}
	```
	- Только модератор имеет возможность создать дом
	- В случае успешного запроса возвращается полная информация о созданном доме

5. **Создание квартиры. _/flat/create_:**
	```
	curl --location 'localhost:8080/flat/create' \
	     --header 'Content-Type: application/json' \
	     --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjM4Nzc3NjksImlhdCI6MTcyMzgzNDU2OSwiVXNlciI6eyJJRCI6IjBlMTJiMGUyLWExZGItNGJmMS05YzY3LTg2ODQyNWJiNjQ0YSIsIlJvbGUiOiJtb2RlcmF0b3IifX0.TMQV4lSQ2Uf-l3gJRZNkMg-HVyYa8wZyNE5lCtZAdnM' \
	     --data '{
 		"number": 3,
 		"house_id": 1,
 		"price": 100,
 		"rooms": 2
	     }'
	```
	```
	{
	    "id": 1,
    	    "house_id": 1,
    	    "price": 100,
 	    "rooms": 2,
    	    "status": "created"
	}
	```
	- Квартиру может создать любой пользователь
	- Объявление получает статус модерации created
   	- У дома, в котором создали новую квартиру, обновляется дата последнего добавления жилья
	- При успешном запросе возвращается полная информация о квартире

6. **Модерация квартиры. _/flat/update_:**
	```
	curl --location 'localhost:8080/flat/update' \
	     --header 'Content-Type: application/json' \
	     --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2MzA5MzYsImlhdCI6MTcyNDA5NDkzNiwiVXNlciI6eyJJRCI6ImUzOTI4NmU3LWIxOTQtNDEyMy05NGZiLWM2OTI2MDM0NWY0MCIsIlJvbGUiOiJtb2RlcmF0b3IifX0.SV9iBTids72XLluboHhvpvZ5vEF0G1TiY9FkYChVdsw' \
	     --data '{
		"id": 1,
		"house_id": 1,
		"status": "on moderation"
	     }'
	```
	```
	{
    	    "id": 1,
    	    "house_id": 1,
    	    "price": 100,
    	    "rooms": 2,
    	    "status": "on moderation"
	}
	```
	- Статус модерации квартиры: created, approved, declined или on moderation
	- Только модератор может изменить статус модерации
	- При успешном запросе возвращается полная информация об обновленной квартире

7. **Получение списка квартир по номеру дома. _/house/{id}_:**
	```
	curl --location 'localhost:8080/house/1' \
	     --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2MzA5MzYsImlhdCI6MTcyNDA5NDkzNiwiVXNlciI6eyJJRCI6ImUzOTI4NmU3LWIxOTQtNDEyMy05NGZiLWM2OTI2MDM0NWY0MCIsIlJvbGUiOiJtb2RlcmF0b3IifX0.SV9iBTids72XLluboHhvpvZ5vEF0G1TiY9FkYChVdsw'
	```
	```
 	{
 	    "flats": [
		{
		    "id": 1,
		    "house_id": 1,
		    "price": 100,
		    "rooms": 2,
		    "status": "on_moderation"
        	}
 	    ]
	}
	```
	- Обычный пользователь увидит квартиры со статусом модерации approved
   	- Модератор увидит квартиры с любым статусом модерации

### Список самостоятельных решений по заданию:

   1. добавил в api.yaml в ручке flat/create необходимое поле – номер добавляемой квартиры, решив проблему, описанную выше в пункте 1;
   2. добавил в api.yaml в ручке register обязательные поля;
   3. добавил в api.yaml в ручке flat/update поля house_id и status вместо rooms и price, так как, исходя из текста в readme, подразумевалось изменения статуса квартиры модератором, а не редактирование ее характеристик;
   4. в описании ручки login происходит расхождение м/у readme, в котором говорится о почте и пароле, и api.yaml, в котором говорится об uuid и пароле. Выбрали вариант api.yaml;
   5. пользовал пакет go-transaction-manager при добавлении квартиры, чтобы консистентно обновлять строку с датой последней добавленной квартиры;
   6. пользовал PK=(house_id, number) для задания уникального идентификатора квартиры;
   7. удалил headers из ответа 500;
