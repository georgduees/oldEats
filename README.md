<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
***
***
***
*** To avoid retyping too much info. Do a search and replace for the following:
*** georgduees, oldEats, georg_duees, georg@duees.de, Schmeckt wie Damals, Historisches digitales Kochbuch, alte Rezepte in neuem Format
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!--[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
-->


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/georgduees/oldEats">
    <img src="images/logo.png" alt="Logo" width="100" height="100">
  </a>

  <h3 align="center">Schmeckt wie Damals</h3>

  <p align="center">
    Historisches digitales Kochbuch, alte Rezepte in neuem Format
    <br />
    <a href="https://github.com/georgduees/oldEats"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/georgduees/oldEats">View Demo</a>
    ·
    <a href="https://github.com/georgduees/oldEats/issues">Report Bug</a>
    ·
    <a href="https://github.com/georgduees/oldEats/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Inhaltsverzeichnis</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">Über das Projektt</a>
      <ul>
        <li><a href="#built-with">Verwendete Software</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Einstieg</a>
      <ul>
        <li><a href="#prerequisites">Voraussetzungen</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Verwendung</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Unterstützung</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Danksagungen</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## Über das Projekt

[![Product Name Screen Shot][product-screenshot]](https://example.com)


### Verwendete Software

* []()
* []()
* []()



<!-- GETTING STARTED -->
## Einstieg

Um eine lokale Kopie des Projektes zu starten folge am besten den folgenden Schritten.

### Voraussetzungen

Dies ist eine List der Dinge, die vorhanden sein sollten, damit du das Projekt ausführen kannst.

* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/georgduees/oldEats.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```



<!-- USAGE EXAMPLES -->

## Verwendung

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_



<!-- ROADMAP -->
## Roadmap 

1. Daten aufarbeiten in PDF-Format
2. Mirroring/Ablage der Daten zum direkten Abruf über Server
3. Aufteilen der Daten in "Rezepte / Texte / Bilder" 
4. Aufbau API für Rezepte / Texte / Infos

# ToDo

# Aufgabe

Die Daten aus der Quelle müssen teilweise ergänzt werden und erweitert. Als Beispiel wird z.B. bei Backrezepten weder von Temperaturen als auch von der Backdauer geschrieben.
Dieses oder ähnliche Probleme sind typisch.
Damit die Daten zu "modernen und neuen Rezepten werden" würde ich vorschlagen das ganze in zwei Schritten zu managen.
Tipps / Weitere Informationen sollten von Nutzern/Besuchern angegeben werden können, ähnlich "Backtemperatur 180*C und 30 Minuten". 

1. Einmal die original-Daten wiedergeben als "Read-Only"-Schnittstelle

2. Die Schnittstelle nutzen um ein "Daten-Feeder" zu bauen/ Tool welches die Rezepte importiert in das "neue Rezept-Format / die Datenbank".
Im ersten Schritt wird so händisch Rezept für Rezept eingepflegt, anschließend wenn die Rezepte bestehen kann die Maske als "Editor genutzt werden".
Als Rezept-Formate gibt es z.B. das Open Recipe Format, welches yaml-Dateien anlegt.
3. Tipps / Kommentare können angegeben werden zu den Rezepten ähnlich Reddit-Threads mit Voting.
   Am meisten gevotete Tipps werden zu den Rezepten angezeigt und ggf. als "Backtemperatur" angegeben.
   Gleiches kann für Zutaten verwendet werden. "9/10 Nutzern ersetzten die Gurke durch eine Tomate".

4. API-Nutzung in Kochbuch-Website/Blog/Porta

# Aufbau

1. Aufbereitung OCR-Rescan / Fix der Stammdaten
2. Erst einmal eine Original-API, welche simpel die Texte und Bilder liefert Seite pro Seite im Buch.
   Grafik und Text aus OCR pro Buchseite
   Aufteilung  der Rezepte, Inhaltsverzeichnis der Bücher zur Hilfe.
3. Datenbank für Rezept Alage in neuem Format
    Open Recipe Format https://schema.org/Recipe
    https://schema.org/Recipe
    YAML-Files
4. Rezept-Editor/Korrektur-Tool:
    4.1 Rezepte korrigieren mit Webeditor:
    4.2 Anzeige Original-Rezept und Text aus OCR-Scan
    4.3 Eingabe / Zuweisung von Text zu "Rezepteigenschaften"/Daten zuweisung.
        Beispiel "1/4 Pfund Zucker". Wird markiert & übertragen zu Rezeptzutat in neuer Datenbank mit "Inforamtionen verknüpft".
5. API-Aufbau um Rezeptdaten zurückzugeben.
6. Blog/Webverzeichnis der Rezepte
7. Aufbereitung
    Ernährungsampel
    Nutrition facts

# Links

GitHub Link:
https://github.com/georgduees/oldEats

# Tools

 Datenbank PostGres/MariaDB


See the [open issues](https://github.com/georgduees/oldEats/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Unterstützen

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## Lizenz

Not yet anything licensed. Except the Project Icon:
## Ressource

* Project Icon <a target="_blank" href="https://icons8.de/icon/2470/essen">Essen</a> Icon von <a target="_blank" href="https://icons8.de">Icons8</a>


<!-- CONTACT -->
## Kontakt

Your Name - [@georg_duees](https://twitter.com/georg_duees) - georg@duees.de

Project Link: [https://github.com/georgduees/oldEats](https://github.com/georgduees/oldEats)



<!-- ACKNOWLEDGEMENTS -->
## Danksagungen

* []()
* []()
* []()





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/georgduees/repo.svg?style=for-the-badge
[contributors-url]: https://github.com/georgduees/oldEats/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/georgduees/repo.svg?style=for-the-badge
[forks-url]: https://github.com/georgduees/oldEats/network/members
[stars-shield]: https://img.shields.io/github/stars/georgduees/repo.svg?style=for-the-badge
[stars-url]: https://github.com/georgduees/oldEats/stargazers
[issues-shield]: https://img.shields.io/github/issues/georgduees/repo.svg?style=for-the-badge
[issues-url]: https://github.com/georgduees/oldEats/issues
[license-shield]: https://img.shields.io/github/license/georgduees/repo.svg?style=for-the-badge
[license-url]: https://github.com/georgduees/oldEats/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/georgduees
