FROM golang:1.22-bullseye

# add grpc compiler
RUN apt-get update && apt-get install -y protobuf-compiler golang-goprotobuf-dev

# add user
RUN useradd -ms /bin/bash vanio

# set user
RUN umask 000

# Create app directory
WORKDIR /app

# loop infinitely to keep the container running
CMD ["tail", "-f", "/dev/null"]
