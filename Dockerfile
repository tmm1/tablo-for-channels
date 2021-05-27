FROM alpine as build
RUN apk add go git
WORKDIR /usr/src/tablo-for-channels
ADD go.* *.go .
ADD tablo tablo
RUN go build

FROM alpine
WORKDIR /usr/src/app
ENV GIN_MODE=release
ENV PORT=80
COPY --from=build /usr/src/tablo-for-channels/tablo-for-channels .
ENTRYPOINT ./tablo-for-channels -ip $TABLO_IP