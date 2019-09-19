from golang:1.13.0

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go install

EXPOSE 8000
CMD ["iknow"]
