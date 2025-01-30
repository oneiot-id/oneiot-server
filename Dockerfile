FROM golang:1.23

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./app
RUN go build -v -o /usr/local/bin/app


EXPOSE 8000

CMD ["app"]
