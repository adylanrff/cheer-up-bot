FROM golang:1.13.7-alpine3.11

WORKDIR src/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build cmd/main/cheer_up_bot.go
CMD ["./cheer_up_bot"]
