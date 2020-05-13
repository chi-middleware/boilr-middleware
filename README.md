# boilr-middleware

This is a [boilr template](https://github.com/tmrts/boilr) for creating a middleware for Chi. Get started by installing the template:

```console
$ boilr template download chi-middleware/boilr-middleware chi-middleware -f
```

Create a project in directory sample:

```console
$ boilr template use chi-middleware sample
```

Enter year:

```text
[?] Please choose a value for "Year" [default: "2020"]:
```

Enter author:

```text
[?] Please choose a value for "Author" [default: "Lauris BH"]:
```

Enter the git repository host:

```text
[?] Please choose a value for "RepoHost" [default: "github.com"]:
```

Enter the git repository owner:

```text
[?] Please choose a value for "RepoOwner" [default: "chi-middleware"]:
```

Enter the git repository name:

```text
[?] Please choose a value for "RepoName" [default: ""]: sample
```

Enter middleware name:

```text
[?] Please choose a value for "Name" [default: ""]: sample
```

Enter middleware description:

```text
[?] Please choose a value for "Description" [default: ""]: my middleware description
```

Enter middleware function name that will be exported:

```text
[?] Please choose a value for "MiddlewareFuncName" [default: ""]: MyMiddleware
```

## Example

```console
$ boilr template use chi-middleware sample
[✔] Created .drone.yml
[✔] Created .gitignore
[✔] Created .golangci.yml
[✔] Created .revive.toml
[?] Please choose a value for "Year" [default: "2020"]:
[?] Please choose a value for "Author" [default: "Lauris BH"]:
[✔] Created LICENSE
[✔] Created Makefile
[?] Please choose a value for "RepoHost" [default: "github.com"]:
[?] Please choose a value for "Name" [default: ""]: sample
[?] Please choose a value for "Description" [default: ""]: Sample middleware description.
[?] Please choose a value for "RepoOwner" [default: "chi-middleware"]:
[?] Please choose a value for "RepoName" [default: ""]: sample
[?] Please choose a value for "MiddlewareFuncName" [default: ""]: Sample
[✔] Created README.md
[✔] Created go.mod
[✔] Created middleware.go
[✔] Created middleware_test.go
[✔] Successfully executed the project template chi-middleware in /home/user/go-projects/chi-middleware/sample
```
