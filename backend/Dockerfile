FROM golang:1.23-alpine AS dependencies

WORKDIR /app

COPY ./go.mod .

COPY ./go.sum .

RUN go mod download


FROM dependencies AS build

COPY ./ .

RUN go build -o /app


FROM dependencies AS test

COPY ./ .

CMD ["go", "test", "-v"]

FROM alpine:3.15 AS run

WORKDIR /app

COPY --from=build /app/area51 .

EXPOSE 8080

CMD ["/app/area51"]
