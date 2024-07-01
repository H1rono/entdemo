FROM golang:1.22.4-bookworm

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD [ "go", "run", "main.go" ]
