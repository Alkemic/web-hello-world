FROM golang:1.13
WORKDIR /go/src/github.com/Alkemic/web-hello-world/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o web-hello-world .

FROM scratch
COPY --from=0 /go/src/github.com/Alkemic/web-hello-world/web-hello-world .
CMD ["./web-hello-world"]