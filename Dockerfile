FROM golang:1.15
WORKDIR .
COPY . .
RUN go build -o server .
CMD ["./server"]