function solveRiddle() {
  enableControls(false);
  showJugLabels();
  disableLoader();
  solveWaterJugRiddle();
  enableControls(true);
}

function solveWaterJugRiddle() {

  let volumeX, volumeY, volumeZ;
  [volumeX, volumeY, volumeZ] = getRiddleInputsAndDisplayProblemStatement();

  const gaugePanel = new GaugePanel(volumeX, volumeY);

  let queryParams = "x=" + volumeX + "&y=" + volumeY + "&z=" + volumeZ;
  let url = "http://localhost:3000/v1/riddles/water_jug/solution?" + queryParams;
  console.log(url);

  fetch(url).then(response => {
    return response.json();
  }).then(solutionSteps => {
    console.log(solutionSteps);
    let steps = solutionSteps.steps;

    for (let i = 0; i < steps.length; i++) {
      let timeout = (i + 1) * 2500;
      gaugePanel.updateGaugesPanel(steps[i].x, steps[i].y, steps[i].state, timeout);
    }


  }).catch(err => {
    console.log('failed to fetch solution from riddle API', err)
  });

}

function getRiddleInputsAndDisplayProblemStatement() {
  let volumeXStr = document.getElementById("x").value;
  let volumeYStr = document.getElementById("y").value;
  let volumeZStr = document.getElementById("z").value;

  // display the problem statement based on the inputs
  document.getElementById(GaugePanel.getGaugePanelStatusElementID).innerHTML =
    "Measure " + volumeZStr + " gallons of water with " + volumeXStr + " gallon jug and " + volumeYStr + " gallon jug";

  // TODO validate inputs
  let volumeX = parseInt(volumeXStr).valueOf();
  let volumeY = parseInt(volumeYStr).valueOf();
  let volumeZ = parseInt(volumeZStr).valueOf();
  return [volumeX, volumeY, volumeZ]
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

function disableLoader() {
  document.getElementById("loader").outerHTML = "";
}

function showJugLabels() {
  document.getElementById("xLabel").innerHTML = "Jug X gauge";
  document.getElementById("yLabel").innerHTML = "Jug Y gauge";
}

function reload() {
  location.reload();
}
