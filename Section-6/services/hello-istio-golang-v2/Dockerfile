FROM golang:1.13.6-alpine

ENV PORT 8080
EXPOSE 8080

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["hello-istio-golang-v2"]
