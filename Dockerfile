# # FROM golang:latest
# # WORKDIR /app
# # COPY main /app
# # ENTRYPOINT ["./main"]

# FROM golang:latest 
# WORKDIR /app 
# ADD . /app
# RUN go get -d -v; go get ./... ; go build -o main . ; cp main /app/
# CMD ["./main"]

# # FROM golang:latest
# # WORKDIR /app
# # ADD . /app
# # RUN go get ./... ; go build -o column_grouping; cp column_grouping /app/

# # ENTRYPOINT ["./column_grouping"]

FROM golang:latest

# RUN mkdir -p /app

# ENV SRC_DIR=/go/src/column_grouping
# ADD . $SRC_DIR
# ADD ./config.json /app/config.json

# RUN go get ./... ; go build -o main; cp main /app/

WORKDIR /app
COPY . /app

ENTRYPOINT ["./main"]

