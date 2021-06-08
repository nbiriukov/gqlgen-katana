# GQLGen Katana v2

This repo is a PoC for using an SDL-first GraphQL server using the gqlgen library for Go.

## How to code

Do not forget to generate stuff after schema.graphqls or schema.resolvers.go changes! Go to country folder and run script in terminal:

```
$ cd graphql/us
$ go run github.com/99designs/gqlgen generate
```

## How to test

Start server:

```
$ make run_server
```

Visit either `http://localhost:8080/au` or `http://localhost:8080/us` to use the GraphQL playground client, or write queries using Postman/Insomnia etc (gql query endpoints are `http://localhost:8080/query/au` or `http://localhost:8080/query/us` and expect POST requests)

## Example queries

### Australia

Get all AU hubs

```
query getHubs {
  hubs {
    id
    country
    name
    address
    state
    priority
  }
}
```

Create new AU hub

```
mutation createHub {
  createHub(
    input: {
      country: "au"
      name: "Sydney"
      address: "123 Sydney St"
      state: "NSW"
      priority: 2
    }
  ) {
    id
    country
    name
    address
    state
    priority
  }
}
```

### United States

Get all US hubs (note that in this example US hubs do not use the priority field as a test)

```
query getHubs {
  hub {
    id
    country
    name
    address
    state
  }
}
```

Get single US hub

```
query getHub  {
  hub(id: 1) {
    id
    country
    name
    address
    state
  }
}
```

Create new US hub

```
mutation createHub {
  createHub(
    input: {
      country: "us"
      name: "New York"
      address: "123 New York St"
      state: "NY"
    }
  ) {
    id
    country
    name
    address
    state
  }
}
```

Update hub by id

```
mutation updateHub {
  updateHub(
    id: 1
    input: {
      country: "gb"
      name: "Manchester"
      address: "123 Kings road"
      state: "MN"
    }
  ) {
    id
    country
    name
    address
    state
  }
}
```

Delete hub by id

```
mutation deleteHub {
  deleteHub(
    id: 2
  )
}
```

Get schema

```
query getSchema {
  schema
}
```

## JSON schema format

General file structure

```
{
  menu: []          [MenuItem] - List of menu items (could contain subitems)
  ...*name*:        String - Model name for queries and mutations
    ModelSchema     ModelSchema - Schema of the model
}
```

MenuItem:

```
{
  model: "dc"     String - Model name for queries and mutations. If empty, page for menu item will not be created
  title: "DC"     String, required - Model title for displaying in menu
  single: false   Boolean, false - True if model is singleton, and has no table view (like settings, configs, etc)
  items: []       [MenuItem] - List of subitems
}
```

ModelSchema:

```
{
  id: "name"          String - Indetifier of the id field
  fields: []         [FieldSchema]
}
```

FieldSchema:

```
{
  "name": "name"                  String, required - Field name, for data accessing
  "type": "string"                String, required - Field type, for select proper component to display
  "description": "Id"             String, required - Field title
  "disabled": true                Boolean, false - If set, user can not edit this field
  "required": true                Boolean, false - If set, field value should exists
  "enum": []                      [Any] - List of possible values
  "sortable": false               Boolean, true - If set, field could be sortable in table view
  "cssClass": "font-weight-bold"  String - css class for decorating values in table view
}
```

## How to add new resources

Coming soon...

## Todo

- Expose Graphql schemas by country via endpoint to be consumed by client to generate forms using form generation library (uniform)
- Separate metadata resources on a brand based level
- Separate static and weekly metadata resources
