# Step 1 - build
FROM golang:1.21 AS build

ENV CGO_ENABLED=0

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -ldflags "-s -w" -o /gh-actions-go

# Step 2 - UPX Compression
FROM alpine:3.17 AS compress

RUN apk add upx
COPY --from=build /gh-actions-go /
RUN upx /gh-actions-go

# Step 3 - use distroless to reduce image size
FROM gcr.io/distroless/static-debian11

COPY --from=compress gh-actions-go /

ENTRYPOINT ["/gh-actions-go"]
