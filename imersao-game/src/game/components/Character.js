class Character extends Animation {
  static initWitch(sprit) {
    const motionMatrix = generateMotionMatrix(220, 270, 4, 16);
    const dimension = new Dimension(164, 184);
    const spritDimension = new Dimension(220, 270);
    const collisionFactor = new CollisionFactor(
      abs((164 * 0.55) / 2 - 164 / 2),
      abs((184 * 0.8) / 2 - 184 / 2),
      164 * 0.5,
      184 * 0.8
    );

    return new Character(
      motionMatrix,
      sprit,
      parseInt(width / 5),
      35,
      dimension,
      spritDimension,
      collisionFactor
    );
  }

  constructor(
    motionMatrix,
    image,
    x,
    variationY,
    dimension,
    spritDimension,
    collisionFactor
  ) {
    super(motionMatrix, image, x, variationY, dimension, spritDimension);
    this.collisionFactor = collisionFactor;

    this.initialY = height - this.dimension.height - variationY;
    this.gravity = 3;
    this.jumpSpeed = 0;

    this.y = this.initialY;
    this.jumps = 0;
    this.invencible = false;
  }

  jump(sound) {
    if (this.jumps < 2) {
      sound.play();
      this.jumpSpeed = -40;
      this.jumps++;
    }
  }

  stayInvencible() {
    this.invencible = true;
    setTimeout(() => {
      this.invencible = false;
    }, 1000);
  }

  isColliding(enemy) {
    if (this.invencible) {
      return false;
    }

    return enemy.collisionFactors.some((factor) => {
      return collideRectRect(
        this.x + this.collisionFactor.x,
        this.y + this.collisionFactor.y,
        this.collisionFactor.width,
        this.collisionFactor.height,
        enemy.x + factor.x,
        enemy.y + factor.y,
        factor.width,
        factor.height
      );
    });
  }

  applyGravity() {
    this.y += this.jumpSpeed;
    this.jumpSpeed += this.gravity;
    if (this.y > this.initialY) {
      this.y = this.initialY;
      this.jumps = 0;
    }
  }
}
