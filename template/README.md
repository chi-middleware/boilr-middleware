# [Chi](https://{{RepoHost}}/go-chi/chi) {{Name}} middleware

{{Description}}

[![Documentation](https://godoc.org/{{RepoHost}}/{{RepoOwner}}/{{RepoName}}?status.svg)](https://pkg.go.dev/{{RepoHost}}/{{RepoOwner}}/{{RepoName}})
[![codecov](https://codecov.io/gh/{{RepoOwner}}/{{RepoName}}/branch/master/graph/badge.svg)](https://codecov.io/gh/{{RepoOwner}}/{{RepoName}})
[![Go Report Card](https://goreportcard.com/badge/{{RepoHost}}/{{RepoOwner}}/{{RepoName}})](https://goreportcard.com/report/{{RepoHost}}/{{RepoOwner}}/{{RepoName}})
[![Build Status](https://cloud.drone.io/api/badges/{{RepoOwner}}/{{RepoName}}/status.svg?ref=refs/heads/master)](https://cloud.drone.io/{{RepoOwner}}/{{RepoName}})

## Usage

Import using:

```go
import "{{RepoHost}}/{{RepoOwner}}/{{RepoName}}"
```

Use middleware:

```go
    r := chi.NewRouter()
    r.Use({{RepoName}}.{{MiddlewareFuncName}})
```
