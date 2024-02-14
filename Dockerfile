FROM golang AS build

WORKDIR /go/src/bot
ADD . /go/src/bot

RUN go build -o bot

FROM gcr.io/distroless/base-debian12

COPY --from=build /go/src/bot/bot /

CMD ["/bot"]
