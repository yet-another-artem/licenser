FROM golang:alpine AS build-stage
WORKDIR /app
COPY . .
RUN go build -o main

FROM bash:4.1.17-alpine3.22 AS run-stage
WORKDIR /app
COPY --from=build-stage /app/main .
HEALTHCHECK NONE
CMD [ "./main", "data", "yet.another.artem" ]