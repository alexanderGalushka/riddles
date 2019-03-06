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

#### Problem Statement

OVERVIEW

Build an application that solves the Water Jug Riddle for dynamic inputs (X, Y, Z). The simulation should have a UI to display state changes for each state for each jug (Empty, Full or Partially Full).
You have an X-gallon and a Y-gallon jug that you can fill from a lake. (Assume lake has unlimited amount of water.) By using only an X-gallon and Y-gallon jug (no third jug), measure Z gallons of water.
GOALS

1. Measure Z gallons of water ​in the most efficient ​way.
2. Build a UI where a user can enter any input for X, Y, Z and see the solution.
3. If no solution, display “No Solution”.

LIMITATIONS

● No partial measurement. Each jug can be empty or full.
● Actions allowed: Fill, Empty, Transfer.
● Use one of the following programming languages: Scala, Java, Nodejs, Go, Python, C,
C++, Kotlin.

DELIVERABLES

The application source code should be on Github and a link should be provided. If this is not an option, a public link to the application source code or a zip archive is also acceptable.

EVALUATION CRITERIAS

● Functionality
● Efficiency (Time, Space)
● Code Quality / Design / Patterns
● Testability
● UI/UX design