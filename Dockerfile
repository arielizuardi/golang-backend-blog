FROM golang

RUN go get -u github.com/golang/dep/cmd/dep 

WORKDIR /go/src/github.com/arielizuardi/sph-backend-coding-challenge

COPY . .

RUN go build -o sph-backend-coding-challenge

EXPOSE 8080
CMD ./sph-backend-coding-challenge
