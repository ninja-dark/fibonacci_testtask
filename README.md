# fibonacci_testtask

Запустить использую Makefile 
1. Сборка Docker образа
```bush
make docker-build
```
2. Скачать образ memcache 
```bush
make docker-cache
```
3. Запустить сервер fibonacci
```bush
make docker-run 
```
пример REST API запроса 

```bush
curl -X GET "localhost:8080/fibonacci?x=1&y=5"
```
