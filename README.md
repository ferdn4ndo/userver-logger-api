# uServer-Logger-API

[![Go Report Card](https://goreportcard.com/badge/github.com/ferdn4ndo/userver-logger-api)](https://goreportcard.com/report/github.com/ferdn4ndo/userver-logger-api)
[![test status](https://github.com/ferdn4ndo/userver-logger-api/actions/workflows/run_tests.yml/badge.svg?branch=main "test status")](https://github.com/ferdn4ndo/userver-logger-api/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

<p align="center">
  <img src="https://raw.githubusercontent.com/ferdn4ndo/userver-logger-api/main/static/userver-logger-logo-dark.png?sanitize=true#gh-dark-mode-only" alt="uServer Logger Logo" width="264px"><img src="https://raw.githubusercontent.com/ferdn4ndo/userver-logger-api/main/static/userver-logger-logo-light.png?sanitize=true#gh-light-mode-only" alt="Hurl Logo" width="264px">
</p>

---

**DISCLAIMER: THIS IS A WORK-IN-PROGRESS! WAIT UNTIL THE FIRST RELEASE BEFORE PRODUCTION USE**

---

A RESTful API developed in [Go](https://go.dev/) using the MSC (Model, Service, Controller) architecture to process and catalog `*.log` files (generated as the standard output of running docker containers), allowing queries with pagination and basic search capabilities.

It's part of the [uServer-Logger-Slim](https://github.com/ferdn4ndo/userver-logger-slim) application, a lightweight option for a logging stack in a docker microservices environment. Compared to an ELK scenario, it would replace both Elasticsearch & Logstash services more lightly (although losing several capabilities too). The goal of this service is to be part of a Log Management tool light enough to run in a multi-container environment inside a low-to-medium EC2 instance.

Built using [go-chi](https://github.com/go-chi/chi), [gorm](https://github.com/go-gorm/gorm), [go-sqlite3](https://github.com/mattn/go-sqlite3). The development version also depends on [air](https://github.com/cosmtrek/air).

## Configuration (Environment)

To configure the container, copy the `.env.template` file to `.env` (you can use the command below). An explanation of each of the variables is also available in this section.

```
cp .env.template .env
```

Then edit the file to tweak the settings as you wish before running the container.

### Variables

* **VIRTUAL_HOST**: The virtual hostname to use if you're running the container behind a reverse proxy. (Default: `[EMPTY]`)
* **LETSENCRYPT_HOST**: The virtual hostname to use in the SSL certificate generation by [Let's Encrypt](https://letsencrypt.org/) if you're running the container behind a reverse proxy. (Default: `[EMPTY]`)
* **LETSENCRYPT_EMAIL**: The hostmaster e-mail to use in the SSL certificate generation by [Let's Encrypt](https://letsencrypt.org/) if you're running the container behind a reverse proxy. (Default: `[EMPTY]`)
* **BASIC_AUTH_USERNAME**: The username to use in the Basic Authentication of the API endpoints. (Default: `[EMPTY]`) **[REQUIRED]**
* **BASIC_AUTH_PASSWORD**: The password to use in the Basic Authentication of the API endpoints. (Default: `[EMPTY]`) **[REQUIRED]**
* **SERVER_PORT**: The port used to expose the API service. (Default: `5000`)
* **LOG_FILES_FOLDER**: The location of the log files to be watched. (Default: `/log_files`)
* **TMP_FOLDER**: The location of the temporary files created while running the service. (Default: `/go/src/github.com/ferdn4ndo/userver-logger-api/tmp`)
* **DATA_FOLDER**: The location of the temporary files created while running the service. (Default: `/go/src/github.com/ferdn4ndo/userver-logger-api/data`)
* **FIXTURE_FOLDER**: The location of the fixture files for preloading internal service data. (Default: `/go/src/github.com/ferdn4ndo/userver-logger-api/fixture`)
* **DATABASE_FILE**: The filename of the SQLite database file (inside the `data` folder) to store the parsed log entries. (Default: `sqlite.db`)
* **TEST_DATABASE_FILE**: The filename of the SQLite database file (inside the `data` folder) to use during the tests. (Default: `test.sqlite.db`)
* **EMPTY_DATABASE_FILE**: The filename of the SQLite database file (inside the `fixture` folder) without any table, to be used when preparing a new test environment. (Default: `empty.sqlite.db`)

TEST_DATABASE_FILE=test.sqlite.db
EMPTY_DATABASE_FILE=empty.sqlite.db
## How to run

### In Production

To build the image:

```
# Navigate to the project folder and run
docker build -f ./Dockerfile --tag userver-logger-api:latest .
```

For a single container run (that will expose port `5000` by default):

```
# Assuming .env file is at the current location
docker run -d --rm -e 5000 -v "$DATA_DIR":/data --env-file ./env "$CONTAINER_NAME":local
```

Docker-compose version (will build and run):

```
# Navigate to the project folder and:
docker compose -f docker-compose.yml up --build
```

### In Development

For development purposes, we recommend running the `docker compose` command.

The project has a hot-reload mechanism using [air](https://github.com/cosmtrek/air).

```
# Navigate to the project folder and:
docker compose -f docker-compose.dev.yml up --build
```

## Endpoints

* **GET /health**: provides basic health checking, retrieving a 200 OK (and internally registering a heartbeat) when up & running; This endpoint requires NO authentication;

* **GET /log-entries**: provides basic health checking, retrieving a 200 OK (and internally registering a heartbeat) when up & running; This endpoint requires Basic Authentication (credentials configured in the environment variables);

    * Query parameters:
        
        * `size`: Number of results per page (min: 1, max: 1000, default: 100);
        * `offset`: Number of results to skip before starting the page (min: 0, default: 0);
        * `producer`: The name of the producer to filter the results (exact match);
        * `message`: The message (or part of it) to search in the logs (`LIKE %keyword%` match);

    * Response schema:

        * `items`: The array of entry logs of that page;
        * `total_count`: The total (non-paginated) number of results;
        * `page_count`: The number of results on that specific page;
        * `next_page_offset`: The offset value to fetch the results from the next page (if it's equal to the requested offset, it means you're in the last page);
        * `previous_page_offset`: The offset value to fetch the results from the previous page (if it's equal to the requested offset, it means you're in the first page);

* **POST /log-entries**: creates a log entry. It will retrieve a `201 Created` status code with the created log entry in case of success, or a 4xx with the error message otherwise. This endpoint requires Basic Authentication (credentials configured in the environment variables);

    * Request body:

        * `producer`: The name of the producer of the log entry;
        * `message`: The message (content) of the log entry;

    * Response schema:

        *SAME AS IN 'GET /log-entries/{id}'*

* **GET /log-entries/{id}**: retrieves a single log entry. It will retrieve a `200 Ok` status code with the requested log entry in case of success, or a 4xx with the error message otherwise. This endpoint requires Basic Authentication (credentials configured in the environment variables);

    * URL parameters:

        * `producer`: The name of the producer of the log entry;
        * `message`: The message (content) of the log entry;

    * Response schema:

        * `id`: The unique ID of the log entry;
        * `producer`: The name of the producer of the log entry;
        * `message`: The message (content) of the log entry;
        * `created_at`: The timestamp (in ISO 8601 format) when the log entry was registered in the application (note that this is different then the log creation timestamp, which should be part of the log message).

* **PUT /log-entries/{id}**: updates a single log entry. It will retrieve a `200 Ok` status code with the updated log entry in case of success, or a 4xx with the error message otherwise. This endpoint requires Basic Authentication (credentials configured in the environment variables);

    * Request body:

        * `producer`: The name of the producer of the log entry;
        * `message`: The message (content) of the log entry;

    * Response schema:

        *SAME AS IN 'GET /log-entries/{id}'*

* **DELETE /log-entries/{id}**: deletes a single log entry. It will retrieve a `204 No Content` status code with an empty body on success, or a 4xx with the error message otherwise. This endpoint requires Basic Authentication (credentials configured in the environment variables);

## Testing

To run the test suite for CI/CD pipelines, run:

```
docker exec -it userver-logger-api sh -c "./scripts/run_all_tests.sh"
```

If you want to have a coverage report of the tests, run:

```
docker exec -it userver-logger-api sh -c "./scripts/run_all_tests_with_coverage.sh"
```

If you only want the number (float) of the coverage percentage, run:

```
docker exec -it userver-logger-api sh -c "./scripts/get_test_coverage_percentage.sh"
```


## F.A.Q.

### 1 - Why using the SQLite driver?
R: Because the logging container aims to be one of the very first services started on a web application stack. It should avoid any other later service dependency, and it can be potentially used to monitor the main database container (therefore not being able to depend on it).

### 2 - I found a bug / I want a new feature. What should I do?
R: Open an issue in this repository. Please describe your request as detailed as possible (remember to attach binary/big files externally), and wait for feedback. If you're familiar with software development, feel free to open a Pull Request with the suggested solution. Contributions are welcomed!

## License

This application is distributed under the [MIT](https://github.com/ferdn4ndo/userver-logger-api/blob/main/LICENSE) license.

## Contributors

[ferdn4ndo](https://github.com/ferdn4ndo)

Any help is appreciated! Feel free to review / open an issue / fork / make a PR.

