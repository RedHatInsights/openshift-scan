# FROM registry.access.redhat.com/devtools/go-toolset-7-rhel7 as builder
FROM golang:1.9 AS builder
MAINTAINER Jeremy Crafts <jcrafts@redhat.com>

# RUN sudo subscription-manager register --username=rhn-support-jcrafts --password=ISeeYouConnor subscription-manager attach --pool=8a85f9833e1404a9013e3cddf99305e6

# COPY ./insights-client-3.0.2-2.fc25.noarch.rpm  /insights-client-3.0.2-2.fc25.noarch.rpm

# RUN yum -y update-minimal --security --sec-severity=Important --sec-severity=Critical --setopt=tsflags=nodocs && \
#     yum install -y golang git && \
#     yum install -y /insights-client-3.0.2-2.fc25.noarch.rpm && \
#     yum clean all

ENV GOPATH=/go
RUN mkdir -p /go/src/github.com/RedHatInsights/insights-ocp/scanner
WORKDIR /go/src/github.com/RedHatInsights/insights-ocp/scanner

COPY ./insights-scanner.go ./insights-scanner.go
COPY vendor ./vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o /insights-scanner

ENV EGG=/etc/insights-client/rpm.egg

FROM scratch
COPY --from=builder /insights-scanner /insights-scanner
# COPY /insights-scanner /insights-scanner
COPY ./insights-client-3.0.2-2.fc25.noarch.rpm  /insights-client-3.0.2-2.fc25.noarch.rpm
ENTRYPOINT ["./insights-scanner"]
