function displayGuage(elementID, maxGuageValue, currentGuageMeasure) {
  var config = liquidFillGaugeDefaultSettings();
  config.maxValue = maxGuageValue;
  config.circleThickness = 0.1;
  config.circleColor = "#7b5080";
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