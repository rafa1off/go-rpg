FROM golang:1.21-alpine AS builder

WORKDIR /code

COPY go.mod go.sum /code/

RUN go mod download

COPY . .

RUN go build -o ./exec/exec main.go


FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /code/exec /exec

EXPOSE 8000

ENTRYPOINT [ "exec/exec" ]