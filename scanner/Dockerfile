FROM centos:centos7

MAINTAINER Lindani Phiri <lphiri@redhat.com>

COPY ./insights-scanner /insights-scanner
COPY ./insights-client-3.0.2-2.fc25.noarch.rpm  /insights-client-3.0.2-2.fc25.noarch.rpm

RUN yum -y update-minimal --security --sec-severity=Important --sec-severity=Critical --setopt=tsflags=nodocs && \
    yum install -y /insights-client-3.0.2-2.fc25.noarch.rpm && \
    yum clean all

ENV EGG=/etc/insights-client/rpm.egg

ENTRYPOINT ["/insights-scanner"]
