FROM golang:1.20-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY main.go ./
COPY html ./html

RUN go build -o server main.go

FROM alpine:3.19
WORKDIR /app

COPY --from=build /app/server ./server
COPY --from=build /app/html ./html

ENV PORT=8080
EXPOSE 8080

CMD ["/app/server"]
