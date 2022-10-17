# utrade

## Run

```bash
docker-compose up
```

## Add Invoice

```bash
$ curl -X POST http://localhost:8080/api/v1/workspace/w/invoices -H 'Content-Type: application/json' -d '{"name":"invoice_1"}'
{"name":"invoice_1", "ID":"9745c3dc-a523-4d02-b1d2-9fbf455cff98"}

# use same ID again
curl -X POST http://localhost:8080/api/v1/workspace/w/invoices -H 'Content-Type: application/json' -d '{"name":"invoice_1","ID":"9745c3dc-a523-4d02-b1d2-9fbf455cff98"}'
{"Message":"Conflict","Code":""}
```

## Get Invoice

```bash
$ curl http://localhost:8080/api/v1/workspace/w/invoice/9745c3dc-a523-4d02-b1d2-9fbf455cff98
{"name":"invoice_1", "ID":"9745c3dc-a523-4d02-b1d2-9fbf455cff98"}
```