FROM golang:latest 
RUN mkdir /app 
RUN go get github.com/gin-gonic/gin && go get github.com/go-sql-driver/mysql github.com/jordan-wright/email
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
EXPOSE 8000
CMD ["/app/main"]
