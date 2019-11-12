FROM node:12-slim

LABEL description="An image suitable for running feature tests"
LABEL maintainer="Pete Wall <pete@petewall.net>"

RUN apt-get update && \
    apt-get install -y chromedriver && \
    rm -rf /var/lib/apt/lists/*
