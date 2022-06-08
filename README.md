# Godo
A simple Todo List API made for learning purposes.

## Made with
- [Golang](https://go.dev/)
- [Gorm](https://gorm.io/)
- [Gin](https://github.com/gin-gonic/gin)
- [SQLite](https://www.sqlite.org/index.html)

## Running on your local machine
```sh
go run . --migrate # If this is the first time you run the app
```

```sh
go run . --mode=release # To run the app in release mode
```
You can use an API Client like [Insomnia](https://insomnia.rest/) to test the API endpoints. There is a file inside `docs/` you can import to it that contains all requests the API accepts.
