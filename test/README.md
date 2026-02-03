# acme-dns E2E Testing Suite

This directory contains the end-to-end (E2E) testing suite for `acme-dns`. The suite runs in a containerized environment to ensure a consistent and isolated test execution.

## Overview

The E2E suite consists of:
- A Dockerized `acme-dns` server.
- A Python-based `tester` container that performs API and DNS operations.
- A GitHub Actions workflow for CI/CD integration.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Running Locally

To run the full E2E suite locally, execute the following command from the root of the repository:

```bash
docker compose -f test/e2e/docker-compose.yml up --build --abort-on-container-exit
```

The `tester` container will return an exit code of `0` on success and `1` on failure.

## Test Flow

The `tester.py` script follows these steps:
1.  **Wait for Ready**: Polls the `/health` endpoint until the API is available.
2.  **Account Registration**: Registers a NEW account at `/register`.
3.  **TXT Update**: Performs TWO sequential updates to the TXT records of the newly created subdomain.
4.  **DNS Verification**: Directly queries the `acme-dns` server (on port 53) used in the test to verify that the TXT records have been correctly updated and are resolvable.

## CI/CD integration

The tests are automatically run on every push and pull request to the `master` branch via GitHub Actions. The workflow configuration can be found in `.github/workflows/e2e.yml`.
