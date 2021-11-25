# Lunar Colony

## Description

Colonizing the Moon is going to be step one to colonize Mars. Phase 3 of the colonization would
require having some sort of economy.
Trading, traveling and other types of business will be the start of creating a sustainable economy.

## Project Requirements

Build a microservice that has the following requirement:

* Serves the Public API
* Takes time in UTC and converts to Lunar Standard Time
* Calculates time needed to deliver package from Earth to Moon

## Known Facts & Assumptions

* It takes 1 day from the warehouse to the Earth space station.
* The trip from the Earth to the Lunar Colony is about 3 days (warehouse -> space station NOT included)
* The time specified above in Earth days

## Software Requirements

Go 1.17.3 or above.
Docker & Docker-compose lts

## Installation

Run the following command to start the service:

```
make run
```

## Testing

Testing consists of: Unit Testing packages

All the test suits can be run with the following command:

```
make test
```