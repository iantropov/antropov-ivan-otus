# Homework 1: Social Network Blank

Заготовка социальной сети

Чтобы запустить:

1. Завести .env файл, например, `cp sample.env .env`
2. Поднять сервисы через `docker-compose up`
3. Проверить работу через Postman (импортировав коллекцию `social_network.postman_collection.json`)

### Roadmap (for myself)

+ Go Hello World
+ Handle HTTP requests
+ Marshal/unmarshal JSON
+ Validate HTTP requests
+ Add dockercompose with MySQL
+ Connect to MySQL from code - https://go.dev/doc/tutorial/database-access
+ Add users table
+ Save users into the table
+ Fix utf8 problem
+ Secure storing of passwords
+ Check user password
+ Implement user/register
+ Implement login
+ Implement user/get
+ Add dotenv
+ Add JWT tokens
+ Check for SQL Injections - https://go.dev/doc/database/sql-injection
+ Improve error handling
+ Clean up
+ Add Postman collection
+ Write ReadMe