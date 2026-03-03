# AoC 2023 Jour 10

**Version anglaise disponible :** [README.md](README.md)

## Aperçu

Solution Go pour [Advent of Code 2023 - Jour 10](https://adventofcode.com/2023/day/10) : _Pipe Maze_

## Exécution

Depuis ce dossier (`aoc/2023/aoc20231210`) :

```powershell
go run .
```

Aide CLI :

```powershell
go run . -h
```

Options disponibles :

- `-demo-images` : génère uniquement les PNG de prévisualisation de test
- `-pixel-size` : taille des pixels pour le rendu PNG (défaut : `2`)
- `-no-ascii` : désactive les visualisations ASCII (console et fichiers `.txt`)
- `-png-only` : alias de `-no-ascii`

Usages courants :

| Cas d'usage | Commande |
|---|---|
| Exécution standard (ASCII + PNG) | `go run .` |
| PNG uniquement (sans sortie ASCII) | `go run . -png-only` |
| PNG uniquement avec tuiles plus grandes | `go run . -png-only -pixel-size 4` |
| Générer seulement les PNG de prévisualisation de test | `go run . -demo-images` |
| Prévisualisations de test avec taille personnalisée | `go run . -demo-images -pixel-size 16` |
| Afficher l'aide CLI | `go run . -h` |

## Tests

Depuis ce dossier (`aoc/2023/aoc20231210`) :

```powershell
go test ./...
```

## Visualisation

Le solveur génère des visualisations ASCII et PNG.

### Sortie ASCII

Lors de `go run .`, le programme affiche et sauvegarde :

- `visualization_loop.txt`
  - Vue de la boucle uniquement
  - Les tuiles de la boucle utilisent les caractères de tuyaux
  - Les tuiles hors boucle sont notées `.`
- `visualization_distances.txt`
  - Vue des distances le long de la boucle
  - Les tuiles de la boucle utilisent des symboles hexadécimaux (`0-9`, `A-F`) selon la distance au départ
  - Les tuiles hors boucle sont notées `.`

Si `-no-ascii` ou `-png-only` est activé, cette sortie ASCII est ignorée.

### Sortie PNG

Lors de `go run .`, le programme sauvegarde aussi :

- `visualization_loop_part1.png`
  - Rendu partie 1 (boucle vs hors boucle)
- `visualization_enclosed_part2.png`
  - Rendu partie 2 (boucle + intérieur + extérieur)

### Code couleur

Palette utilisée pour les PNG :

- **Bleu** (`RGB 0,100,200`) : tuiles appartenant à la boucle principale
- **Vert** (`RGB 100,200,100`) : tuiles à l'intérieur de la boucle (vue partie 2)
- **Gris clair** (`RGB 200,200,200`) : tuiles à l'extérieur de la boucle

### Modèle de rendu

- Chaque tuile de la grille est rendue comme un carré plein (`pixelSize`)
- Valeur CLI par défaut : `-pixel-size 2`
- Une valeur plus grande donne une image plus zoomée

### Exemples

L'exemple ci-dessous utilise `./data/input_1_1.txt`.

- Vue boucle (`visualization_loop.txt`) :

```text
.....
.F-7.
.|.|.
.L-J.
.....
```

- Vue distances (`visualization_distances.txt`) :

```text
.....
.012.
.1.3.
.234.
.....
```

- Rappel légende PNG :
  - Bleu : boucle
  - Vert : intérieur (image partie 2)
  - Gris clair : extérieur

- Fichiers images d'exemple :
  - [Exemple boucle partie 1](./test_loop_part1.png)
  - [Exemple intérieur partie 2](./test_enclosed_part2.png)

Aperçu boucle partie 1 :

<a href="./test_loop_part1.png">
  <img src="./test_loop_part1.png" alt="Aperçu boucle partie 1" width="220" />
</a>

Aperçu intérieur partie 2 :

<a href="./test_enclosed_part2.png">
  <img src="./test_enclosed_part2.png" alt="Aperçu intérieur partie 2" width="220" />
</a>

## Explication de l'algorithme

### Aperçu du problème (Partie 1)

Vous avez un ensemble de tuyaux arrangés sur une grille 2D.

- `|` connecte nord et sud
- `-` connecte est et ouest
- `L` connecte nord et est
- `J` connecte nord et ouest
- `7` connecte sud et ouest
- `F` connecte sud et est
- `.` représente le sol (pas de tuyau)
- `S` est la position de départ

L'entrée peut contenir des tuyaux qui ne font pas partie de la boucle principale. On doit d'abord identifier cette boucle principale en partant de `S`, puis calculer la plus grande distance minimale à `S` le long de la boucle.

Données de test (partie 1) :

- `./data/input_1_1.txt` → distance max attendue : 4
- `./data/input_1_2.txt` → distance max attendue : 4
- `./data/input_1_3.txt` → distance max attendue : 8

L'entrée puzzle est `./data/input.txt`.

### Algorithme de résolution (Partie 1)

1. Lire la grille et repérer `S`
2. Déduire l'orientation implicite de `S` via ses voisins
3. Parcourir la boucle en suivant les connexions de tuyaux
4. Calculer les distances minimales depuis `S` dans les deux sens de la boucle
5. Retourner la distance maximale obtenue

### Aperçu du problème (Partie 2)

Pour la partie 2, on conserve la même boucle principale, mais l'objectif devient le comptage des tuiles **enfermées** à l'intérieur de la boucle.

Données de test (partie 2) :

- `./data/input_2_1.txt` → 4 tuiles intérieures
- `./data/input_2_2.txt` → 4 tuiles intérieures
- `./data/input_2_3.txt` → 8 tuiles intérieures
- `./data/input_2_4.txt` → 10 tuiles intérieures

### Algorithme de résolution (Partie 2)

La résolution utilise un **ray-casting** (règle pair/impair) :

1. Identifier la boucle principale (comme en partie 1)
2. Pour chaque tuile hors boucle, lancer un rayon horizontal vers la droite
3. Compter les intersections avec la frontière de la boucle
4. Nombre impair d'intersections = tuile intérieure, sinon extérieure
5. Additionner les tuiles intérieures

Points d'attention :

- Les coins (`L`, `J`, `7`, `F`) doivent être traités soigneusement pour éviter le double comptage
- Les tuiles parasites non connectées à la boucle principale sont ignorées pour la frontière
