
Goa version v3.7.2

[Design](./design/design.go)

The expected `metadata` Attribute of `MapOf(String, String)` type is
mapped to the HTTP transport as the following `MapParams("metadata")`

The `metadata` query parameter can be passed to the URL as the following:
`GET /multiply/?metadata[testKey]=val`

The service will correctly receive `map[testKey:val]`, but the [OpenAPI generated](./gen/http/openapi3.yaml#L15)
it's not listing `metadata` in the endpoint.

## Expected

`metadata` query parameter to be exposed in OpenAPI like the other mapped attributes, example:

```
      parameters:
      - in: query
        name: metadata
        description: this should be on openapi
        required: false
        schema:
          type: object
          additionalProperties: true
          example:
            testKey: val
        style: deepObject
```


### Additional

### sample CLI issue

The sample CLI that is auto generated via goa example `goa example calcsvc/design` is producing a wrong URL (removing attribute name) when calling an endpoint with MapOf/MapParams, example:

```
./calc-cli calc multiply --type 1 --metadata '{"testKey": "val"}'
```

creates the following URL: `GET /multiply/?testKey=val&type=1` but instead should create: `GET /multiply/?metadata[testKey]=val&type=1`
