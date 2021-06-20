# Dokumentation zum Abschluss bzw. zur Übergabe

Vorlesung: Anwendungsprojekt Informatik

Zuständige Dozierende: Hr. Prof. Dr. Erik Behrends

Kurs: TIF20A

Projekt: Webanwendung zur Vorlesungsplanung

Gruppenmitglieder:
- Bieringer, Christoph Karlhans Otto
- Krüger, Jennifer Laureen
- Kubon, Benedikt Daniel
- Meier, Jannic
- Weishar, Luc Rémi
- Zaremba, Fabian


## Beschreibung der Installation

### Abhängigkeiten

- MySQL Datenbankserver
- Go Toolchain
- nodejs / npm
- git


### Installation & Ausführung

Annahme: MySQL Datenbankserver läuft auf localhost mit User mysql, Passwort mysql. Datenbank lectureplan enthält valide Daten, z.B. ```backend/res/demo.sql```.

- ```git clone https://github.com/fabian-z/dh-application-lectureplan.git``` 
- ```cd dh-application-lectureplan```
- ```cd frontend```
- ```npm install && npm run build-dev```
- ```cd ..```
- ```cd backend```
- ```go build && ./backend```

## Beschreibung des Abgabeinhaltes
### Was beinhalten die Mockups
#### Must-Haves
- Vorlesungszeitraum zur Planung festlegen
- Kursangabe beziehungsweise Kursauswahl
- Planungsmöglichkeiten der Vorlesung durch die Möglichkeit...
  - ...noch zu planende Vorlesungen mit Namen, Stundenzahl und ggf. mit Prüfungsleistungen zu versehen  
  - ...Termine (innerhalb der jeweiligen Theoriephase) einzutragen, zu berbeiten und zu löschen
  - ...Termine im Kurskalender in Wochen- und/oder Monatsansicht erscheinen zu lassen
  - ...Termine mit einer Start- und einer Endzeit zu versehen, denen ein Vorlesungstitel und den Namen des Dozierenden zu hinterlegen
  - ..., dass Vorlesungen lediglich im Rahmen von 08:00 - 17:00 Uhr zu veranstalten sind
  - ...das mindestens eine Stunde Mittagspause (zwischen 11:30 - 14:00 Uhr) einzuplanen ist
  - ...der Erkennung, ob alle Vorlesungsstunden eingeplant sind
- Einhaltung der DSGVO für Produktivdaten


#### Nice-to-haves
- Planung für mehrere Kurse
- Login mit Benutzername und Passwort
- Entwicklung von Prozessen zur Benachrichtigung, Erinnerung und Bestätigung von Terminen
- Raumplanung
- Export als PDF zum Drucken bzw. Versand per E-Mail
- Vorbereitung von Datenstrukturen zur algorithmischen Planung


