FROM openshift/jenkins-2-centos7

COPY plugins.txt /opt/openshift/configuration/plugins.txt

USER root
RUN cd /usr/local && \
    curl https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz | tar xz && \
    mkdir /go

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH


RUN /usr/local/bin/install-plugins.sh /opt/openshift/configuration/plugins.txt