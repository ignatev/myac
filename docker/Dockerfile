FROM ignatev/go-build-env-git AS build-env
ADD . /src
WORKDIR /src
RUN go build -o myac

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY config.yml /app/
COPY --from=build-env /src/myac /app/
ENTRYPOINT ./myac
