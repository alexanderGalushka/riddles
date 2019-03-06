RIDDLES web application
==========================

RIDDLES is the application for curious minds. Enjoy detailed step by step solutions for mathematical riddles.
For some of the riddles you have control over the input parameters.

E.g. for infamous water jug riddle you can enter X, Y, Z values, where X and Y are the capacities of the 
jugs and Z is the goal water measurement.


### SETUP DEV

- install golang, the service if currently running on 1.10 version
- configure $GOPATH environment variable
- install docker
- install npm
- make sure make is installed

### How to run

- to run this application do `make run`

- it will bring up 2 web servers: node.js to display ui and golang for riddle API
- go to your browser and type in http://localhost:8080 to see the application running


### TEST

- run unit tests `make unit`
- run static analysis golang tools `make lint`