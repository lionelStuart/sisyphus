FROM golang

MAINTAINER jim

WORKDIR /app

COPY app .

EXPOSE 8080

ENTRYPOINT ["./main"]
