function solveRiddle() {

  let volumeX, volumeY, volumeZ;
  [volumeX, volumeY, volumeZ] = getRiddleInputsAndDisplayProblemStatement();

// initialize gauges with values from input X and Y
  scaleGaugesBasedOnVolume(volumeX, volumeY);
  let gaugeX = displayGauge("gaugeX", volumeX, 0);
  let gaugeY = displayGauge("gaugeY", volumeY, 0);

  let solutionSteps = getWaterJugRiddleSolution(volumeX, volumeY, volumeZ);

  enableControls(false);

  for (let i = 0; i < solutionSteps.length; i++) {
    setTimeout(function () {
      let state = solutionSteps[i].state;
      console.log(state);
      document.getElementById("gauge-panel-status").innerHTML = state;
      let x = solutionSteps[i].x;
      console.log(x);
      let y = solutionSteps[i].y;
      console.log(y);
      gaugeX.update(x);
      gaugeY.update(y);
    }, (i + 1) * 2500);
  }

  enableControls(true);

}

function getRiddleInputsAndDisplayProblemStatement() {
  let volumeXStr = document.getElementById("x").value;
  let volumeYStr = document.getElementById("y").value;
  let volumeZStr = document.getElementById("z").value;

  // display the problem statement based on the inputs
  document.getElementById("gauge-panel-status").innerHTML =
    "Measure " + volumeZStr + " gallons of water with " + volumeXStr + " gallon jug and " + volumeYStr + " gallon jug";

  // TODO validate the inputs
  let volumeX = parseInt(volumeXStr).valueOf();
  let volumeY = parseInt(volumeYStr).valueOf();
  let volumeZ = parseInt(volumeZStr).valueOf();
  return [volumeX, volumeY, volumeZ]
}

function getWaterJugRiddleSolution(x, y, z) {

  let queryParams = "x=" + x + "&y=" + y + "&z=" + z;
  let url = "http://localhost:3000/v1/riddles/water_jug/solution?" + queryParams;
  console.log(url);

  // axios.get(uri)
  //   .then(function (response) {
  //     console.log(response);
  //   })
  //   .catch(function (error) {
  //     console.log(error);
  //   });

  let requestConfig = {
    method: 'GET',
    mode: 'no-cors',
    headers: new Headers(
      {
        "Content-Type": "application/json",
        "Accept": "application/json"
      }
    )
  };

  fetch(url, requestConfig)
    .catch(error => console.log('failed to fetch solution from riddle API', error))
    .then(response => console.log(response.json()));


  let riddleSteps = [
    {
      "state": "Fill up X",
      "x": 5,
      "y": 0,
    },
    {
      "state": "Transfer from X to fill up Y",
      "x": 2,
      "y": 3,
    },
    {
      "state": "Empty Y",
      "x": 2,
      "y": 0,
    },
    {
      "state": "Transfer from X to Y",
      "x": 0,
      "y": 2,
    },
    {
      "state": "Fill up X",
      "x": 5,
      "y": 2,
    },
    {
      "state": "Transfer X to Y",
      "x": 4,
      "y": 3,
    },
  ];

  return riddleSteps;
}



function enableControls(controlsFlag) {
  let solveButton = document.getElementById("solveButton");
  solveButton.disabled = controlsFlag;
  let inputX = document.getElementById("x");
  inputX.disabled = controlsFlag;
  let inputY = document.getElementById("y");
  inputY.disabled = controlsFlag;
  let inputZ = document.getElementById("z");
  inputZ.disabled = controlsFlag;
}