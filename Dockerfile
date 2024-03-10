FROM golang:1.16.3-alpine3.13

WORKDIR /app

copy . .

# download all dependencies
RUN go get -d -v ./...

# build go application
RUN go build -o main .

#expose the port
EXPOSE 8000


#Run the application
CMD ["./main"]