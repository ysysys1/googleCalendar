FROM golang:1.16.13-buster as dev-step

RUN apt-get update && apt-get install -y git
RUN mkdir -p /go/src/github.com/calendar-open
WORKDIR /go/src/github.com/calendar-open
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "main.go"]
