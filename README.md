# hspgo
HSP API wrapper for Go

Documentation
----------------
https://wiki.openraildata.com/index.php/HSP

http://godoc.org/github.com/sb89/hspgo

Install
----------------
```
go get github.com/sb89/hspgo

```
Example
----------------
```
c := NewClient("test@test.com", "password")

resp, err := c.GetServiceDetails("123456789")

...
```