const gaugeXElementID = "gaugeX",
  gaugeYElementID = "gaugeY",
  gaugePanelStatusElementID = "gauge-panel-status";

let gauges = {};

/**
 * GaugePanel is class encapsulating two gauges and status message
 *
 * @param {int} volumeX           gauge X capacity
 * @param {int} volumeY           gauge Y capacity
 */
class GaugePanel {

  constructor(volumeX, volumeY) {
    scaleGauges(volumeX, volumeY);
    gauges.gaugeX = createGauge(gaugeXElementID, volumeX, 0);
    gauges.gaugeY = createGauge(gaugeYElementID, volumeY, 0);
  }

  static get getGaugeXElementID() {
    return gaugeXElementID;
  }

  static get getGaugeYElementID() {
    return gaugeYElementID;
  }

  static get getGaugePanelStatusElementID() {
    return gaugePanelStatusElementID;
  }

  getGaugeX() {
    return gauges.gaugeX
  }

  getGaugeY() {
    return gauges.gaugeY
  }

  updateGaugesPanel(levelX, levelY, state, timeout) {
    setTimeout(function () {
      console.log(state);
      console.log(levelX);
      console.log(levelY);
      document.getElementById(gaugePanelStatusElementID).innerHTML = state;
      gauges.gaugeX.update(levelX);
      gauges.gaugeY.update(levelY);
    }, timeout);
  }

}

function scaleGauges(volumeX, volumeY) {
  const minAllowedScaleFactor = 0.15;
  let maxWidthPercentage = 30;
  let maxHeightValue = 250;
  if (volumeX > volumeY) {
    let factor = volumeY * 1.0 / volumeX;
    if (factor < minAllowedScaleFactor) {
      factor = minAllowedScaleFactor
    }
    let widthPercentageForY = maxWidthPercentage * factor;
    let heightValueForY = maxHeightValue * factor;
    document.getElementById(gaugeYElementID).setAttribute("width",
      widthPercentageForY.toString() + "%");
    document.getElementById(gaugeYElementID).setAttribute("height", heightValueForY.toString());
  }
  if (volumeX < volumeY) {
    let factor = volumeX * 1.0 / volumeY;
    if (factor < minAllowedScaleFactor) {
      factor = minAllowedScaleFactor
    }
    let widthPercentageForX = maxWidthPercentage * factor;
    let heightValueForX = maxHeightValue * factor;
    document.getElementById(gaugeXElementID).setAttribute("width",
      widthPercentageForX.toString() + "%");
    document.getElementById(gaugeXElementID).setAttribute("height", heightValueForX.toString());
  }
}

function createGauge(elementID, maxGaugeValue, currentGaugeMeasurement) {
  let config = liquidFillGaugeDefaultSettings();
  config.maxValue = maxGaugeValue;
  config.circleThickness = 0.1;
  config.circleColor = "#FFFFFF";
  config.textColor = "#A9A9A9";
  config.waveColor = "#7FA9D1";
  config.textVertPosition = 0.8;
  config.waveAnimateTime = 1000;
  config.waveAnimate = true;
  config.waveCount = 1;
  config.waveHeight = 0.05;
  config.displayPercent = false;
  return loadLiquidFillGauge(elementID, currentGaugeMeasurement, config);
}