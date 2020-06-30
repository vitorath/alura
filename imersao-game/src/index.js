const states = new State();
const game = new Game();
console.log(states);
const menu = new Menu(states);

const scenas = { game, menu };

function keyPressed() {
  game.keyPress(key);
}

function preload() {
  const sounds = new Sounds();
  const sprits = new Sprits();
  const backgrounds = new Backgrounds();
  const fonts = new Fonts();
  game.setAssets(sprits, sounds, backgrounds);
  menu.setAssets(backgrounds, fonts);
}

function setup() {
  //console.log(windowWidth - 400, windowHeight - 400);
  createCanvas(1280, 720);
  game.setup();
  menu.setup();
  frameRate(40);
  // sounds.track.loop();
}

function draw() {
  scenas[states.scene].draw();
}
