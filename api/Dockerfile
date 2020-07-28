FROM golang

FROM golang:1.13-buster as build

WORKDIR /go/src/app
COPY go.mod /go/src/app
COPY go.sum /go/src/app
RUN go mod download

COPY . /go/src/app
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
CMD ["/app"]