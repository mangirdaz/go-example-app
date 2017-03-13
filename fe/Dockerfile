FROM docker.io/library/golang
MAINTAINER mangirdas@judeikis.lt
RUN mkdir /deployment
COPY       fe /deployment/fe
COPY       image.tmpl /deployment/image.tmpl
COPY       index.tmpl /deployment/index.tmpl
ENTRYPOINT ["/deployment/fe"]
# Document that the service listens on port 8001.
EXPOSE 8080
