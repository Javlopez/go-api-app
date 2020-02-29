# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:latest as builder

ARG port=8080
ENV ENV_PORT=${port}


LABEL maintainer="Javier Lopez Lopez <sjavierlopez@gmail.com>"


WORKDIR /go/src/go-lana
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/go-lana .
EXPOSE ${port}

# Command to run the executable
CMD ["sh", "-c", "./main --port=${ENV_PORT}"] 