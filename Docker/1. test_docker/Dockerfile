FROM golang:1.18-buster AS gobuilder
COPY hello.go /app/hello.go
WORKDIR /app
RUN go build -o hello hello.go

FROM scratch
COPY --from=gobuilder /app/hello /
CMD ["/hello"]

#https://nayoungs.tistory.com/entry/Docker-Docker%EC%97%90-Go-%EB%B0%B0%ED%8F%AC%ED%95%98%EA%B8%B0
