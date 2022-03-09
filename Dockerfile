FROM golang:1.18-rc-alpine AS build

WORKDIR /go/src/api
ADD . /go/src/api

RUN go build -o api

FROM gcr.io/distroless/base-debian11

COPY --from=build /go/bin/api /

CMD ["/api"]
