FROM golang:1.13.8 as build

MAINTAINER "Xander Guzman <xander.guzman@xanderguzman.com>"

WORKDIR /go/src/github.com/theshadow/hived
COPY . .

ENV GO111MODULE=on
ENV PYTHONPATH=/usr/lib/python2.7/dist-packages

ENV OPT_EDITOR=no

# System
RUN apt-get update && apt-get upgrade -y

# Tools
RUN apt-get install -y python-pip python-sphinx zip
RUN go get github.com/readthedocs/godocjson

# Libraries
RUN pip install sphinx-autoapi sphinxcontrib-golangdomain

RUN make

FROM nginx:alpine

COPY --from=build /go/src/github.com/theshadow/hived/_build/docs /usr/share/nginx/html
