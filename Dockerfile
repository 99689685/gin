FROM golang:latest

WORKDIR /mydata/ginBlog/gin-web
COPY . .

RUN go build main.go

EXPOSE 10012
ENTRYPOINT ["./main-go-linux"]