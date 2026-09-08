# Ручная проверка Notes API

Запустите `make run`, затем используйте второй терминал.

```bash
curl -i -X POST http://localhost:8080/notes \
  -H 'Content-Type: application/json' \
  -d '{"title":"HTTP","text":"request and response"}'

curl -i http://localhost:8080/notes
curl -i http://localhost:8080/notes/1

curl -i -X PATCH http://localhost:8080/notes/1 \
  -H 'Content-Type: application/json' \
  -d '{"text":"router selects handler"}'

curl -i -X DELETE http://localhost:8080/notes/1
curl -i http://localhost:8080/notes/1
```

Ошибочные запросы:

```bash
curl -i -X POST http://localhost:8080/notes \
  -H 'Content-Type: application/json' \
  -d '{"title":'

curl -i -X POST http://localhost:8080/notes \
  -H 'Content-Type: application/json' \
  -d '{"title":"   ","text":"empty title"}'

curl -i http://localhost:8080/notes/abc
curl -i http://localhost:8080/notes/999
```

До каждого запроса запишите ожидаемый status. После ответа отдельно найдите status, headers и body.
