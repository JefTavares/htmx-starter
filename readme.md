# Which technologies were used in this project??

- Go (https://go.dev/)
- air (https://github.com/cosmtrek/air)
- templ (https://templ.guide/)
- htmx (https://htmx.org/)
- bulma.css (https://bulma.io/)

# How to run the project?

Execute the make comand below to start 2 processes:

1. go web server running with air
2. templ live reload process

```bash
make start
```

Run the command below to stop all previously running processes

```bash
make stop
```

# Notas sobre o htmx

## hx-boost
```html
hx-boost="true"
```

## hx-get

Faz um get para o endereço

```html
<p hx-get="/users">Users</p>
```

E subustitui o `<p>` com todo o conteudo de `users.templ`

```go
	@page() {
		<h1>Users</h1>
		<a hx-boost="true" href="/">Home</a>
	}
```

Para resolver o problema 
```html
<p hx-get="/users" hx-target="body">Users</p>
```