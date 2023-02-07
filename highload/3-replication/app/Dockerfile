FROM golang:1.19-alpine as build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o server-app *.go

FROM scratch
COPY --from=build /app/server-app /server-app

ENTRYPOINT [ "/server-app" ]