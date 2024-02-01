FROM golang:1.21 as builder

WORKDIR /go/src/module

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# See https://stackoverflow.com/questions/60400478/golang-based-docker-image-build-works-but-not-scratch-based-image
RUN CGO_ENABLED=0 go build -o /go/bin/server ./cmd/server/main.go

EXPOSE 8080
CMD ["/go/bin/service"]

FROM scratch

COPY --from=builder /go/bin/server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
