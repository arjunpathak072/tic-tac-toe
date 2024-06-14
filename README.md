# Tic Tac Toe Game in Go and SDL

A simple implementation of the classic Tic Tac Toe game using the Go programming language and the SDL library for graphics and input handling.

## Getting Started

### Prerequisites

* Go 1.22 or later installed on your system
* SDL 2.0 or later installed on your system
* A compatible graphics driver (e.g. OpenGL, DirectX, etc.)

### Building and Running

1. Clone this repository: `git clone https://github.com/arjunpathak072/tic-tac-toe.git`
2. Change into the project directory: `cd tic-tac-toe`
3. Build the project: `go build .`
4. Run the game: `./tic-tac-toe`

## Gameplay

* The game is played on a 3x3 grid, with two players: X and O.
* Players take turns clicking on empty squares to place their mark.
* The game checks for wins after each move, and the game ends if a player has won.

## Features

* Simple and intuitive gameplay
* Graphics and input handling using SDL
* Written in Go for performance and concurrency

## Acknowledgments

* The SDL library for providing a cross-platform graphics and input handling API.
* The Go community for providing a great language and ecosystem.

## TDB

- [ ] Allow players to restart the game through graphical input
- [ ] Allow users pick the shape they want to play as
- [ ] Add some visual feedback when a player wins the game
- [ ] Add sound effects