RIDDLES web application
==========================

RIDDLES is the application for curious minds. Enjoy detailed step by step solutions for mathematical riddles.
For some of the riddles you have control over the input parameters.

E.g. for infamous water jug riddle you can enter X, Y, Z values, where X and Y are the capacities of the 
jugs and Z is the goal water measurement.


### SETUP DEV

- install latest golang version
- configure $GOPATH environment variable
- install docker
- install npm
- make should be installed by default

### How to run

- client and api are run as 2 separate applications:
  - open terminal, run client: `make jsclient`
  - open terminal, run api: `make goapi`
- go to http://localhost:8080 to test out the application, api is running on http://localhost:3000


### TEST

- run unit tests `make unit`