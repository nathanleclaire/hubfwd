from golang:1.4.2

# TODO: Vendor if these get too hairy.
run go get -u github.com/gorilla/mux
run go get -u github.com/Sirupsen/logrus

run mkdir -p /go/src/github.com/nathanleclaire/hubfwd
workdir /go/src/github.com/nathanleclaire/hubfwd
copy . /go/src/github.com/nathanleclaire/hubfwd

cmd ["./build.sh"]
