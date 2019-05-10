FROM golang:1.11

WORKDIR $GOPATH/src/github.com/saromanov/pinger
COPY ..
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/pinger

FROM scratch
COPY --from=build-env /go/bin/pinger /go/bin/pinger
ENTRYPOINT ["/go/bin/pinger"]