# pokedex

Welcome to a tiny, terminal-friendly Pokédex written in Go. Perfect for explorers, collectors,
and anyone who loves pokémon-flavored command-line adventures. (I exaggerate)

Think of this repo as a pocket-sized CLI world where you can:

- `catch` — try your luck and capture pokémon
- `explore` — wander around and discover new areas
- `inspect` — get the lowdown on a specific pokémon
- `map` — see where you've been (or where you should go next)
- `pokedex` — open your gathered collection
- `help` — get command guidance
- `exit` — leave the adventure (for now)

Current ideas & TODOs

- Support the `up` arrow to cycle through previous commands
- Add more unit tests and tighten up test coverage
- Improve code organization for easier testing and maintenance
- Implement battles, leveling, and evolution
- Add autocompletion for areas and commands
- Add random encounters and richer exploration mechanics

Quick start

1. Build the binary:

```bash
go build -o pokedex
```

1. Run the CLI (or just run with `go run`):

```bash
./pokedex
# or
go run main.go
```

Try a few commands once inside the CLI:

```text
> explore
You wander into a sun-dappled clearing and spot a wild Pidgey.
> catch pidgey
Nice! You caught a Pidgey. It has been added to your `pokedex`.
> pokedex
1: Pidgey — Normal/Flying
```

Testing

Run unit tests for the helper packages:

```bash
go test ./...
```
