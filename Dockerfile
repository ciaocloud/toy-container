#FROM ubuntu:latest
FROM golang:latest
LABEL maintainer="xingw@splunk.com"
ENV APP_ROOT="/toy-container" \
    VERSION="0.6"
WORKDIR ${APP_ROOT}
COPY . ${APP_ROOT}
#ENTRYPOINT ["./START.sh"]
#CMD ["run", "main.go", "run"]