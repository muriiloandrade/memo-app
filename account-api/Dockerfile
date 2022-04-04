# Create base image with golang
FROM golang:1.18-bullseye as base

# Create another stage called "dev" based off of our "base" stage (so we have golang available)
FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Run the air command in the directory where our code will live
WORKDIR /opt/app/account-api

CMD ["air"]

FROM base as built

WORKDIR /go/app/api
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/run ./*.go

FROM busybox

COPY --from=built /tmp/run .

EXPOSE 3000

ENV GIN_MODE=release

CMD ["./run"]