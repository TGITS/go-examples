# AoC 2023 Jour 5

**Version anglaise disponible :** [README.md](README.md)

## Aperçu

Solution en Go pour [Advent of Code 2023 - Jour 5](https://adventofcode.com/2023/day/5)

## Exécution

Depuis ce dossier (`aoc/2023/aoc20231205`) :

```powershell
go run .
```

## Tests

Depuis ce dossier (`aoc/2023/aoc20231205`) :

```powershell
go test ./...
```

## Benchmarks

Ce module inclut des benchmarks dédiés pour comparer :
- `BruteForce` (approche originale de la partie 2 graine par graine)
- `Optimized` (approche pipeline d'intervalles)

Depuis ce dossier (`aoc/2023/aoc20231205`), lancez :

```powershell
# Optimisé sur l'entrée complète (exécution unique)
go test ./almanac -bench "BenchmarkPart2Optimized_Input$" -benchmem -benchtime=1x -run "^$"

# Force brute sur l'entrée complète (exécution unique, peut être très long)
go test ./almanac -bench "BenchmarkPart2BruteForce_Input$" -benchmem -benchtime=1x -run "^$"

# Comparaison rapide sur l'entrée de test
go test ./almanac -bench "BenchmarkPart2(Optimized|BruteForce)_InputTest$" -benchmem -benchtime=1x -run "^$"
```

Remarques :
- Les benchmarks sont "propres" : les journaux d'analyse des énigmes sont silencieux pendant la configuration du benchmark.
- Exécuter l'optimisé et la force brute sur l'entrée complète dans des commandes séparées évite une sortie incomplète lors de l'interruption de la force brute.

Ou exécutez le script helper :

```powershell
powershell -ExecutionPolicy Bypass -File ".\run_benchmarks.ps1"

# Ignorer la force brute sur l'entrée complète
powershell -ExecutionPolicy Bypass -File ".\run_benchmarks.ps1" -SkipBruteForce
```

Script Bash équivalent :

```bash
bash ./run_benchmarks.sh

# Ignorer la force brute sur l'entrée complète
bash ./run_benchmarks.sh --skip-bruteforce
```

## Explication de l'Algorithme

### Aperçu du Problème

Ce puzzle nécessite de mapper des ID de graines à travers une série de transformations (graine → sol → engrais → eau → lumière → température → humidité → localisation) en utilisant des mappages basés sur des intervalles définis dans un "Almanach". L'objectif est de trouver le numéro de localisation minimum parmi toutes les graines.

### Partie 1 : Mappage de Graines Individuelles (Force Brute)

**Approche :**
La solution naïve traite chaque graine séparément et enchaîne toutes les fonctions de transformation :

1. Parser les graines comme des valeurs individuelles
2. Pour chaque graine, appliquer le pipeline de transformation :
   - Rechercher quelle règle de mappage (le cas échéant) s'applique à la valeur actuelle
   - Si trouvée, traduire la valeur en utilisant : `destination + (valeur - source)`
   - Si aucune règle ne s'applique, la valeur se mappe à elle-même (mappage identité)
   - Passer à l'étape de transformation suivante
3. Collecter tous les numéros de localisation finaux et retourner le minimum

**Complexité Temporelle :** O(n × m) où n est le nombre de graines et m est le nombre total de règles de mappage sur tous les étages.

**Limitation :** La partie 2 spécifie des plages de graines (paires de départ + longueur), qui peuvent représenter des milliards de graines individuelles, rendant cette approche impraticable.

### Partie 2 : Pipeline d'Intervalles (Optimisé)

**Idée Clé :** Au lieu d'itérer chaque graine individuelle, nous pouvons propager des plages de graines à travers les mappages sous forme d'intervalles. Les intervalles présentant un comportement de transformation identique peuvent être regroupés et traités comme une unité unique.

**Structure de Données :**
- **Intervalle :** Une plage semi-ouverte `[début, fin)` représentant des plages continues de valeurs de graine/valeur
- **Les plages semi-ouvertes** simplifient l'arithmétique et la fusion : une plage exclut toujours son point final, évitant les erreurs de décalage d'un

**Étapes de l'Algorithme :**

1. **Parser les Plages de Graines :**
   - Interpréter les données de graines comme des paires : `(début, longueur)` → intervalle `[début, début + longueur)`
   - Exemple : les graines `79 14 55 13` deviennent les intervalles `[79, 93)` et `[55, 68)`

2. **Normaliser les Intervalles :**
   - Trier les intervalles par position de début
   - Fusionner les intervalles qui se chevauchent ou sont adjacents pour réduire la fragmentation
   - Exemple : `[79, 93)` et `[88, 100)` fusionnent en `[79, 100)`

3. **Appliquer le Mappage aux Intervalles (`applyMapToIntervals`) :**
   Pour chaque intervalle d'entrée, itérer à travers les règles de mappage (triées par début de source) :
   - **Identifier les lacunes :** Les parties de l'intervalle non couvertes par une règle de mappage conservent leur identité (pas de transformation)
   - **Gérer les chevauchements :** Pour les parties de l'intervalle qui se chevauchent avec une règle de mappage :
     - Calculer la région de chevauchement : `[max(intervalle.début, règle.source), min(intervalle.fin, règle.source + règle.plage))`
     - Traduire vers la destination : `valeur_mappée = règle.destination + (valeur_source - règle.source)`
     - Préserver la longueur de l'intervalle car le mappage est linéaire
   - Utiliser un curseur pour suivre la position dans l'intervalle et éviter les traitements en double

4. **Propagation du Pipeline :**
   - Commencer avec les intervalles de graine normalisés
   - Appliquer chacun des 7 mappages en séquence : graine-vers-sol, sol-vers-engrais, ..., humidité-vers-localisation
   - Normaliser les intervalles après chaque étape pour prévenir la croissance exponentielle

5. **Extraire le Résultat :**
   - Après toutes les transformations, examiner les intervalles de localisation résultants
   - Retourner la valeur de début minimale parmi tous les intervalles

**Exemple de Parcours :**
```
Intervalle de graine d'entrée [79, 80):
  Étape 1 : règle graine-vers-sol (98→50, plage 2) et (50→52, plage 48)
            → 79 tombe dans la plage [50, 98), se mappe à 52 + (79-50) = 81 → [81, 82)
  Étape 2 : règle sol-vers-engrais (15→0, plage 37) et (52→37, plage 2) et (0→39, plage 15)
            → 81 tombe dans la plage [52, 54), se mappe à 37 + (81-52) = 66 → [66, 67)
  ... (continuer à travers les étapes restantes)
  Final : intervalle de localisation [82, 83), minimum = 82
```

**Analyse de Complexité :**
- **Temps :** O(k × (n log n + n × m)) où k est le nombre d'étages de mappage (7), n est le nombre d'intervalles, m est le nombre de règles de mappage par étage
  - Le nombre d'intervalles croît comme O(n × m) par étage dans le pire cas
  - La normalisation (tri + fusion) le garde praticable
- **Espace :** O(n × m × k) pour stocker les intervalles à chaque étage du pipeline
- **Performance Pratique :** Sur l'entrée complète de l'énigme, cela s'exécute en millisecondes par rapport à des heures pour la force brute (qui itérerait ~600 milliards de graines)

**Pourquoi Cela Fonctionne :**
1. **Nature Linéaire des Mappages :** Chaque règle applique un décalage constant : si la valeur A se mappe à B, alors A+1 se mappe à B+1
2. **Préservation des Intervalles :** L'ensemble des intervalles qui n'ont "jamais rencontré une règle de mappage" peut être traité comme un intervalle unique à travers plusieurs étapes
3. **Normalisation :** La fusion réduit la prolifération des intervalles, en gardant la représentation compacte

Cet algorithme démontre comment la compréhension de la structure mathématique d'un problème (fonctions linéaires par morceaux) peut mener à des améliorations de performance dramatiques via un changement de représentation.
