FROM node:6
EXPOSE 8080

RUN mkdir /usr/bigzam
WORKDIR /usr/bigzam
COPY . /usr/bigzam

RUN npm install

COPY node_modules/bootstrap/dist/css/bootstrap.min.css app/public
COPY node_modules/bootstrap/dist/css/bootstrap.min.css.map app/public
COPY node_modules/bootstrap/dist/js/bootstrap.min.js app/public
CMD ["node", "server.js"]