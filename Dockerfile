FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Directory which will be mounted
RUN mkdir /root/test.dir
RUN mkdir /mnt
COPY test-key /root

COPY efsctl /root/
WORKDIR /root/

# Initialize EFS
CMD ["./efsctl", "init", "-s", "test.dir", "-t", "/mnt", "-k", "test-key"]
