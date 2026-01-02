# Projet 3 : Quiz

Ce projet consiste en la création d'un jeu de quiz de culture générale (orienté jeux vidéo) développé en Go. Il se divise en deux étapes : une application en ligne de commande (CLI) et une interface web interactive.

---

## Phase 1 : Version Terminal (CLI)

La première phase se concentre sur la structuration des données et l'interaction séquentielle avec l'utilisateur.

### Objectif
Manipuler des structures de données complexes (`struct`, `slice`), gérer des boucles de jeu et valider les entrées utilisateur (choix multiples).

### Fonctionnalités
* **Niveaux de difficulté** :
    * Facile
    * Moyen
    * Difficile
* **Système de questions** : Affichage successif de questions avec 4 choix possibles.
* **Suivi du score** : Comptabilisation des bonnes réponses en temps réel.
* **Validation** : Vérification de la saisie (A, B, C ou D) pour éviter les erreurs.

### Installation et Lancement
1. Accéder au répertoire :
```bash
cd "Phase 1 - Quiz CLI"
```
2. Exécuter le programme :
```bash
go run main.go
```

### Utilisation
1. Choisir un niveau de difficulté dans le menu.
2. Répondre aux questions en saisissant la lettre correspondante (A, B, C, D).
3. Consulter le score final affiché à la fin de la série.

---

## Phase 2 : Extension Web (HTML + Go)

Cette phase porte le quiz sur une interface web, introduisant la gestion d'état entre les pages via des requêtes HTTP.

### Objectif
Créer une expérience utilisateur fluide où le score et la progression sont maintenus tout au long du questionnaire, en utilisant le modèle MVC.

### Fonctionnalités
* **Navigation par étapes** : Une question par page avec validation et passage automatique à la suivante.
* **Feedback final** : Page de résultats affichant le score et un message personnalisé selon la performance (de "Nul" à "GOAT").
* **Architecture MVC** : Séparation de la logique de présentation (`template`), de contrôle (`controller`) et de routage (`router`).
* **Contenu dynamique** : Chargement des questions selon le niveau choisi.

### Installation et Lancement
1. Accéder au répertoire :
```bash
cd "Phase 2 - Quiz Web"
```
2. Démarrer le serveur :
```bash
go run main.go
```
3. Accéder à l'application :
Ouvrir un navigateur à l'adresse : `http://localhost:8080`

### Structure du Projet Web
* **main.go** : Initialisation du serveur et des données.
* **controller/** : Logique de gestion du quiz (vérification réponses, calcul score).
* **router/** : Gestion des routes URL pour chaque niveau.
* **template/** : Fichiers HTML pour l'affichage des questions et résultats.
* **static/** : Feuilles de style CSS.

---

## Annexes : Thématiques

Le quiz porte principalement sur la culture vidéoludique :
* Histoire des consoles (PlayStation, Nintendo, etc.)
* Personnages emblématiques (Mario, Sonic, Link, etc.)
* Faits marquants de l'industrie.
