# Gedis

Gedis is an in-memory Redis inspired database written in Go.

## Why?

Just wanted to challenge myself to create a Database. + I just wanted to build a Go project since I am very much foreign with the language (not having used it in anything other than backend development)

## How does this work?

Its simple, a `map[string]string` stored in memory, while writing to cache for persistence.

Backups are also supported. Creating a backup will take the map from meory & spit out a `backup-[TIMESTAMP].(json|csv)` file with the map encoded as JSON or CSV. When loading a backup, the type of the file is inferred based on the file extension. If the name of a backup file is changed to include other stuff like extra ".", it will fail. After the type is inferred of either JSON or CSV, it decodes it & then places it in memory + cache.

It also starts a server by default on port 5000 (If PORT env variable is not set, the default is 5000). This API allows for `GET`, `SET` & `DELETE` so it can be accessed by an application using a network request.

## Commands

The database is very simple with 5 commands.

NOTE: The commands are case-insensitive, so doesn't matter if you are a SQL guy or not.

### Open commands

These commands can be accessed from both the CLI and the Rest API.

### `GET`

Gets a value with the given key from memory.

#### Syntax

```sh
GET <key>

# example, if the key is "foo"
GET foo
```

### `SET`

Sets a value with a given key. Also written to cache for persistence.

#### Syntax

```sh
SET <key> <value>

# example with the key="foo" & value="bar"
SET foo bar
```

### `DELETE`

Deletes a value with the given key.

#### Syntax

```sh
DELETE <key>

# example, delete the value with key "foo"
DELETE foo
```

### CLI-only commands

### `create_backup`

Creates a `backup-[TIMESTAMP].(json|csv)` in the `backups` folder.

Supported Backup export type

- JSON
- CSV

```sh
# by default, it exports as JSON
create_backup

# export as CSV
create_backup --csv

# any flag not-meaningful to the database is ignored (exported as JSON)
create_backup --never-gonna-give-you-up
```

### `load_backup`

Loads a backup file.

It looks for backups in the `backups/` folder so anything else is ignored.

#### Syntax

```sh
load_backup <backup-file-name>

# json
load_backup backup-060623-17:12:51.json

# csv
load_backup backup-060623-17:12:51.csv
```

## Using the Rest API

The database's server only has one endpoint, the root (/)
Here you can send `GET`, `POST` (`SET`) & `DELETE` Requests to perform actions on the database. Make sure to also pass `?key=...` & `&value=...` based on the type of action.

## Getting Started

- Clone the repo
```sh
$ git clone https://github.com/Dev-Siri/gedis.git
```

- Compile the project. Make sure you have [Go](https://go.dev) installed on your system.
```sh
$ go build -tags netgo -ldflags '-s -w' -o gedis
```

- Then just run the binary.
```sh
$ ./gedis
```

This will start a server on your machine w/ PORT env (default 5000) + a CLI that will allow you to interact with the database.

## License

This project is MIT licensed, see [LICENSE.md](LICENSE.md).