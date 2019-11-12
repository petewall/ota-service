FROM node:12-slim

LABEL description="An image suitable for running feature tests"
LABEL maintainer="Pete Wall <pete@petewall.net>"

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' | tee /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    apt-get install -y chromedriver google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*
