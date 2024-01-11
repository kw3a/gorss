# GORSS
## API
The api is conformed of <ip>:8080/v1/<endpoint>
# Endpoints
## POST /users/
Body expected:

```json
{
  "name": "example_name"
}
```

Correct return:
httpStatusCreated (201)
```json
{
  "id": "7668cd15-52dc-4cde-8e4b-2dc00837927f",
  "created_at": "2024-01-10T23:11:17.16728Z",
  "updated_at": "2024-01-10T23:11:17.16728Z",
  "name": "example_name"
}

```