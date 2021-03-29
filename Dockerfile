FROM golang AS build

WORKDIR /app
ADD src ./src
WORKDIR src
RUN CGO_ENABLED=0 go build -o /out/server .

FROM scratch

COPY --from=build /out/server /server

ENTRYPOINT ["/server"]