FROM golang:1.15.2
COPY . pushbot
ENV GO111MODULE=on

RUN cd pushbot &&  go get ./...  &&  make clean &&  make

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/pushbot/dist/ .
RUN ls
ENTRYPOINT ["/root/pushbot"]

