class Life {
  constructor(sprit, total, start) {
    this.total = total;
    this.start = start;
    this.sprit = sprit;

    this.currentLifes = this.start;
    this.width = 25;
    this.height = 25;
    this.startPositionX = 20;
  }

  draw() {
    for (let i = 0; i < this.currentLifes; i++) {
      const margin = i * 10;
      const position = this.startPositionX * (i + 1) + margin;
      image(this.sprit, position, 20, this.width, this.height);
    }
  }

  addLife() {
    if (this.currentLifes <= this.total) {
      this.currentLifes++;
    }
  }

  removeLife() {
    this.currentLifes--;
  }
}
