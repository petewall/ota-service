FROM ubuntu AS get-chromedriver
RUN apt-get update && \
    apt-get install -y unzip wget
RUN wget https://chromedriver.storage.googleapis.com/81.0.4044.69/chromedriver_linux64.zip
RUN unzip chromedriver_linux64.zip
RUN chmod +x chromedriver

FROM node:15-slim

ENV CI=true
LABEL description="An image suitable for running feature tests"
LABEL maintainer="Pete Wall <pete@petewall.net>"

COPY --from=get-chromedriver chromedriver /usr/bin/chromedriver
RUN apt-get update && \
    apt-get install -y gnupg2 wget && \
    wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' | tee /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*
