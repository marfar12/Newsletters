# Newsletters

## Run PostgreSQL in Docker

`docker run -d --name newsletters -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=newsletters postgres:15`

## Migrate

### Migrate up

`migrate -path "db/migrations" -database "<DB_CONNECTION_STRING>" up`

### Migrate down

`migrate -path "db/migrations" -database "<DB_CONNECTION_STRING>" down`

## Add enviroment variables

`BASE_URL=http://localhost:3000`\
`SENDGRID_API_KEY=<SENDGRID_API_KEY>`\
`SENDGRID_EMAIL=<SENDGRID_EMAIL>`\
`SECRET=<JWT_SECRET>`\
`DB_CONNECTION_STRING=<DB_CONNECTION_STRING>`

## Run app

`go run cmd/main.go`
