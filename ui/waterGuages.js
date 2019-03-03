function displayGuage(elementID, maxGuageValue, currentGuageMeasure) {
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
function scaleGuagesBasedOnVolume(volumeX, volumeY) {
  const minAllowedScaleFactor = 0.15;
  var maxWidthPercentage = 30;
  var maxHeightValue = 250;
  if (volumeX > volumeY) {
    var factor = volumeY * 1.0 / volumeX;
    if (factor < minAllowedScaleFactor) {
      factor = minAllowedScaleFactor
    }
    var widthPercentageForY = maxWidthPercentage * factor;
    var heightValueForY = maxHeightValue * factor;
    document.getElementById("guageY").setAttribute("width",
      widthPercentageForY.toString() + "%");
    document.getElementById("guageY").setAttribute("height", heightValueForY.toString());
  }
  if (volumeX < volumeY) {
    var factor = volumeX * 1.0 / volumeY;
    if (factor < minAllowedScaleFactor) {
      factor = minAllowedScaleFactor
    }
    var widthPercentageForX = maxWidthPercentage * factor;
    var heightValueForX = maxHeightValue * factor;
    document.getElementById("guageX").setAttribute("width",
      widthPercentageForX.toString() + "%");
    document.getElementById("guageX").setAttribute("height", heightValueForX.toString());
  }

}