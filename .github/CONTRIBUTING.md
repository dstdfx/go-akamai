# Contribution guide

## Basics

Prior to create an issue that describes a problem/feature request or whatever 
before sending a pull-request.

## Project structure

Every Akamai API is implemented in it's separate package.

PAPI (https://developer.akamai.com/api/core_features/property_manager/v1.html) example:
https://github.com/dstdfx/go-akamai/tree/master/akamai/papi

Any package that implements methods to work with a needed API uses the
following structure:

```
api_object/              # Name of the directory should desrcibe API object (property/cpcode/activation)
├── doc.go               # Documentation that will be available at the godoc.org
├── requests.go          # Methods to work with the API
├── requests_opts.go     # Structures that contain options for a POST/PATCH calls
├── schemas.go           # Structures that contain unmarshalled responses.
└── testing
    ├── fixtures.go      # Tests fixtures.
    └── requests_test.go # Tests for all implemented requests.
```

## Test

Please implement tests for all methods that you're creating.
Check for examples: https://github.com/dstdfx/go-akamai/tree/master/akamai/papi/v1/property/testing

To run linters:
```bash
make golangci-lint
```

To run unit tests:
```bash
make test
```
