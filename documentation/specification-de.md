# Pflichtenheft 

Vorlesung: Anwendungsprojekt Informatik
Zuständige Dozierende: Fr. Dr. Kristina Birn & Hr. Prof. Dr. Erik Behrends
Kurs: TIF20A
Projekt: Webanwendung zur Vorlesungsplanung

Gruppenmitglieder:
- Bieringer, Christoph Karlhans Otto
- Krüger, Jennifer Laureen
- Kubon, Benedikt Daniel
- Meier, Jannic
- Weishar, Luc Rémi
- Zaremba, Fabian

## Einleitung

Ausgangssituation ist die zeitaufwändige manuelle Planung von Vorlesungsterminen mittels Tischkalender und manueller Absprache des Sekretariats mit einzelnen Dozenten.

- Zeitaufwendige manuelle Organisation und Verschiebung von Vorlesungsterminen
- Häufige Absprachen zwischen Dozent / Sekretariat
- Planung mit internen / externen Dozenten unterschiedlich
- Arbeitsmittel bisher: Tischkalender & Bleistift

Ziel des Projektes ist eine Minimierung des personellen Aufwandes zur Organisation der Vorlesungsterminen eines oder mehrerer Kurse.
Umgesetzt werden soll eine grundlegenden Konzeptionierung und die Implementierung eines entsprechenden Webanwendungsprototypen.
Dabei soll der zeitliche und organisatorische Aufwand um t-n Zeiteinheiten, gegenüber dem bisherigen Zeitaufwand t, reduziert werden.

Ein Prototyp / MVP soll bis einschließlich den 20. Juni 2021 durch den Lenkungskreis und das Kernteam, sowie mithilfe der internen und externen Fachpromotern, realisiert und den Stakeholdern zur weiteren Evaluierung zur Verfügung gestellt werden.

Besonderer Kundennutzen:
- Erhöhung der organisatorischen Effizienz (Verringerung des Zeitaufwandes)
- Steigerung der Attraktivität der DHBW für Dozenten
- Überführung analoger Planungsprozesse des Studienbetriebes in digitale Lösungenansätze zur Vereinheitlichung der Optimierung der Planungsabläufe

## Technische Lösung

### Rahmenbedingungen

