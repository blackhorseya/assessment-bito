# Bito assessment

## Introduction

This repository contains the code for the Bito assessment.

## How to use

### On a Local Machine

1. Clone the repository

    ```shell
    git clone git@github.com:blackhorseya/assessment-bito.git
    ```

2. Install the dependencies

    ```shell
    go mod download
    ```

3. Show help message

    ```shell
    go run adapter --help
    ```

4. Run HTTP server with the following command

    ```shell
    go run adapter start api
    ```

### With Docker

1. Pull the Docker image

    ```shell
    docker pull ghcr.io/blackhorseya/assessment-bito:latest
    ```

2. Show help message

    ```shell
    docker run -it --rm ghcr.io/blackhorseya/assessment-bito:latest --help
    ```

3. Run the Docker container with HTTP server

   ```shell
   docker run -it --rm -p 30000:30000 ghcr.io/blackhorseya/assessment-bito:latest start api
   ```
