FROM golang:latest

RUN mkdir -p /agenda-cs
RUN apt-get update && apt-get install -y sqlite3
ADD . /agenda-cs/
WORKDIR /agenda-cs
RUN cd cli && go get -d -v && go build
RUN cd services/service &&  go get -d -v && go build

CMD ["/agenda-cs/services/service/service"]

