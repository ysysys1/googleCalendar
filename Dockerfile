FROM golang:1.16.13-buster as dev-step

RUN apt-get update && apt-get install -y git
RUN mkdir -p /go/src/github.com/calendar-open

RUN go install github.com/golang/mock/mockgen@v1.6.0

WORKDIR /go/src/github.com/calendar-open
COPY go.mod go.sum ./
RUN go mod download

# For vscode
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
  github.com/ramya-rao-a/go-outline@latest \
  github.com/cweill/gotests/gotests@latest \
  github.com/fatih/gomodifytags@latest \
  github.com/josharian/impl@latest \
  github.com/haya14busa/goplay/cmd/goplay@latest \
  golang.org/x/tools/gopls@latest \
  honnef.co/go/tools/cmd/staticcheck@latest  

EXPOSE 8080

CMD ["go", "run", "main.go"]
