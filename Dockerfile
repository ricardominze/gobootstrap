FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY . . 
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./gobootstrap

FROM scratch AS runner

WORKDIR /
COPY --from=BUILDER /build/.env .
COPY --from=BUILDER /build/gobootstrap .
COPY --from=BUILDER /build/infra/db/migrations/ ./infra/db/migrations/

EXPOSE 8000
ENTRYPOINT [ "/gobootstrap" ]