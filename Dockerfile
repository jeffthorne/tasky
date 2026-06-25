# Building the binary of the App
FROM golang:1.19 AS build

WORKDIR /go/src/tasky
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /go/src/tasky/tasky


FROM alpine:3.17.0 AS release

ARG exposed_username=CHANGEME

# --- WIZ EXERCISE REQUIREMENT ---
RUN mkdir --parents /app && \
    echo "${exposed_username}" > /app/wizexercise.txt

WORKDIR /app
COPY --from=build  /go/src/tasky/tasky .
COPY --from=build  /go/src/tasky/assets ./assets
EXPOSE 8080
ENTRYPOINT ["/app/tasky"]

