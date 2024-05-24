FROM golang:alpine as builder
#FROM golang:1.16 AS builder

ARG APP_RELATIVE_PATH
COPY . /src
WORKDIR /src/${APP_RELATIVE_PATH}

RUN echo "Building the Docker image..."
RUN echo ${APP_RELATIVE_PATH}
RUN echo pwd

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
	&& mkdir -p bin/ \
    && go mod tidy \
    && go build -o ./bin/ ./cmd/main.go

FROM alpine AS runner
ARG APP_RELATIVE_PATH

WORKDIR /src/${APP_RELATIVE_PATH}/cmd

COPY --from=builder /src/${APP_RELATIVE_PATH}/bin/ .
COPY --from=builder /src/${APP_RELATIVE_PATH}/config/ ../config/

# expose 不然在容器外没有可以访问服务的端口
# httpserver
#EXPOSE 10000
# userservice
EXPOSE 10001
# videoservice
#EXPOSE 10002
# socialservice
#EXPOSE 10003
VOLUME /data/conf

#CMD ["./main"]