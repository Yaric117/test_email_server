FROM golang:1.18

# dependencies
RUN go install github.com/cosmtrek/air@latest
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate


# init
RUN mkdir -p /app/src
WORKDIR /app/src

# vendor
COPY src/go.mod src/go.sum /app/src/
RUN  go mod download

# src
COPY . /app

# build
RUN go build -o /go/bin/app

# run
CMD ["/go/bin/app"]