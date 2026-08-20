FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /teapot ./

# ---

FROM scratch

COPY --from=build /teapot /teapot

EXPOSE 4180
CMD ["/teapot"]
