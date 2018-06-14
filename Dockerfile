FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o myac

FROM alpine
WORKDIR /app
COPY --from=build-env /src/myac /app/
ENTRYPOINT ./myac
