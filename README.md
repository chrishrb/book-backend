# book-backend

Backend lambdas for the bachelor thesis book application.

## Local development

```bash
docker-compose up -d
make all
make local
```

### POST /books

```bash
curl -X POST 'http://localhost:3000/books' \
    -H 'Content-Type: application/json' \
    -d '{"title": "example", "description": "Example description", "year": 2023}'
```

### GET /books

```bash
curl -X GET 'http://localhost:3000/books' | jq
```

### GET /books/{id}

```bash
curl -X GET 'http://localhost:3000/books/{id}' | jq
```

### DELETE /books/{id}

```bash
curl -X DELETE 'http://localhost:3000/books/{id}' | jq
```

### PUT /books/{id}

```bash
curl -X PUT 'http://localhost:3000/books/{id}' \
    -H 'Content-Type: application/json' \
    -d '{"title": "example", "description": "Example description", "year": 2023}'
```

## Deploy to AWS

Use terraform for this step, not the SAM CLI.
