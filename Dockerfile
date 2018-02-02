

# ----------- minimize version ----------#
# FROM iron/base
# WORKDIR /app
# COPY main /app
# ENTRYPOINT ["./main"]
# ----------- minimize version ----------#


# ----------- golang standard version ----------#
FROM golang:latest
WORKDIR /app
ENV SRC_DIR=/go/src/column_grouping
ADD . $SRC_DIR
COPY . /app
RUN cd $SRC_DIR ; go get ./... ; go build -o main; cp main /app/
ENTRYPOINT ["./main"]
# ----------- golang standard version ----------#

