class Game {
  constructor() {
    this.mapIndex = 0;
    this.map = [
      {
        enemy: 0,
        speed: 15,
      },
      {
        enemy: 1,
        speed: 10,
      },
      {
        enemy: 1,
        speed: 20,
      },
      {
        enemy: 2,
        speed: 40,
      },
    ];
  }

  setAssets(sprits, sounds, backgrounds) {
    this.sprits = sprits;
    this.sounds = sounds;
    this.backgrounds = backgrounds;
  }

  setup() {
    this.scenario = new Scenario(this.backgrounds.game, 3);
    this.character = Character.initWitch(this.sprits.witch);
    this.enemies = EnemyFactory.buildEnemies(this.sprits);
    this.ponctuation = new Ponctuation();
    this.lifes = new Life(this.sprits.heart, 3, 3);
  }

  keyPress(key) {
    if (key === "ArrowUp") {
      this.character.jump(this.sounds.jump);
    }
  }

  draw() {
    this.scenario.show();
    this.scenario.move();

    if (this.__isGameOver()) {
      image(this.sprits.gameOver, width / 2 - 200, height / 3);
      noloop();
    }

    this.lifes.draw();

    const currentMapLine = this.map[this.mapIndex];
    this.__showEnemies(currentMapLine);
    this.__colliderWithMainCharacter();

    this.character.show();
    this.character.applyGravity();

    this.ponctuation.show();
    this.ponctuation.add();
  }

  __showEnemies(currentMapLine) {
    this.visibleEnemy = this.enemies[currentMapLine.enemy];
    this.visibleEnemy.speed = currentMapLine.speed;

    this.visibleEnemy.show();
    this.visibleEnemy.move();

    this.__changeEnemies();
  }

  __changeEnemies() {
    if (this.visibleEnemy.isVisible()) {
      this.mapIndex++;
      this.visibleEnemy.appear();
      if (this.mapIndex > this.map.length - 1) {
        this.mapIndex = 0;
      }
    }
  }

  __colliderWithMainCharacter() {
    if (this.character.isColliding(this.visibleEnemy)) {
      this.lifes.removeLife();
      this.character.stayInvencible();
    }
  }

  __isGameOver() {
    return this.lifes.currentLifes === 0;
  }
}
