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
-d '{"source_user_id": 1, "target_user_id": 2, "product_id": 300}'
```

it will return transaction id:

```bash
{"tix":"3kvAHBE25xQ4eYE6BMsnHuYwwszds5Z1f9KcVnwCVWnUfzdm8EDTQetcpdjzhzrUWfd54aF5MFfiMEy65LtWoatt"}
```

## How to view balance for particular user_id in SOL and USD

```bash
curl -X POST http://127.0.0.1:8080/balance -H "Content-Type: application/json" -d '{"user_id": 2}'
{"balance":"1.0000000000","balance_usd":"19.59"}
```

## How to create goods

```bash
curl -X POST http://localhost:8080/goods \
-H "Content-Type: application/json" \
-d '{"name": "Book", "price": 1, "created_by": 11}'
```

## How to get product's details

```bash
 curl -X GET http://localhost:8080/goods/2
 {"id":2,"product_name":"Chair","count":1}
```

## How to check whethere user have balance sufficient balance to buy product

```bash
curl -X POST http://localhost:8080/can-i-buy \
-H "Content-Type: application/json" \
-d '{"user_id": 7, "product_id": 5}'
{"result":true}
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
