# googlekg
Google Knowledge Graph Search API bindings.

This package provides the API bindings for google knowledge graph search api.

Detailed API reference: https://developers.google.com/knowledge-graph/reference/rest/v1

## Install

```
go get -u https://github.com/Navid2zp/googlekg
```

## Example:
```go
req, err := googlekg.Make("YOUR_API_KEY")
if err != nil {
    fmt.Println(err)
}
// Set a query string
req.SetQuery("github")

res, err := req.Do()
if err != nil {
    fmt.Println(err)
}

fmt.Println(res)

// Response will be a KG struct type when response status code is 200.
```

**Methods:**

```go
// To set "limit" parameter
req.SetLimit(10)

// To set "query" parameter
req.Query("github")

// To set "indent" parameter
req.SetIndent(true)

// To set "prefix" parameter
req.SetPreFix(true)

// To set "ids" parameters
req.SetIDs([]string{"/m/0dl567"})

// To set "languages" parameters
req.SetLanguages([]string{"en", "fa"})

// To set "types" parameters
req.SetTypes([]string{"Person", "Thing"})
```

## License
MIT
