FROM golang:alpine AS build-stage
WORKDIR /app
COPY . .
RUN go build -o main

FROM gcr.io/distroless/static-debian12 AS run-stage
WORKDIR /app
COPY --from=build-stage /app/main .
ENTRYPOINT [ "./main", "data", "yet.another.artem" ]