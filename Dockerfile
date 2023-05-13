FROM golang@sha256:0a03b591c358a0bb02e39b93c30e955358dadd18dc507087a3b7f3912c17fe13

# To update the base alpine image, please refer to
# https://github.com/docker-library/repo-info/blob/master/repos/golang/remote/1-alpine.md
# and get the latest sha256

LABEL maintaner="Fernando Constantino <const.fernando@gmail.com>"

# Install GCC + git + SSL ca certificates
# GCC is required to build the sqlite3 dependency
# Git is required for fetching the dependencies
RUN apk update \
    && apk add \
      git \
      gcc \
      musl-dev \
      curl \
      sqlite \
      coreutils \
    && apk upgrade \
    && rm -rf /var/cache/apk/*

# Create appuser
ENV USER=appuser
ENV UID=1000

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /go/src/github.com/ferdn4ndo/userver-logger-api/

COPY ./ /go/src/github.com/ferdn4ndo/userver-logger-api/

RUN git config --global --add safe.directory /go/src/github.com/ferdn4ndo/userver-logger-api

RUN go mod download \
    github.com/go-chi/chi/v5 \
    github.com/go-chi/docgen \
    github.com/go-chi/render \
    gorm.io/gorm \
    github.com/mattn/go-sqlite3 \
    gorm.io/driver/sqlite \
    github.com/ajg/form

RUN go mod tidy

RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -o /go/bin/server . \
    && chmod +x /go/bin/server

EXPOSE 5555

HEALTHCHECK --interval=30s --timeout=10s --retries=6 --start-period=5s CMD curl -s -f http://localhost:5555/health

ENTRYPOINT ["/go/bin/server"]
