FROM golang:latest

COPY app/ /go/app

WORKDIR /go/app

RUN go get -u github.com/gin-gonic/gin

CMD ["go", "run", "app.go"]