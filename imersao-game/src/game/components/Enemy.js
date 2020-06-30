class Enemy extends Animation {
  constructor(
    motionMatrix,
    image,
    x,
    variationY,
    dimension,
    spritDimension,
    collisionFactors,
    flying = false
  ) {
    super(motionMatrix, image, x, variationY, dimension, spritDimension);
    this.collisionFactors = collisionFactors;
    this.speed = 10;
    this.x = width;

    this.flying = flying;
    this.oldSenoide = 0;
    this.amplitude = 500;
  }

  move() {
    if (this.flying) {
      this.y = this.__getAxisYRandomFlying(this.y);
    }

    this.x = this.x - this.speed;
  }

  appear() {
    this.x = width;
  }

  isVisible() {
    return this.x < -this.dimension.width;
  }

  __getAxisYRandomFlying(y) {
    const senoide = parseInt(sin((2 * Math.PI * this.x) / this.amplitude) * 80);
    const returnedY = y + senoide - this.oldSenoide;

    if (this.x + this.dimension.width < 0) {
      this.amplitude = random(500, 5000);
    }

    this.oldSenoide = senoide;
    return returnedY;
  }
}
