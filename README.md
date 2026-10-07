# wowhun

Félbehagyott Go-kísérlet 2024 végéről: a kódban egy küldetés- (`Quest`) és egy NPC-adatmodell,
valamint egy külső „wowapi" küldetés-végpont válaszának típusa van meg. Működő alkalmazás nincs
benne, és a kódból nem derül ki, pontosan mi lett volna a végcél.

**Nincs karbantartva:** az utolsó commit 2024-12-26-i, azóta nem változott. A munkafában ráadásul
nem commitolt módosítások vannak (`main.go`, `go.mod`, `go.sum`, `Dockerfile`,
`docker-compose.yml`).

## Felépítés

```
main.go               belépési pont: csak egy gorm + PostgreSQL megnyitás vázlata
dbmodel/              adatmodellek: Quest (cím, leírás, cél, jutalom, játékos neve/kasztja), Npc
ext/wowapi/           a külső API kliensének kezdete: a küldetés-válasz típusa (GetQuestResponse)
                      és a kliens-azonosítók
api/                  apimodel/, controller/, router/ – üres mappák
Dockerfile            golang:1.23.4-alpine, `go build -o main .`
docker-compose.yml    app (8080-as port) + PostgreSQL 13 (db_data kötet)
```

Függőségek (`go.mod`): Go 1.23.3, `gorm.io/gorm`, `gorm.io/driver/postgres`.

## Indítás

A terv szerint `docker compose up --build` indítaná: az `app` szolgáltatás a 8080-as portot
teszi ki, mellette egy `postgres:13-alpine` adatbázis fut, a kapcsolat adatai `DB_*` környezeti
változókban érkeznek.

**Jelenlegi állapotában nem fordul le:** a `main.go` nem deklarált változókra hivatkozik
(`db`, `err`, `connectionString`), a `DB_*` változókat semmi nem olvassa, és HTTP-szerver sincs
a kódban, tehát a 8080-as porton semmi nem hallgat.
