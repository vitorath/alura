class Button {
  constructor(text, x, y, handleMousePressed) {
    this.text = text;
    this.x = x;
    this.y = y;
    this.button = createButton(this.text);
    this.button.addClass("manager-button");
    this.button.mousePressed(() => handleMousePressed());
  }

  draw() {
    this.button.position(this.x, this.y);
    // this.button.center("horizontal");
  }
}
