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
- my own codeg

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

| feature       | gorm | go-jet | bob | carta | sqlc                                                                       |
| ------------- | ---- | ------ | --- | ----- | -------------------------------------------------------------------------- |
| multi tenancy | -    | -      | -   | -     | workaround [discussion](https://github.com/sqlc-dev/sqlc/discussions/1108) |
