FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build .

EXPOSE 3000

CMD [ "./social-network" ]