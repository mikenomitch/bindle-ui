FROM golang

WORKDIR /app

COPY ./bindle-linux ./bin/bindle
COPY ./nomad-linux ./bin/nomad

ENV PATH="./bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go get -t .
RUN go build -o serve .
EXPOSE 9000
CMD ["./serve"]
