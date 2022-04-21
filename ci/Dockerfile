FROM node:18-slim

ENV CI=true
LABEL description="An image suitable for running feature tests"
LABEL maintainer="Pete Wall <pete@petewall.net>"

# Install tools
RUN apt-get update && \
    apt-get install -y curl gnupg2 wget unzip

# Install Google Chrome
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' | tee /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

# Install Chromedriver
RUN export CHROME_VERSION=$(google-chrome --product-version | sed -e "s/\.[0-9]\+$//") && \
    export CHROME_DRIVER_VERSION=$(curl https://chromedriver.storage.googleapis.com/LATEST_RELEASE_${CHROME_VERSION}) && \
    wget -N https://chromedriver.storage.googleapis.com/${CHROME_DRIVER_VERSION}/chromedriver_linux64.zip -P . && \
    unzip chromedriver_linux64.zip && \
    rm chromedriver_linux64.zip && \
    mv -f chromedriver /usr/local/bin/ && \
    chmod +x /usr/local/bin/chromedriver
