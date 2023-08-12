
##################################
##         bulid server         ##
##################################
FROM golang:1.20-alpine as build
WORKDIR /app
COPY . .

RUN go env -w GO111MODULE=on \
        && go env -w GOPROXY=https://goproxy.cn,direct \
        && go env -w CGO_ENABLED=0

RUN go build -o server .



####################################
##            deploy              ##
####################################
FROM alpine:latest
WORKDIR /app

COPY --from=build   /app/server  ./
COPY --from=build   /app/deploy  ./deploy

EXPOSE 8000
ENTRYPOINT ["/app/server"]

