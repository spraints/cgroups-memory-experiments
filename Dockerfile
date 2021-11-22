FROM golang:1.17.2

COPY . /src
WORKDIR /src
# Compile the C code.
# Remove the C code because Go refuses to build when there's C code present.
# Compile the Go code.
RUN set -x && cc -o c child.c && rm child.c && go build -o p .
