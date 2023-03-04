FROM golang:1.19.3-alpine as builder

WORKDIR /work
ENV GOPROXY https://goproxy.cn,direct

COPY . .
RUN go mod tidy
RUN go build -o=yzh_music_api .


FROM alpine as test_app
# FROM scratch as app

COPY --from=builder /work/yzh_music_api .

EXPOSE 8080
ENTRYPOINT ["./yzh_music_api"]