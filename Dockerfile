FROM golang:1.6-onbuild

RUN apt-get update -qq && DEBIAN_FRONTEND=noninteractive apt-get -q -y install libpq-dev

RUN mkdir -p /app
WORKDIR /app