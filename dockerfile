FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
RUN chmod +x ./main
RUN ls -al
CMD ["./main"]
