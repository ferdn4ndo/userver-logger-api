FROM golang@sha256:344193a70dc3588452ea39b4a1e465a8d3c91f788ae053f7ee168cebf18e0a50
LABEL maintaner="Fernando Constantino <const.fernando@gmail.com>"

# To update the base alpine image, please refer to
# https://github.com/docker-library/repo-info/blob/master/repos/golang/remote/1-alpine.md
# and get the latest sha256 for 'linux; amd64'

ARG BUILD_DATE
ARG BUILD_VERSION
ARG VCS_REF

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="ferdn4ndo/userver-logger-api"
LABEL org.label-schema.description="An Alpine-based Go API to process and catalog `*.log` file, allowing queries with pagination and basic search capabilities."
LABEL org.label-schema.vcs-url="https://github.com/ferdn4ndo/userver-logger-api"
LABEL org.label-schema.usage="/README.md"
LABEL org.label-schema.vcs-ref=$VCS_REF
LABEL org.label-schema.version=$BUILD_VERSION
LABEL org.label-schema.docker.cmd="docker run --rm --env-file ./.env ferdn4ndo/userver-logger-api"
LABEL org.label-schema.docker.cmd.devel="docker compose -f docker-compose.dev.yml up --build"
LABEL org.label-schema.docker.cmd.test="docker exec -it userver-logger-api sh -c ./scripts/run_all_tests.sh"

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
