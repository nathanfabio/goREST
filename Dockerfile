FROM golang:1.22

#set working directory
WORKDIR /go/src/app

#copy the source code
COPY . .

#expose the port
EXPOSE 8080

#build the go app
RUN go build -o main cmd/main.go

#run the executable
CMD ["./main"]