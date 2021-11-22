FROM golang:1.17.2

COPY . /src
WORKDIR /src
RUN cc -o c child.c && rm child.c
RUN go build -o p .
