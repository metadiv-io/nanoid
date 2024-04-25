# NanoID Generator

## Installation

```bash
go get -u github.com/metadiv-io/nanoid
```

## Generate NanoID

```go
uuid := nanoid.New(&nanoid.Opt{
    Numbers: true, // 0123456789
    Lowercase: true, // abcdefghijklmnopqrstuvwxyz
    Uppercase: true, // ABCDEFGHIJKLMNOPQRSTUVWXYZ
    Symbols: true, // _-
    Length: 21, // must be greater than 0
    ExcludeAlike: true, // exclude alike characters: 1lI0Oouv5Ss
})
```

## Safe Generate NanoID

Generate NanoID with safe characters:

1. Numbers
2. Lowercase
3. Uppercase
4. Exclude Symbols
5. Exclude Alike Characters
6. Length: 21

```go
uuid := nanoid.NewSafe()
```
