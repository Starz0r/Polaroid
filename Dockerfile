# Set Builder Image
FROM golang:1.13.15-alpine3.12 as builder

# Add Build Dependencies and Working Directory
RUN apk --no-cache add build-base git tar wget
RUN wget https://github.com/go-task/task/releases/download/v2.0.0/task_linux_amd64.tar.gz
RUN tar zxvf task_linux_amd64.tar.gz && cp ./task /usr/bin/
RUN chmod +x /usr/bin/task
RUN wget https://github.com/gobuffalo/packr/releases/download/v1.30.1/packr_1.30.1_linux_amd64.tar.gz
RUN tar zxvf packr_1.30.1_linux_amd64.tar.gz && cp ./packr /usr/bin/
RUN chmod +x /usr/bin/packr
RUN mkdir /build
ADD . /build/
WORKDIR /build

# Compile
ENV GO111MODULE=on
RUN task build BUILD_OUTNAME=service

# Move to Base Image and Run
FROM alpine:3.12.0
RUN apk update && apk upgrade && apk add ca-certificates
COPY --from=builder /build/service /app/
WORKDIR /app
CMD ["./service"]