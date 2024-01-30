# CRUD and data mapping

## interesting projects

- <https://github.com/cdfmlr/crud>
- <https://betterprogramming.pub/go-generic-repo-fd31b0300e0e>
- Gorm
- Go-Jet
- SQ
- Bob
- Carta
- sqlc
- SQLBoiler
- my own codeg

## Features

| short | function                                              |
| ----- | ----------------------------------------------------- |
| C     | Create                                                |
| R     | Read one                                              |
| U     | Update one                                            |
| D     | Delete one                                            |
| P     | Partial update                                        |
| L     | List many                                             |
| A     | Audit or some kind of hookable api to implement audit |
| H     | Hooks for audit or timestamps                         |
| M     | Multi create                                          |

| Project | CRUDPLAHM        |
| ------- | ---------------- |
| codegen | C-UDPLA-         |
| Gorm    | CRUD-L-H         |
| Go-Jet  | -R---L--         |
| SQ      | almost plain sql |
| Bob     |                  |

## questions

- how to organizes repos?
- how to handle partial updates?
- how to handle audit?
- how to handle multi tenancy?

## features to look at

- simple crud code
- partial updates
- hooks (e.g. audit or timestamps)

## decision table

| feature       | gorm                                                      | go-jet | bob | carta | sqlc                                                                       |
| ------------- | --------------------------------------------------------- | ------ | --- | ----- | -------------------------------------------------------------------------- |
| multi tenancy | not easy. https://github.com/bartventer/gorm-multitenancy | -      | -   | -     | workaround [discussion](https://github.com/sqlc-dev/sqlc/discussions/1108) |
