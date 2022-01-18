FROM golang:1-alpine as go_build
WORKDIR /app
COPY ./src ./
RUN go build ddns.go

# IMAGE ASSEMBLY
FROM alpine:latest
WORKDIR /app
ARG APP_NAME="main"
COPY --from=go_build /app/${APP_NAME} /app/${APP_NAME}

ENTRYPOINT [ "/app/${APP_NAME}" ]