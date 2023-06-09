FROM golang:1.20.4-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

ENV PORT=8080

EXPOSE 8080

RUN go build -tags netgo -ldflags=-s -ldflags=-w -o /bin/gedis

ENTRYPOINT [ "/bin/gedis" ]