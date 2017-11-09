FROM centos
EXPOSE 8080

RUN yum install -y epel-release     
RUN yum install -y nodejs

RUN mkdir /usr/bigzam
WORKDIR /usr/bigzam
COPY . /usr/bigzam

RUN npm install
RUN mv node_modules/bootstrap/dist/css/bootstrap.min.css app/public
RUN mv node_modules/bootstrap/dist/css/bootstrap.min.css.map app/public
RUN mv node_modules/bootstrap/dist/js/bootstrap.min.js app/scripts
CMD ["node", "server.js"]