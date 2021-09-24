FROM golang:1.16-buster AS build

RUN go version

COPY . /github.com/Hudayberdyyev/weather-api/
WORKDIR /github.com/Hudayberdyyev/weather-api/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o weather cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=build /github.com/Hudayberdyyev/weather-api/weather .
COPY --from=build /github.com/Hudayberdyyev/weather-api/configs configs/

CMD ["./weather"]