FROM golang:1.24-alpine AS build
WORKDIR /go/app
ADD . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/user ./cmd/main.go

FROM scratch
COPY --from=build /go/app/bin/user /go/app/user

CMD ["/go/app/user"]
EXPOSE 8080