class Animation {
  constructor(
    motionMatrix,
    image,
    x,
    variationY = 0,
    dimension = new Dimension(),
    spritDimension = new Dimension()
  ) {
    this.image = image;
    this.motionMatrix = motionMatrix;
    this.x = x;
    this.y = height - dimension.height - variationY;
    this.dimension = dimension;
    this.spritDimension = spritDimension;
    this.motionFrame = 0;
  }

  show() {
    //rect(this.x, this.y, this.dimension.width, this.dimension.height);
    image(
      this.image,
      this.x,
      this.y,
      this.dimension.width,
      this.dimension.height,
      this.motionMatrix[this.motionFrame].x,
      this.motionMatrix[this.motionFrame].y,
      this.spritDimension.width,
      this.spritDimension.height
    );
    this.animate();
  }

  animate() {
    this.motionFrame++;

    if (this.motionFrame > this.motionMatrix.length - 1) {
      this.motionFrame = 0;
    }
  }
}
