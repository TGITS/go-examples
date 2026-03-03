# AoC 2023 Jour 6

**Version anglaise disponible :** [README.md](README.md)

## Aperçu

Solution Go pour [Advent of Code 2023 - Jour 6](https://adventofcode.com/2023/day/6).

## Exécution

Depuis ce dossier (`aoc/2023/aoc20231206`) :

```powershell
go run .
```

## Tests

Depuis ce dossier (`aoc/2023/aoc20231206`) :

```powershell
go test ./...
```

## Benchmarks

Depuis ce dossier (`aoc/2023/aoc20231206`) :

```powershell
# Exécuter tous les benchmarks boatrace
go test ./boatrace -bench . -benchmem -run "^$"

# Exemples one-shot
go test ./boatrace -bench "BenchmarkNumberOfWaysToWin_SmallRace$" -benchmem -benchtime=1x -run "^$"
go test ./boatrace -bench "BenchmarkRecord(Breaking|BreakingsProducts)_InputTest$" -benchmem -benchtime=1x -run "^$"
```

Ou lancez les scripts helper :

```powershell
powershell -ExecutionPolicy Bypass -File ".\run_benchmarks.ps1"
```

```bash
bash ./run_benchmarks.sh
```

## Fichiers d'entrée

- `data/input_test.txt`
- `data/input.txt`
  - Ce fichier est spécifique à chaque utilisateur AoC et n'est pas versionné dans le dépôt : vous devez fournir votre propre fichier d'entrée.
