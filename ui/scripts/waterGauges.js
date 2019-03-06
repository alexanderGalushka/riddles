function displayGauge(elementID, maxGuageValue, currentGuageMeasure) {
  var config = liquidFillGaugeDefaultSettings();
  config.maxValue = maxGuageValue;
  config.circleThickness = 0.1;
  config.circleColor = "#676468";
  config.textColor = "#A9A9A9";
  config.waveColor = "#3F78AA";
  config.textVertPosition = 0.8;
  config.waveAnimateTime = 1000;
  config.waveAnimate = true;
  config.waveCount = 1;
  config.waveHeight = 0.05;
  config.displayPercent = false;
  return loadLiquidFillGauge(elementID, currentGuageMeasure, config);
}

// TODO abstract common code out
function scaleGaugesBasedOnVolume(volumeX, volumeY) {
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
    document.getElementById("gaugeY").setAttribute("width",
      widthPercentageForY.toString() + "%");
    document.getElementById("gaugeY").setAttribute("height", heightValueForY.toString());
  }
  if (volumeX < volumeY) {
    let factor = volumeX * 1.0 / volumeY;
    if (factor < minAllowedScaleFactor) {
      factor = minAllowedScaleFactor
    }
    let widthPercentageForX = maxWidthPercentage * factor;
    let heightValueForX = maxHeightValue * factor;
    document.getElementById("gaugeX").setAttribute("width",
      widthPercentageForX.toString() + "%");
    document.getElementById("gaugeX").setAttribute("height", heightValueForX.toString());
  }

}