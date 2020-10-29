Permet de démarrer rapidement un projet WordPress avec tous les outils pré-installer.

- Dernière version de **WordPress**
- Thème de roots/sage en version 9
- Installation des extensions **Yoast SEO, Contact Form 7, Classic Editor, Duplicate Post, Custom Post Type UI et ACF Pro**


## Usage

```
make-wp install
```

## Description

### Installation de WordPress

1) Téléchargement automatique de WordPress dans le répertoire désiré,
2) Création du fichier `wp-config.php`
3) Suppression des thèmes et extensions par défaut
4) Installation par défaut d'un fichier `wp-cli.yml` à la racine du projet pour démarrer le serveur en 127.0.0.1:8000

### Installation du thème Sage

1. Téléchargement de la dernière version de Sage.
2. Installation de wordplate/acf pour gérer les ACFs directement depuis le thème PHP.
3. Modification du fichier `resources/functions.php` pour y rajouter acf.
4. Création d'un répertoire `app/Fields` et d'un fichier `app/acf.php`.

### Installation de WordPress

1. Installation de base de WordPress
2. Nettoyage des données par défaut de WordPress
3. Mise en place de la structure de lien
4. Installation de toutes les extensions


## Todo

- Demander les informations lors de l'installation de WordPress