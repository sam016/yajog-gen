# yajog-gen

This repository is created just for testing the gojay generator for `map[string]interface{}`

## Steps

Install gojay

```bash
go install github.com/francoispqt/gojay/gojay
```

Get this test repository

```bash
go get -u github.com/sam016/yajog-gen
```

Test code generation

```bash
gojay -s dictionary.go -p true -t PlainDictionary,UnMarDictionary

gojay -s dictionary.go -p true -t PlainClassroom,UnMarClassroom
```
