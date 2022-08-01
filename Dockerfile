FROM scratch

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/Mistsink/kuwo-api
COPY . $GOPATH/src/github.com/Mistsink/kuwo-api

EXPOSE 8080
ENTRYPOINT ["./kuwo-api"]