- Umsetzung als Free / Libre Open Source Software unter MIT Lizenz
- Projektorganisation und Management mit OpenProject (Eigene Instanz auf bwCloud)
- Organisatorische Diagramme mit draw.io (https://app.diagrams.net/)
- Dokumentation im Markdown Format

### VCS & Hosting

- Öffentliche Versionskontrolle (VCS) mit Git, Hosting durch GitHub
- Continuous Integration (Build & Statische Analyse) mit GitHub Actions
- Serverbetriebssystem Debian 11 Bullseye
- Hostingplatform bwCloud (OpenStack, Verwaltung durch OpenSSH)

### Datenbank

- Relationales Datenbankschema & EER Modell mit MySQL Workbench
- Datenbankserver MariaDB

### Backend

- Backendentwicklung mit Go (memory-safe, high performance)
- Statische Analyse mit integriertem ```go vet``` bzw. ```gopls```
- Codeformatierung mit ```go fmt```
- Race Conditions Testing mit ```go build -race```
- Keine Abhängigkeiten von C Libraries oder Betriebssystem -> Cross-Platform / Cross-Architecture
- HTTP Routing Library github.com/go-chi/chi
- SQL Boilerplate & Mapping mit github.com/jmoiron/sqlx
- SAML2 / Shibboleth SSO Serviceprovider mit github.com/crewjam/saml
- REST API mit JSON Payloads zur Kommunikation mit Frontend
- TLS Mindestversion v1.2 mit Zertifikaten über ACME (CA Let's Encrypt)

### Frontend

- Erstellung von Mockups & Klickprototypen mit https://mockplus.com
- Frontend mit Javascript mit ES Modulen entwickelt nach ECMAScript2021 Standard
- Babel als Transpiler (mit preset-env und browserquery für target >= 0,25% Market & not dead)
- Webpack als Bundler für CSS & JS
- CSS & JS Komponenten Framework Bootstrap v5
- CalenderJS als angepasstes Modul für Kalenderentwicklung
- ESLint mit angepasster Konfiguration zur statischen Analyse
- JSBeautify zur einheitlichen Codeformatierung

- Webfont Roboto, konvertiert mit FontSquirrel
- Iconfont FontAwesome v5

## Beschreibung der Anforderungen bzw. Komponenten

### Mockups / Konzeption

- Konzeption des Userinterfaces
- Mockups zur Visualisierung und Unterstützung der Anwendungsentwicklung
- Strukturierte Entwicklung der User Experience

### Minimum Viable Product / Prototyp

- Backend
- Frontend

*Hier wird beschrieben, welche Anforderungen geliefert werden. 
Dies kann z.B. nach Teilprojekten gegliedert werden. 
Die Anforderungen werden beschrieben und können mit Screenshots von den Skizzen, Wireframes oder Mockups dargestellt werden. 
Eine Einteilung in die Pflichtanforderungen (MUST-HAVEs) und zusätzlich geplanten Anforderungen (NICE-TO-HAVEs) ist sinnvoll.*

*Durch die Beschreibungen soll der Umfang (Scope) eines sinnvollen Lieferobjekts festgelegt werden. 
Falls möglich können Akzeptanzkriterien angegeben werden, sodass dadurch der Nachweis erbracht wird, dass eine Anforderung erfüllt ist.*

## Bereitstellung der Mockups und der Webanwendung

- Bereitstellung und Auslieferung der Mockups erfolgt über Screenshots im Rahmen der Dokumentation sowie als Klickprototypen über das Cloudangebot von https://mockplus.com
- Der Prototyp der Webanwendung / das MVP wird funktionsfähig auf einer Serverinstanz der bwCloud bereitgestellt (siehe oben).

*Hier wird dargestellt wie das Ergebnis ausgeliefert und verwendet wird (Mockups oder Klickprototypen in der Cloud, Prototyp der Webanwendung usw.)*

## Qualität und Test

- Statische Analyse für Backend & Frontend Code
- Ausschluss von Sicherheitslücken durch OOB Fehler, Memory Safety Problemen und Type Confusion im Backend durch Auswahl der passenden Programmiersprache
- Kontinuierliche statische Prüfung von dynamisch typisiertem JS Code zur Reduzierung von Type Confusion
- Kontinuierliche ausführliche Code Reviews im Team
- Orientierung an OWASP Top 10 und Mitre CWE bei Entwicklung Code Review
- Manuelle Integrationstests mit festgelegten Testdaten und Überprüfung von Edge Cases

*Gibt es Qualitätskriterien? Wie werden Tests durchgeführt?*

## Projektplanung

### Terminplanung

### Meilensteine

### Einsatzplanung

| Rolle                | Person                    |
| -------------------- | --------------------------|
| Projektmanager       | Jennifer Krüger           |
| Stellvertreter       | Fabian Zaremba            |
| Technische Leitung   | Fabian Zaremba            |
| Lenkungskreis        | Krüger, Zaremba           |
| Frontendentwicklung  | Weishar, Meier, Kubon     |
| Backendentwicklung   | Zaremba                   |
| API Entwicklung      | Bieringer                 |

*Terminplan, Meilensteine, Einsatzplanung (Kapazitäten, Rollen, Zuständigkeiten)*

## Anhang

### Algorithmische Planung

- Betrieb mit regelmäßigen Updates über stabile Linuxdistribution für den Servereinsatz (Debian)
  - Bullseye Sicherheitsupdates bis voraussichtlich ca. 2024
  
- Go v1 Kompatibilitätsplanung (Link)  

- Weiterentwicklung der Planung von manueller Terminfestlegung
  - Festlegung von Constraints und Darstellung als SAT Problem
  - Verwendung etablierter SAT Solver zur approximativen zeitlich begrenzten Lösung (Partial MaxSAT)
  - Ansatzpunkte siehe auch "Solving the Course-timetabling Problem of Cairo University Using Max-SAT", Mohamed El Halaby (2018), https://arxiv.org/abs/1803.05027



*Begriffe und Definitionen
Hinweise zu Betrieb, Wartung und Weiterentwicklung*
