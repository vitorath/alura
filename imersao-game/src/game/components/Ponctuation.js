class Ponctuation {
  constructor() {
    this.points = 0;
  }

  show() {
    textAlign(RIGHT);
    fill("#FFF");
    textSize(37);
    text(parseInt(this.points), width - 30, 70);
  }

  add() {
    this.points += 0.1;
  }
}
