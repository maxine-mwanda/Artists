FROM golang

WORKDIR /code
COPY . .
RUN go build main.go
EXPOSE 5000
CMD ["./code/main"]