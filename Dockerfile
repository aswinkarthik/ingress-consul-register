FROM golang:1.9.2 as source
RUN mkdir -p /go/src/github.com/aswinkarthik93
COPY . /go/src/github.com/aswinkarthik93/ingress-consul-register
RUN cd /go/src/github.com/aswinkarthik93/ingress-consul-register \
    && go build -o ingress-consul-register

FROM debian:stretch
RUN apt-get update -y \
    && apt-get install ca-certificates -y \
    && update-ca-certificates --verbose
COPY --from=source /go/src/github.com/aswinkarthik93/ingress-consul-register/ingress-consul-register .
CMD ["./ingress-consul-register", "start"]
