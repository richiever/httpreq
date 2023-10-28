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
R.Get() // returns the GET request in []byte
R.GetAsString() // Returns request as a string
```
## Post request Example 

```Go
R := RequestsPostRequestsPost{url: "https://jsonplaceholder.typicode.com/posts/1", postBody: map[string]string{"name": "Toby", "email": "Toby@example.com"}, requestType: "application/json"}
R.Post() // returns the POST request in type http.Response
R.GetAsString() // Returns request as a string
```



