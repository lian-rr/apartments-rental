FROM golang:1.8
WORKDIR /go/src/app
COPY . .
EXPOSE 80
RUN go get -d -v ./...
RUN go install -v ./...
RUN make
CMD ["/go/src/app/apartments"]