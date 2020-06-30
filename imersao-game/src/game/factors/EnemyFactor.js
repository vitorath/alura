class EnemyFactory {
  static buildDroplets(sprit) {
    const motionMatrix = generateMotionMatrix(105, 100, 4, 28);
    const dimension = new Dimension(52, 50);
    const spritDimensions = new Dimension(105, 100);
    const collisionFactor = [
      new CollisionFactor(
        abs((52 * 0.7) / 2 - 52 / 2) + 10,
        abs((50 * 0.7) / 2 - 52 / 2),
        52 * 0.7 + 10,
        50 * 0.7
      ),
    ];

    return new Enemy(
      motionMatrix,
      sprit,
      width,
      35,
      dimension,
      spritDimensions,
      collisionFactor
    );
  }

  static buildFlyingDroplets(sprit) {
    const motionMatrix = generateMotionMatrix(200, 150, 3, 16);
    const dimension = new Dimension(102, 76);
    const spritDimensions = new Dimension(200, 150);
    const collisionFactor = [
      new CollisionFactor(
        abs((102 * 0.5) / 2 - 102 / 2) + 10,
        abs((76 * 0.5) / 2 - 76 / 2),
        102 * 0.5 + 10,
        76 * 0.5
      ),
    ];

    return new Enemy(
      motionMatrix,
      sprit,
      width - 52,
      200,
      dimension,
      spritDimensions,
      collisionFactor,
      true
    );
  }

  static buildTroll(sprit) {
    const motionMatrix = generateMotionMatrix(400, 400, 5, 28);
    const dimension = new Dimension(400, 400);
    const spritDimensions = new Dimension(400, 400);
    const collisionFactor = [
      new CollisionFactor(200, 80, 120, 220),
      new CollisionFactor(50, 200, 180, 100),
    ];
    return new Enemy(
      motionMatrix,
      sprit,
      width,
      -25,
      dimension,
      spritDimensions,
      collisionFactor
    );
  }

  static buildEnemies(sprits) {
    const enemies = [];

    const droplets = this.buildDroplets(sprits.droplets);
    const troll = this.buildTroll(sprits.troll);
    const flyingDroplets = this.buildFlyingDroplets(sprits.flyingDroplets);

    enemies.push(droplets);
    enemies.push(troll);
    enemies.push(flyingDroplets);

    return enemies;
  }
}
