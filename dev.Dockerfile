FROM cosmtrek/air

LABEL maintaner="Fernando Constantino <const.fernando@gmail.com>"

ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /go/src/github.com/ferdn4ndo/userver-logger-api

RUN apt-get update && apt-get install -y git gcc musl-dev ca-certificates curl sqlite3 coreutils && update-ca-certificates

COPY ./ /go/src/github.com/ferdn4ndo/userver-logger-api/

RUN go mod download github.com/go-chi/chi/v5 github.com/go-chi/docgen github.com/go-chi/render gorm.io/gorm gorm.io/driver/sqlite

RUN go mod verify

EXPOSE 5000

HEALTHCHECK --interval=30s --timeout=10s --retries=6 --start-period=5s CMD curl -s -f http://localhost:5000/health

ENTRYPOINT ["./entrypoint.dev.sh"]
