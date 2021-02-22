FROM golang:latest

COPY app/ /go/app

WORKDIR /go/app

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/google/uuid

CMD ["go", "run", "app.go"]