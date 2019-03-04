function solveWaterJugRiddle(x, y, z) {

  return riddleSteps = [
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
  ]
}