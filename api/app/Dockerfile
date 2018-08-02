FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o main ./app
CMD ["/app/main"]
