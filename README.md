# golang-greetings

Modulo que saluda

Parte del estudio de Go: [https://go.dev/doc/](https://go.dev/doc/)

Al inicio no tengo `go.mod` ni `go.sum`. Ejecuto:

```zsh
go mod init github.com/LuisPalacios/golang-greetings
go mod tidy
```

## Publica

Verifico que pasa los tests

```zsh
go test ./...   ((o también "go test -v")
```

Publico mi primera versión inestable,

```zsh
git add .
git commit -m "changes for v0.1.1"
git tag v0.1.1
```

Push del TAG a Origin

```zsh
git push origin v0.1.1
```

Hago que el módulo esté disponible

```zsh
go list -m github.com/LuisPalacios/golang-greetings@v0.1.1
```

## Consumo

En el proyecto que consume este módulo basta con incluir lo siguiente en su `import`

```go
    import (
        // mi módulo
        "github.com/LuisPalacios/golang-greetings"
    )
```

Y por supuesto haber inicializado y hecho un `mod tidy`

```zsh
go mod init luispa.com/hello
go mod tidy
```
