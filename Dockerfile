FROM golang AS build

ENV CGO_ENABLED=0
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/sqlite

WORKDIR /app
ADD src ./src
WORKDIR src
RUN go mod download gorm.io/gorm
RUN go mod download gorm.io/driver/sqlite
RUN go build -tags netgo -a -o /out/server .

FROM alpine

COPY --from=build /out/server /server

ENTRYPOINT ["/server"]