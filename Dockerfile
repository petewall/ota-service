FROM node:12

COPY --from=cfplatformeng/needs:latest /usr/local/bin/needs /usr/local/bin/

WORKDIR /usr/src/ota-service
COPY [ "index.js", "needs.json", "package.json", "package-lock.json", "/usr/src/ota-service/" ]
RUN npm install --only=production 

CMD [ "/bin/bash", "-c", "needs check && npm start" ]