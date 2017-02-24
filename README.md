Go Router
====================================

# Example

## Params

```
type Param struct {
	Name string `schema:"name"`
	Hoge []string `schema:"hoge"`
	School School `schema:"school"`
}

type School struct {
	Name string `schema:"name"`
}
```

## Exec

```
$ curl http://localhost:8080/?hoge=a&hoge=2&name=shinohara&school.name=mf
main.Param{
    Name:   "",
    Hoge:   {"a"},
    School: main.School{},
}
```

