function generateMotionMatrix(dWidth, dHeight, maxWidthSprits, totalSprits) {
  let widthIndicator = 0;
  let heightIndicator = 0;

  const motionMatrix = [];

  while (totalSprits > 0) {
    const position = {
      x: widthIndicator * dWidth,
      y: heightIndicator * dHeight,
    };
    motionMatrix.push(position);

    if (widthIndicator === maxWidthSprits - 1) {
      widthIndicator = 0;
      heightIndicator += 1;
    } else {
      widthIndicator += 1;
    }

    totalSprits -= 1;
  }

  return motionMatrix;
}
