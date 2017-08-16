FROM lphiri/insights-client
MAINTAINER Lindani Phiri <lphiri@redhat.com>

RUN yum update -y && \
    yum install -y golang && \
    yum clean all


COPY .  /go/src/github.com/RedHatInsights/insights-ocp

RUN GOBIN=/usr/bin \
    GOPATH=/go \
    CGO_ENABLED=0 \
    go install -a -installsuffix cgo /go/src/github.com/RedHatInsights/insights-ocp/cmd/insights-ocp.go && \
    mkdir -p /var/lib/insights-ocp

EXPOSE 9000

WORKDIR /var/lib/insights-ocp

ENTRYPOINT ["/usr/bin/insights-ocp"]
