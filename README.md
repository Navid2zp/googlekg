# googlekg
Google Knowledge Graph Search API bindings.

This package provides the API bindings for google knowledge graph search api.

Detailed API reference: https://developers.google.com/knowledge-graph/reference/rest/v1

## Install

```
go get -u https://github.com/Navid2zp/googlekg
```

## Example:
```
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

**Methods available:**

```
// To set "limit" parameter
req.SetLimit(limit int)

// To set "query" parameter
req.Query(query string)

// To set "indent" parameter
req.SetIndent(indent bool)

// To set "prefix" parameter
req.SetPreFix(prefix bool)

// To set "ids" parameters
req.SetIDs(ids []string)

// To set "languages" parameters
req.SetLanguages(languages []string)

// To set "types" parameters
req.SetTypes(types []string)
```

## License
MIT