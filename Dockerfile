FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /go-backend

FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /

COPY --from=builder /go-backend /go-backend

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-backend"]
