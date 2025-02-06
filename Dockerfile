FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./gobootstrap

FROM scratch AS runner

WORKDIR /
COPY --from=builder /build/.env .
COPY --from=builder /build/gobootstrap .
COPY --from=builder /build/infra/db/migrations/ ./infra/db/migrations/

EXPOSE 8000
ENTRYPOINT [ "/gobootstrap" ]
