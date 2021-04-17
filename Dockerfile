FROM golang AS build

ENV CGO_ENABLED=1
ENV GO111MOD=on

WORKDIR /app/src
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

ADD src .

RUN go build -tags netgo -a -o /out/server .

FROM debian:stable-slim
COPY --from=build /out/server /server

ENTRYPOINT ["/server"]