# Bito assessment

[![Go](https://github.com/blackhorseya/assessment-bito/actions/workflows/go.yaml/badge.svg)](https://github.com/blackhorseya/assessment-bito/actions/workflows/go.yaml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=blackhorseya_assessment-bito&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=blackhorseya_assessment-bito)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=blackhorseya_assessment-bito&metric=coverage)](https://sonarcloud.io/summary/new_code?id=blackhorseya_assessment-bito)
![GitHub Release](https://img.shields.io/github/v/release/blackhorseya/assessment-bito)

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
