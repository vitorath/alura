class Menu {
  constructor(states) {
    this.states = states;
    console.log(this.states);
  }

  setup() {
    this.startButton = new Button(
      "Iniciar",
      width / 2 - 87,
      (height / 7) * 5,
      () => this.__changeToGameScene()
    );
  }

  setAssets(backgrounds, fonts) {
    this.backgrounds = backgrounds;
    this.fonts = fonts;
  }

  draw() {
    this.__setupBackground();
    this.__setupText();
    this.__setupStartButton();
  }

  __setupBackground() {
    image(this.backgrounds.menu, 0, 0, width, height);
  }

  __setupText() {
    textFont(this.fonts.menu);
    textAlign(CENTER);
    textSize(50);
    text("As aventuras de", width / 2, height / 3);
    textSize(150);
    text("Hipsta", width / 2, (height / 5) * 2.5);
  }

  __setupStartButton() {
    this.startButton.draw();
  }

  __changeToGameScene() {
    this.startButton.button.remove();
    this.states.scene = "game";
  }
}
