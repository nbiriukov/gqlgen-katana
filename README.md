# GQLGen Katana v2

This repo is a PoC for using an SDL-first GraphQL server using the gqlgen library for Go.

## How to test

Start server:

```
$ make run_server
```

Visit either `http://localhost:8080/au` or `http://localhost:8080/us` to use the GraphQL playground client, or write queries using Postman/Insomnia etc (gql query endpoints are `http://localhost:8080/au/query` or `http://localhost:8080/us/query` and expect POST requests)

## Example queries

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

Get all US hubs (note that in this example US hubs do not use the priority field as a test)
```
query getHubs {
  hubs {
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

## How to add new resources

Coming soon...

## Todo

- Expose Graphql schemas by country via endpoint to be consumed by client to generate forms using form generation library (uniform)
- Separate metadata resources on a brand based level
- Separate static and weekly metadata resources