# utrade

## Run

```bash
docker-compose up
```

## Add Invoice

```bash
curl -X POST http://localhost:8080/api/v1/workspace/w/invoices -H 'Content-Type: application/json' -d '{"name":"invoice_1"}'
```

## Get Invoice

```bash
curl http://localhost:8080/api/v1/workspace/w/invoice/abcd
```