# Spécifications

## **1. Contexte et objectifs**

**Public cible** : Développeurs, utilisateurs soucieux de sécurité, ou toute personne ayant besoin de générer des mots de passe robustes.
**Objectif principal** : Créer une application de type **TUI** (_Text User Interface_) permettant de générer des mots de passe aléatoires, personnalisables et sécurisés, avec une interface intuitive et des options de copie/export. Cette application est également créée à des fins pédagogiques afin de monter en compétence sur le langage [Go](https://go.dev/), [BubbleTea](https://github.com/charmbracelet/bubbletea) et éventuellement [LipGloss](https://github.com/charmbracelet/lipgloss).

### **Vision produit**
Le produit vise à offrir un générateur de mots de passe en terminal qui soit :
- **Simple** : utilisable en quelques secondes sans apprentissage complexe.
- **Fiable** : fondé sur des choix techniques cohérents avec un usage sécurité.
- **Pédagogique** : conçu de manière à favoriser la compréhension du code, des choix d'architecture et du fonctionnement de BubbleTea.

### **Objectifs du MVP**
- Proposer une interface terminal simple et rapide à prendre en main.
- Générer des mots de passe robustes à partir de critères configurables.
- Donner un retour immédiat sur la qualité du mot de passe généré.
- Poser une base de code modulaire et testable pour les évolutions futures.
- Servir de support d'apprentissage concret pour progresser en Go et sur le modèle d'architecture BubbleTea.

### **Problème utilisateur**
Les générateurs de mots de passe existants sont souvent soit trop minimalistes, soit trop riches pour un besoin simple. Le projet cherche à proposer une alternative locale, rapide et lisible, orientée terminal, qui permette de :
- générer un mot de passe sûr sans ouvrir un navigateur ;
- comprendre immédiatement les paramètres utilisés ;
- obtenir un retour clair sur la robustesse du résultat.

### **Personas cibles**
- **Développeur en terminal** : veut générer rapidement un mot de passe robuste sans quitter son environnement de travail.
- **Utilisateur sensibilisé à la sécurité** : veut maîtriser les critères de génération et éviter les mots de passe faibles.
- **Apprenant Go/BubbleTea** : veut comprendre comment structurer une application TUI propre, testable et maintenable.

### **Hors périmètre du MVP**
- Synchronisation avec un service distant ou un compte utilisateur.
- Persistance longue durée de l'historique.
- Chiffrement avancé ou gestion complète de coffre-fort de mots de passe.
- Intégration native avec des gestionnaires de mots de passe tiers.

---

## **2. Fonctionnalités principales**

### **Fonctionnalités de base (MVP)**
| ID  | Fonctionnalité                          | Description                                                                                     | Priorité |
|-----|----------------------------------------|-------------------------------------------------------------------------------------------------|----------|
| F01 | Génération de mot de passe aléatoire   | Génère un mot de passe aléatoire selon des critères (longueur, types de caractères).             | Haute    |
| F02 | Personnalisation des critères          | Permet à l’utilisateur de choisir : longueur, inclusion de majuscules, minuscules, chiffres, symboles. | Haute    |
| F03 | Affichage du mot de passe              | Affiche le mot de passe généré en gros dans l’interface.                                       | Haute    |
| F04 | Copie dans le presse-papiers | Permet à l’utilisateur de copier le mot de passe généré dans le presse-papiers (si possible). | Moyenne  |
| F05 | Évaluation de la force du mot de passe | Affiche une estimation de la force (faible, moyenne, forte) en fonction des critères.         | Moyenne  |
| F06 | Génération multiple                    | Permet de générer plusieurs mots de passe d’un coup (ex. : 5 mots de passe).                   | Basse    |
| F07 | Sauvegarde dans un fichier            | Option pour sauvegarder les mots de passe générés dans un fichier texte/chiffré.               | Basse    |

### **Fonctionnalités avancées (évolutions possibles)**
| ID  | Fonctionnalité                          | Description                                                                                     |
|-----|----------------------------------------|-------------------------------------------------------------------------------------------------|
| F08 | Mode "exclusion de caractères"         | Permet d’exclure certains caractères (ex. : `l`, `1`, `O`, `0`).                              |
| F09 | Génération de phrases de passe         | Génère des phrases de passe (ex. : `CorrectHorseBatteryStaple`).                                |
| F10 | Historique des mots de passe           | Stocke et affiche l’historique des mots de passe générés pendant la session.                     |
| F11 | Thèmes visuels                         | Permet de changer les couleurs/thèmes de l’interface.                                          |
| F12 | Intégration avec un gestionnaire       | Export direct vers un gestionnaire de mots de passe (ex. : Bitwarden, KeePass).                 |

### **Périmètre retenu pour la première version**
La première version doit couvrir en priorité : **F01**, **F02**, **F03** et **F05**.

**F04** peut être implémentée si une bibliothèque multiplateforme fiable est retenue sans complexifier excessivement le projet. Les fonctionnalités **F06** à **F12** sont considérées comme des évolutions.

### **Valeur produit du MVP**
Le MVP doit permettre à un utilisateur de lancer l'application, choisir ses critères, générer un mot de passe robuste et comprendre immédiatement si le résultat est satisfaisant. Si cette promesse n'est pas tenue de manière simple et fluide, le MVP n'est pas considéré comme réussi.

### **User stories MVP**
- En tant qu'utilisateur, je veux choisir la longueur du mot de passe afin d'adapter le résultat aux contraintes du service visé.
- En tant qu'utilisateur, je veux activer ou désactiver certaines catégories de caractères afin de maîtriser le niveau de complexité.
- En tant qu'utilisateur, je veux générer un mot de passe en une action afin d'obtenir rapidement un résultat exploitable.
- En tant qu'utilisateur, je veux voir une indication de force afin d'évaluer immédiatement si le mot de passe semble suffisamment robuste.
- En tant qu'utilisateur, je veux recevoir un message clair si ma configuration est invalide afin de la corriger sans quitter l'application.

### **Règles métier**
- La longueur minimale d'un mot de passe doit être définie et validée. Proposition : **8 caractères**.
- La longueur maximale doit rester raisonnable pour l'affichage TUI. Proposition : **128 caractères**.
- Au moins une catégorie de caractères doit être sélectionnée avant toute génération.
- Si plusieurs catégories sont sélectionnées, le mot de passe généré doit garantir la présence d'au moins un caractère de chaque catégorie active.
- La génération doit être basée exclusivement sur une source aléatoire cryptographiquement sûre.
- L'évaluation de force doit être cohérente avec les critères retenus et reproductible à entrée identique.

---

## **3. Spécifications techniques**

### **Langages et librairies**
- **Langage** : Go **1.26.2**.
- **TUI** : [BubbleTea](https://github.com/charmbracelet/bubbletea) (pour l’interface) et [LipGloss](https://github.com/charmbracelet/lipgloss).
- **Génération aléatoire** : `crypto/rand` (pour une génération sécurisée).
- **Parsing/Validation** : Bibliothèque standard (`strings`, `unicode`).
- **Fichiers** : Bibliothèque standard `encoding/json` ou `os` (pour la sauvegarde).
- **Presse-papiers** : Bibliothèque à déterminer quand on en aura besoin.

### **Principes d'implémentation**
- Favoriser un code explicite et lisible plutôt qu'une implémentation trop compacte.
- Limiter la complexité prématurée : une première version simple et propre est préférable à une architecture surdimensionnée.
- Isoler la logique métier de l'interface pour rendre le code testable et compréhensible.
- Documenter les choix non évidents au niveau du code ou de la documentation projet.

### **Architecture logique suggérée**
- **UI/TUI** : gestion de l'affichage, de la navigation et des raccourcis clavier.
- **Domaine** : structures de configuration, génération, évaluation de la force.
- **Infrastructure** : presse-papiers, sauvegarde fichier, sérialisation éventuelle.

Cette séparation permet de tester la logique métier sans dépendre de l'interface terminal.

### **Structure de projet recommandée**
La structure suivante est recommandée pour rester idiomatique Go, compatible BubbleTea et favorable au TDD :

```text
password-generator/
├── cmd/
│   └── password-generator/
│       └── main.go                # Point d'entrée
├── internal/
│   ├── app/
│   │   ├── model.go               # Model BubbleTea (state, update)
│   │   ├── view.go                # Rendu BubbleTea
│   │   └── keymap.go              # Raccourcis clavier
│   ├── domain/
│   │   ├── password/
│   │   │   ├── generator.go
│   │   │   ├── generator_test.go
│   │   │   ├── strength.go
│   │   │   └── strength_test.go
│   │   └── rules/
│   │       ├── validator.go
│   │       └── validator_test.go
│   ├── config/
│   │   └── defaults.go            # Valeurs par défaut (longueur, options)
│   └── infra/
│       └── clipboard/
│           ├── clipboard.go       # Interface + implémentation
│           └── clipboard_stub.go  # Fallback/no-op si non supporté
├── docs/
│   └── decisions.md               # Décisions d'architecture et choix pédagogiques
├── go.mod
├── go.sum
└── README.md
```

### **Positionnement par rapport à la proposition initiale**

La proposition initiale de structuration du projet été la suivante :

```text
password-generator/
├── cmd/            # Point d'entrée de l'application
│   └── main.go
├── internal/
│   ├── model/      # Modèles BubbleTea (état, logique)
│   │   └── model.go
│   ├── view/       # Affichage (vues TUI)
│   │   └── view.go
│   ├── utils/      # Utilitaires (génération, validation, etc.)
│   │   ├── generator.go
│   │   ├── validator.go
│   │   └── clipboard.go
│   └── config/     # Configuration (critères par défaut, thèmes)
│       └── config.go
├── pkg/            # Code réutilisable (si extension future)
├── go.mod          # Dépendances Go
└── README.md       # Documentation
```

- La proposition initiale est globalement conforme à un projet Go + BubbleTea.
- Le dossier `cmd/password-generator/` est une convention Go fréquente et facilite une évolution vers plusieurs binaires.
- Le dossier `internal/model` + `internal/view` peut fonctionner, mais les regrouper dans `internal/app` simplifie la lecture du flux BubbleTea (`model/update/view`) pour l'apprentissage.
- Le dossier `utils` est pratique au début ; à moyen terme, le remplacer par des packages orientés métier (`domain/password`, `domain/rules`) améliore la clarté et les tests.
- Le dossier `pkg` n'est utile que si une API réutilisable externe est réellement prévue ; sinon, il peut être omis au départ.

### **Conventions de tests (TDD)**
- Écrire les tests au plus près du code testé (`*_test.go` dans le même package).
- Prioriser les tests du domaine (`generator`, `validator`, `strength`) avant les détails d'interface TUI.
- Conserver une couverture fonctionnelle des règles métier critiques : longueur, catégories sélectionnées, présence des catégories, robustesse.
- Limiter les tests d'interface aux comportements essentiels (navigation clavier, affichage d'erreurs, mise à jour du résultat).

### **Démarche de développement**
- Le projet suit une approche **TDD** (_Test-Driven Development_) autant que possible sur la logique métier.
- Pour chaque règle métier importante, écrire d'abord un test qui décrit le comportement attendu, puis implémenter le minimum de code nécessaire pour le faire passer.
- Réserver les tests les plus fins et systématiques au domaine : validation de configuration, génération, présence des catégories sélectionnées, évaluation de force.
- Utiliser l'interface TUI comme couche d'orchestration, avec une logique métier déjà sécurisée par les tests.

### **Contraintes pédagogiques**
- Le projet a un objectif explicite de montée en compétence sur Go et BubbleTea.
- Les choix d'implémentation doivent rester accessibles à la relecture et à l'apprentissage.
- Deux modes de contribution sont privilégiés :
   - soit le code est écrit directement par le porteur du projet ;
   - soit toute contribution externe doit rester suffisamment claire pour être comprise, relue et expliquée facilement.
- Les abstractions inutiles, les raccourcis trop implicites et les dépendances superflues sont à éviter.
- Chaque étape importante doit pouvoir être justifiée simplement : ce qui a été fait, pourquoi cela a été fait, et comment cela fonctionne.

### **Exigences non fonctionnelles**
- **Portabilité** : Doit fonctionner sur Windows, Linux et macOS.
- **Sécurité** : Utiliser `crypto/rand` pour la génération aléatoire (pas de `math/rand`).
- **Extensibilité** : Code modulaire pour ajouter facilement de nouvelles fonctionnalités.
- **Tests** : Tests unitaires pour la logique de génération et d’évaluation.
- **Performance** : La génération d'un mot de passe unitaire doit être quasi instantanée à l'échelle utilisateur.
- **UX** : Les actions principales doivent être accessibles au clavier sans souris.
- **Compréhensibilité** : Le code doit rester lisible pour une personne en apprentissage sur Go et BubbleTea.

### **Critères de validation technique**
- Le projet compile sans erreur sur au moins un environnement cible.
- Les tests unitaires couvrent les cas nominaux et les cas d'erreur de validation.
- Aucun usage de `math/rand` n'est autorisé dans le chemin de génération des mots de passe.
- Les messages d'erreur affichés à l'utilisateur doivent être explicites et non bloquants pour l'interface.
- Les éléments essentiels de la logique métier peuvent être expliqués simplement à l'oral ou à l'écrit.

---

## **4. Maquettes et flux utilisateur**

### **Flux principal**
1. L’utilisateur lance l’application.
2. L’interface affiche un menu avec les options de personnalisation (longueur, types de caractères).
3. L’utilisateur valide les critères.
4. Le mot de passe est généré et affiché.
5. L’utilisateur peut :
   - Copier le mot de passe.
   - Le sauvegarder.
   - Générer un nouveau mot de passe.
   - Quitter l’application.

### **Parcours utilisateur cible**
- L'utilisateur comprend l'écran principal sans lire de documentation préalable.
- Il identifie immédiatement les paramètres modifiables et l'action principale.
- Il obtient un résultat en un nombre d'actions réduit.
- Il comprend sans ambiguïté si le mot de passe généré est plutôt faible, moyen ou fort.
- Il sait comment recommencer ou quitter sans hésitation.

### **Comportements attendus de l'interface**
- L'utilisateur peut naviguer entre les champs avec le clavier.
- Une action de génération est possible via un raccourci clair, par exemple `Entrée` ou `g`.
- Une action de sortie est disponible via `q` et éventuellement `Ctrl+C`.
- En cas de configuration invalide, l'interface affiche un message d'erreur lisible sans quitter l'application.
- Après génération, le mot de passe et son niveau de force sont immédiatement mis à jour à l'écran.

### **Principes UX produit**
- Mettre en avant l'action principale : générer.
- Réduire la charge cognitive en affichant uniquement les informations utiles.
- Préserver un vocabulaire simple et cohérent dans tous les libellés.
- Faire apparaître clairement l'état courant de la configuration.
- Éviter toute ambiguïté entre action automatique et action manuelle, notamment pour la copie ou la sauvegarde.

### **Exemple de maquette TUI**
```
+-------------------------------------+
|       GÉNÉRATEUR DE MOT DE PASSE    |
+-------------------------------------+
| Longueur : [12]                     |
| Majuscules : [✓]                    |
| Minuscules : [✓]                    |
| Chiffres    : [✓]                   |
| Symboles   : [✓]                    |
|                                     |
| [Générer]    [Quitter]              |
+-------------------------------------+
| Mot de passe : 7H#k9Lm$2pQ!         |
| Force        : Forte                |
|                                     |
| [Copier]      [Sauvegarder]         |
+-------------------------------------+
```

---

## **5. Critères d'acceptation du MVP**

### **Fonctionnels**
- L'utilisateur peut définir une longueur de mot de passe dans une plage valide.
- L'utilisateur peut activer ou désactiver les catégories de caractères disponibles.
- Un mot de passe valide est généré à la demande et affiché dans l'interface.
- Le niveau de force est recalculé et affiché après chaque génération.
- Si la configuration est invalide, la génération est refusée avec un message explicite.

### **Techniques**
- La logique de génération est testée automatiquement.
- La logique d'évaluation de la force est testée automatiquement.
- Le code UI reste découplé de la logique métier.

### **Ergonomie**
- L'application est entièrement utilisable au clavier.
- Les libellés principaux sont compréhensibles sans documentation externe.
- L'affichage reste lisible sur une taille de terminal standard.
- Le parcours principal de génération reste compréhensible pour un nouvel utilisateur en moins d'une minute.

### **Pédagogie**
- Les composants principaux du projet peuvent être présentés simplement : configuration, génération, score de force, interface TUI.
- Une personne reprenant le projet peut comprendre la structure générale sans connaissance préalable avancée de BubbleTea.
- Les tests servent aussi de documentation exécutable sur les comportements attendus.

---

## **6. Questions ouvertes**
- Quelle bibliothèque de presse-papiers choisir pour garantir un comportement homogène sur Windows, Linux et macOS ?
- Faut-il intégrer la copie dans le MVP, ou la traiter comme une première évolution si son support multiplateforme s'avère trop coûteux ?
- Quelle méthode d'évaluation de la force retenir : heuristique simple basée sur les critères, ou estimation plus avancée de type entropie ?
- Le format de sauvegarde doit-il être du texte brut, du JSON, ou cette fonctionnalité doit-elle rester hors MVP ?

---

## **7. Indicateurs de réussite**
- Le MVP permet de générer un mot de passe valide en quelques interactions seulement.
- Le projet reste assez simple pour être codé ou repris sans perdre sa valeur pédagogique.
- Les tests guident la progression du développement au lieu d'être ajoutés uniquement à la fin.
- La structure du code facilite l'ajout ultérieur de fonctionnalités comme la copie, la sauvegarde ou l'historique.
