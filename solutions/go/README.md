# Advent of Code

My [Go](https://golang.org/) solutions to Advent of Code puzzles.

## Progress

| **Year** | **Stars** |
|:--------:|:---------:|
|   2015   |     0     |
|   2016   |     0     |
|   2017   |     0     |
|   2018   |     0     |
|   2019   |     0     |
|   2020   |     0     |
|   2021   |     0     |
|   2022   |     0     |
|   2023   |     0     |
|   2024   |     50    |

## Usage

### Solving puzzles

Solves depend on the personalized input files generated for each user. These files must be stored within the `infiles` directory at the root of the repository, named `{year}/{day}.txt`, e.g., `2024/09.txt`.

Solve every puzzle with a solution:

```bash
go run cmd/solve/main.go
```

Solve a specific day's puzzle:

```bash
go run cmd/solve/main.go -day 9
```

### Running tests

Tests depend on the example solutions part of each puzzle's description. These files must be stored within the `infiles` directory at the root of the repository, named `{year}/{day}_test.txt`, e.g., `2024/09_test.txt`.

Test every puzzle with solutions:

```bash
go test pkg/puzzles
```

Test a specific day's puzzle:

```bash
go test pkg/puzzles -run Nine
```

### Running benchmarks

Benchmarks depend on the personalized input files generated for each user. These files must be stored within the `infiles` directory at the root of the repository, named `{year}/{day}.txt`, e.g., `2024/09.txt`.

Benchmark every puzzle with a solution:

```bash
go test pkg/puzzles -bench .
```

Benchmark a specific day's puzzle:

```bash
go test pkg/puzzles bench Nine
```