### Bereitstellung der Mockups
- Bereitstellung und Auslieferung der Mockups erfolgt sowohl im Rahmen des [Pflichtenhefts](https://github.com/fabian-z/dh-application-lectureplan/blob/main/documentation/specification-de.md), sowie über das Cloudangebot [Mockplus](https://mockplus.com).
- Übersicht über die Konzeptionierung der einzelnen verlinkten Seiten sind [hier](https://app.mockplus.com/app/share-854623c46b21e3a6e67afbdf44080ea2share-KWYdGLTktnR/rp?hmsr=share) einzusehen
- Abgeschlossene Mockups stehen [hier](https://app.mockplus.com/run/rp/qoamv943rfei/OEVKeyYMGodP?ps=0&ha=0&la=0&fc=1&out=1) zur Einsicht


### Was wird als Code abgeliefert?

#### Frontend

- Prototyp der Kalenderkomponente
  - Integration mit REST API (Anzeige von Terminen aus Datenbank)
  - Anzeige von Wochen / Monaten / Terminlisten
- Modal Fenster
  - Umsetzung "Termin hinzufügen" (Backend Action TODO)
- Sidebar Navigation (Offscreen Canvas)
- SSO Authentication (SAML2)
- Custom Themes / Farben (Bootstrap SCSS)
- Toolchain Configuration (NPM, Webpack, Babel, ESLint, JSBeautify, SASS)
- HTML 5 Templates
- CSS3 Stylesheets, teilweise externes SCSS

#### Backend

- Go Backend
- Datenbank Boilerplate (SQL Queries)
- SSO Authentication
- HTTP Server inkl. ACME Client (TLS über Let's Encrypt)
- AJAX REST API

## Tätigkeitsnachweis der einzelnen Teammitglieder
Die gewünschte Übersicht über alle Tätigkeiten pro Gruppenmitglied ist in zwei Dokumentationen unterteilt.

[Hier](https://docs.google.com/spreadsheets/d/1lBy9JIgU5GNE6L-JWj5fWvleWRICEysQC48-V7u21hA/edit#gid=1636513249)
ist eine grobe Auflistung aller Aufgaben mit den entsprechenden Zeitaufwendungen zu finden, sowie eine Anzeige der gesamten Arbeitszeit pro Teammitglied,
welche für das Anwendungsprojekt angefallen ist.

Zusätzlich dazu sind auf [OpenProject](https://docs.google.com/spreadsheets/d/1lBy9JIgU5GNE6L-JWj5fWvleWRICEysQC48-V7u21hA/edit#gid=1636513249) die verschiedensten Arbeitspakete aufgelistet, welche während des Anwendungsprojektes angefallen sind. 
Hierzu gibt es eine entsprechende anzeige des Fortschrittes der zugewiesenen Aufgaben.
Wichtig dabei ist, dass der Filter gänzlich entfernt werden muss, da sonst die Arbeitspakete nicht angezeigt werden, welche schon abgeschlossen sind.


## Rückblick
### Welche Probleme sind aufgetreten?
#### Projektmanagement

- Fehlendes Fachwissen der Teammitglieder, durch nicht vorhandene Ressourcen führten zwischenzeitlich zu einen sehr geringen 
  Fortschritt einzelner Aufgaben bzw. des gemeinsamen Projektes
- Technische- bzw. Verständnisprobleme während der Nutzung des PM-Tools OpenProject führten zeitweise zu einer Verzögerung der Projektplanung

#### Frontend

- Unerwartete Komplexität im verwendeten Framework Bootstrap (SCSS Toolchain notwendig)
- Zu geringer Zeitrahmen für Ausbau des kompletten Prototypen

#### Backend

- Fehleinschätzung der Vorkenntnisse einzelner Gruppenmitglieder
- Fehlerhafte Registrierung des SSO (Shibboleth Fehler "Nicht unterstützte Anwendung", Erfolreiche Umsetzung mit samltest.id)

### Was wird nicht mit geliefert (Not-Doing)?
- E-Mail Ersatz für Individualkommunikation
- Dozierendenmanagement
- Responsive Design (für Geräte mit kleineren Displays bzw. Smartphones im Hochkantformat)
- Backup der Daten
- Benachrichtigungen und Erinnerungen für das Sekretariat und/oder der Dozierenden (z.B. per E-Mail)
- Schnittstellen zu anderen Anwendungen (Veröffentlichung über Exchange, Moodle, usw.)


### Was wird anders geliefert?
- SAML Identity Provider samltest.id statt idp.dhbw-loerrach.de - Registrierung der bereitgestellten Metadataten offenbar fehlerhaft (Nicht unterstützte Anwendung)
- Konzeption vollständig, Prototyp begrenzt auf Kalenderkomponente und Templates

## Ausblick
### Auflistung offener Punkte
- Einbindung des SAML Identity Provider idp.dhbw-loerrach.de
- Einführung optimaler Vorlesungszeiten (von 09:00 - 12:15 Uhr und von 13:00 - 16:15 Uhr)
- Usability und Vereinfachung: Vorlesungen beginnen und/oder enden immer :00; :15; :30; :45; (mithilfe von Autovervollständigung oder Suchfunktion)


### Zukünftige Aufgaben zur Weiterentwicklung
- Implementierung der "Out-of-Scope" bzw. "Not-Doing" Aufgaben
  - E-Mail Ersatz für Individualkommunikation
  - Dozierendenmanagement
  - Responsive Design (für Geräte mit kleineren Displays bzw. Smartphones im Hochkantformat)
  - Backup der Daten
  - Benachrichtigung und Erinnerungen für das Sekretariat und/oder der Dozierenden (z.B. per E-Mail)
  - Schnittstellen zu anderen Anwendungen (Veröffentlichung über Exchange, Moodle, usw.)
