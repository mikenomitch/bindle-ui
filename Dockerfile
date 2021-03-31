FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go get -t .
RUN go build -o serve .
EXPOSE 9000
CMD ["./serve"]
