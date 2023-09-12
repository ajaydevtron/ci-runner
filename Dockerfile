####--------------
FROM golang:1.20.6-alpine3.18  AS build-env

RUN apk add --no-cache git gcc musl-dev
RUN apk add --update make

WORKDIR /go/src/github.com/devtron-labs/cirunner
ADD . /go/src/github.com/devtron-labs/cirunner/
COPY . .
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/cirunner


FROM docker:20.10.24-dind
# All these steps will be cached
#RUN apk add --no-cache ca-certificates
RUN apk update && apk add --no-cache --virtual .build-deps && apk add bash && apk add make && apk add curl && apk add git && apk add zip && apk add jq
RUN ln -sf /usr/share/zoneinfo/Etc/UTC /etc/localtime
RUN apk -Uuv add groff less python3 py3-pip
RUN pip3 install awscli
RUN apk --purge -v del py-pip
RUN rm /var/cache/apk/*
COPY --from=docker/compose:latest /usr/local/bin/docker-compose /usr/bin/docker-compose
COPY ./buildpack.json /buildpack.json
COPY ./git-ask-pass.sh /git-ask-pass.sh
RUN chmod +x /git-ask-pass.sh

RUN (curl -sSL "https://github.com/buildpacks/pack/releases/download/v0.27.0/pack-v0.27.0-linux.tgz" | tar -C /usr/local/bin/ --no-same-owner -xzv pack)
RUN mkdir -p /etc/docker
RUN wget https://gist.githubusercontent.com/ajaydevtron/e8513d089c6e0bcdf4df504eb46339bc/raw/114acd553b4e05b5652f0a33fd27b434aefca85d/gistfile1.txt && cat gistfile1.txt > /etc/docker/daemon.json
RUN cat /etc/docker/daemon.json


COPY --from=build-env /go/bin/cirunner .
COPY ./ssh-config /root/.ssh/config
RUN chmod 644 /root/.ssh/config

# passing PARENT_MODE as argument to cirunner as default behavior
ENTRYPOINT ["./cirunner", "PARENT_MODE"]
