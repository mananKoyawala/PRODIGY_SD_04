# Sudoku Solver

## Description

This is a program that solves Sudoku puzzles automatically. The program takes an input grid representing an unsolved Sudoku puzzle and uses the backtracking algorithm to fill in the missing numbers. Once solved, the program displays the completed Sudoku grid.

## Features

- Solves any valid 9x9 Sudoku puzzle.
- Uses backtracking algorithm to explore possible solutions.
- Displays the completed Sudoku grid once solved.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/mananKoyawala/PRODIGY_SD_04.git
   cd PRODIGY_SD_04
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

## Usage

1. Run the application.
2. The program will print the unsolved Sudoku grid.
3. The program will then solve the puzzle and print the solved Sudoku grid.

## Example

```bash
Unsolved Sudoku

-------------------------
| 5 3 0 | 0 7 0 | 0 0 0 |
| 6 0 0 | 1 9 5 | 0 0 0 |
| 0 9 8 | 0 0 0 | 0 6 0 |
-------------------------
| 8 0 0 | 0 6 0 | 0 0 3 |
| 4 0 0 | 8 0 3 | 0 0 1 |
| 7 0 0 | 0 2 0 | 0 0 6 |
-------------------------
| 0 6 0 | 0 0 0 | 2 8 0 |
| 0 0 0 | 4 1 9 | 0 0 5 |
| 0 0 0 | 0 8 0 | 0 7 9 |
-------------------------

Solved Sudoku

-------------------------
| 5 3 4 | 6 7 8 | 9 1 2 |
| 6 7 2 | 1 9 5 | 3 4 8 |
| 1 9 8 | 3 4 2 | 5 6 7 |
-------------------------
| 8 5 9 | 7 6 1 | 4 2 3 |
| 4 2 6 | 8 5 3 | 7 9 1 |
| 7 1 3 | 9 2 4 | 8 5 6 |
-------------------------
| 9 6 1 | 5 3 7 | 2 8 4 |
| 2 8 7 | 4 1 9 | 6 3 5 |
| 3 4 5 | 2 8 6 | 1 7 9 |
-------------------------

```

## Contributing

Feel free to fork this project and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT License](LICENSE)
