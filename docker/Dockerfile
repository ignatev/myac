FROM ignatev/go-build-env-git:latest AS build-env
MAINTAINER 'Iskander Ignatev, iskander.ignatev@gmail.com'
COPY . /src
WORKDIR /src
RUN go build -o myac

FROM alpine:3.8
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY config.yml /app/
COPY --from=build-env /src/myac /app/
ENTRYPOINT ./myac
