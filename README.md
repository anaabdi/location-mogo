# LOCATION-MOGO

This is a simple project, using golang and mongodb.
It has a feature. Checking the position given is located in certain area (defined in mongodb).

# Prerequisites

golang: https://golang.org/doc/install
mongodb: https://docs.mongodb.com/manual/installation/

# Installing

go install

location-mogo -h
Usage of location-mogo:
  -mongoURL string
        connection url for mongodb (default "localhost:27017")
  -port string
        openend port for serving request (default "3222")

# Running the tests

go test -v