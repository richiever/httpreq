# HTTPREQ
#### A simple and easy to use module for HTTP requests in Golang.

## Installation:

### To install, use the command:
```cmd
go get github.com/richiever/httpreq
```

and proceed with the command:
```Go
import "github.com/richiever/httpreq"
```

## Or alternatively:
```cmd
git clone https://github.com/richiever/httpreq
```
and simply add the files to your project directory.
## GET request Example 

```Go
R := Requests{"https://jsonplaceholder.typicode.com/posts/1"}
R.Get() // returns the request in []byte
R.GetStringify(R.Get()) // Returns request as a string
```
