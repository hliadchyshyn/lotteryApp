# lotteryApp

A simple HTTP service that distributes lottery tickets by email.

## Setup & Run

```bash
go mod download
go run .
```

The server starts on port `8080` by default.

## Configuration

Set the following variables in `.env` before running:

| Variable         | Description                        | Example |
|------------------|------------------------------------|---------|
| `num_of_tickets` | Number of tickets available        | `2`     |
| `PORT`           | Port to listen on (default: 8080)  | `8080`  |

## API

### `POST /ticket`

Issues a lottery ticket to the given email address.

**Headers:**
```
Content-Type: application/json
```

**Request body:**
```json
{
    "email": "user@example.com"
}
```

**Responses:**

| Status | Meaning                          |
|--------|----------------------------------|
| `200`  | Ticket issued successfully       |
| `400`  | Invalid or malformed request     |
| `403`  | Email already has a ticket       |
| `410`  | No tickets left                  |

**Example request:**
```bash
curl -X POST http://localhost:8080/ticket \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com"}'
```

## Testing

```bash
go test ./...
```
