FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build .

EXPOSE 3000

FROM migrate/migrate

COPY --from=build /app/users-app /users-app
COPY --from=build /app/db/migrations /migrations

CMD [ "./users-app" ]
