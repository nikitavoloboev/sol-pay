## Run http server

```
go run main.go
```

## Run website

Using [bun](https://bun.sh)

```
cd website
bun run dev
```

## How to create user

with some default SOL balance:
```bash
curl -X POST http://127.0.0.1:8080/users -H "Content-Type: application/json" -d '{"id": 2,"name":"John Updated", "email":"test@test.com"}'
```

## How to get user's details
```bash
curl -X GET http://localhost:8080/users/8
{"bought_products":null,"created_products":[{"id":3,"product_name":"Chair"},{"id":4,"product_name":"Laptop"}],"email":"test@test.com","id":8,"name":"John Updated","wallet":"GrrG9YS4C9EwtWanHQYMejF5UqBXSjQ42qAWjNRy5CyX"}
```

## How to send transaction
```bash
curl -X POST http://localhost:8080/pay \
-H "Content-Type: application/json" \
-d '{"source_user_id": 1, "target_user_id": 2}'
```

## How to create goods
```bash
curl -X POST http://localhost:8080/goods \
-H "Content-Type: application/json" \
-d '{"name": "Laptop", "price": 1500, "user_id": 8}'

```

## How to list/update/delete goods
```bash
curl -X GET http://localhost:8080/goods\?user_id=8
```

## to view db:
```bash
sqlite3 test.db

.tables
.schema user_goods
```
