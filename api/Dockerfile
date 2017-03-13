FROM docker.io/library/golang
MAINTAINER mangirdas@judeikis.lt
COPY       api /bin/api
ENTRYPOINT ["/bin/api"]
ENV TIME_BOMB=true
# Document that the service listens on port 8080.
EXPOSE 8080
