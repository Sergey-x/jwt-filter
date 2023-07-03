ARG GO_VERSION=1.18.1

FROM golang:${GO_VERSION}-alpine
WORKDIR /jwt-filter

COPY . .
RUN go mod tidy
RUN go build -buildvcs=false -o jwt-filter .
RUN chmod 755 ./jwt-filter

ENTRYPOINT ./jwt-filter
