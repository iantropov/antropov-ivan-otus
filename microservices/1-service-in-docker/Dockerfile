FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download && go build .

EXPOSE 8000

CMD [ "./service" ]