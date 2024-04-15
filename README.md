### Compile with

```bash
go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o dist/go-olt
```

### config.json

```bash
{
    "db_filename": "/dir/of/sqlite/db"
}
```
