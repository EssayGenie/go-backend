FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go-backend ./

RUN go mod download


RUN CGO_ENABLED=0 go build -o /go-backend

FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /

COPY --from=builder /go-backend /go-backend

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/go-backend"]
