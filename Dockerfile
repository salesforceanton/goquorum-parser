FROM golang:1.22.2

SHELL ["/bin/bash", "-c"]
WORKDIR /app

# build cache
RUN GOPROXY=direct GOSUMDB=off
COPY go.mod go.sum /app/
COPY Makefile /app/

# install deps
RUN cd /app && go mod download
# add recent changes
ADD . /app/

RUN mv docker/services/docker-entrypoint.sh / && \
    chmod +x /docker-entrypoint.sh

ENTRYPOINT [ "/docker-entrypoint.sh" ]
