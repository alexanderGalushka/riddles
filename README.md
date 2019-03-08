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

Please bear in mind that by design client and API are 2 separate web servers (nodjs and golang respectively).

1. open terminal in the riddles project directory, type in the following command to run client: `make jsclient`
  
2. open terminal in the riddles project directory, type in the following command to run run api: `make goapi`
  
3. go to Chrome or Safari, enter http://localhost:8080 to n, api is running on http://localhost:3000


### TEST

- run unit tests `make unit`

### NOTES

- D3 is being used for water gauges visualization
- 100% CSS small loader animation is being added for dramatic affect
- for better perception of the gauges difference, scaling is applied on each svg object, e.g. if X=5, Y=3, X would look
  bigger than Y
- inputs are validated and user would be alerted if the data entry is not adequate
- every step of the riddle is animated with proper explanation
- in UI at any point of time you can start over if you click 'START OVER' button (useful when you need to try out another
  set of inputs or if you have entered large numbers as your inputs and you are borred looking at gauges :) )  
- API returns the array of json objects, the steps to solve the riddle, each step has a state descriptor, current value
  for X and Y  
- water jug riddle edge cases considered on the API side
- the 'water jug steps compute algorithm' is being optimized by using the Bezout's identity which comes down to finding
  greatest common divisor
- 2 jugs have been unified under so-called container and data and functions for the solution are structured around
  container abstraction   
- API is being unit tested
- for technical debt please check the github issues