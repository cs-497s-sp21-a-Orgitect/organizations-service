FROM golang AS build

ENV CGO_ENABLED=1

WORKDIR /app
ADD src ./src
WORKDIR src
RUN go get github.com/mattn/go-sqlite3
RUN go build -tags netgo -a -o /out/server .

FROM scratch

COPY --from=build /out/server /server

ENTRYPOINT ["/server"]