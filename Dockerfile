FROM node:12

LABEL description="An image for running the WallHouse OTA Service"
LABEL maintainer="Pete Wall <pete@petewall.net>"

COPY --from=cfplatformeng/needs:latest /usr/local/bin/needs /usr/local/bin/

WORKDIR /usr/src/ota-service
COPY [ "devices.js", "firmware.js", "index.js", "needs.json", "package.json", "package-lock.json", "/usr/src/ota-service/" ]
COPY public /usr/src/ota-service/public
RUN npm install --only=production

EXPOSE 8266

CMD [ "/bin/bash", "-c", "needs check && npm start" ]