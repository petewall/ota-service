FROM node:15-slim

LABEL description="An image for running the WallHouse OTA Service"
LABEL maintainer="Pete Wall <pete@petewall.net>"

ARG TIMEZONE=America/Chicago
RUN ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime

COPY --from=cfplatformeng/needs:latest /usr/local/bin/needs /usr/local/bin/

WORKDIR /usr/src/ota-service
COPY [ "devices.js", "firmware.js", "index.js", "needs.json", "package.json", "package-lock.json", "/usr/src/ota-service/" ]
COPY public /usr/src/ota-service/public
COPY views /usr/src/ota-service/views
RUN npm install --only=production

EXPOSE 8266
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD curl --fail "http://localhost:8266/healthcheck" || exit 1

CMD [ "/bin/bash", "-c", "needs check && npm start" ]
