# Bito assessment

## How to run the project

On local

```shell
go run ./adapter --help
```

On Docker

```shell
docker build -t assessment-bito:latest .
```

```shell
docker run -it --rm -p 30000:30000 assessment-bito:latest start api
```
