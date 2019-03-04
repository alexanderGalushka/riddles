function solveRiddle() {
  let volumeXStr = document.getElementById("x").value;
  let volumeYStr = document.getElementById("y").value;
  let volumeZStr = document.getElementById("z").value;
// TODO validate the inputs
  let volumeX = parseInt(volumeXStr).valueOf();
  let volumeY = parseInt(volumeYStr).valueOf();
  let volumeZ = parseInt(volumeZStr).valueOf();

// display the problem statement based on the inputs
  document.getElementById("guage-panel-status").innerHTML =
    "Measure " + volumeZStr + " gallons of water with " + volumeXStr + " gallon jug and " + volumeYStr + " gallon jug";

// initialize guages with values from input X and Y
  scaleGuagesBasedOnVolume(volumeX, volumeY);
  let guageX = displayGuage("guageX", volumeX, 0);
  let guageY = displayGuage("guageY", volumeY, 0);

  let solutionSteps = getWaterJugRiddleSolution(volumeX, volumeY, volumeZ);

  enableControls(false);

  for (let i = 0; i < solutionSteps.length; i++) {
    setTimeout(function () {
      let state = solutionSteps[i].state;
      console.log(state);
      document.getElementById("guage-panel-status").innerHTML = state;
      let x = solutionSteps[i].x;
      console.log(x);
      let y = solutionSteps[i].y;
      console.log(y);
      guageX.update(x);
      guageY.update(y);
    }, (i + 1) * 2500);
  }

  enableControls(true);

}

function getWaterJugRiddleSolution(x, y, z) {

  let queryParams = "x=" + x + "y=" + y + "z=" + z;
  let uri = "localhost:3000/riddles/water_jug?" + queryParams;
  console.log(uri);

  // axios.get(uri)
  //   .then(function (response) {
  //     console.log(response);
  //   })
  //   .catch(function (error) {
  //     console.log(error);
  //   });


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