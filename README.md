# HTTPREQ
#### A python-like HTTP request module for the GO language.

## Installation:

### To install, use the command:
```cmd
go get github.com/richiever/httpreq
```

## GET request Example 

```Go
R := Requests{"https://jsonplaceholder.typicode.com/posts/1"}
R.Get() // returns the request in []byte
R.GetStringify(R.Get()) // Returns request as a string
```
