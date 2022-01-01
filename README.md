# Mail content parsing
邮件内容解析

## Simple usage

You just parse a io.Reader that holds the email data. The returned Email struct contains all the standard email information/headers  as public fields.

```go
var reader io.Reader // this reads an email message
email, err := mcopa.Parse(reader) // returns Email struct and error
if err != nil {
    // handle error
}

fmt.Println(email.Subject)
fmt.Println(email.From)
fmt.Println(email.To)
fmt.Println(email.HTMLBody)
```

## Retrieving attachments

Attachments are a easily accessible as `Attachment` type, containing their mime type, filename and data stream.

```go
var reader io.Reader
email, err := mcopa.Parse(reader)
if err != nil {
    // handle error
}

for _, a := range(email.Attachments) {
    fmt.Println(a.Filename)
    fmt.Println(a.ContentType)
    //and read a.Data
}
```

## Retrieving embedded files

You can access embedded files in the same way you can access attachments. They contain the mime type, data stream and content id that is used to reference them through the email.

```go
var reader io.Reader
email, err := mcopa.Parse(reader)
if err != nil {
    // handle error
}

for _, a := range(email.EmbeddedFiles) {
    fmt.Println(a.CID)
    fmt.Println(a.ContentType)
    //and read a.Data
}
```


# 参考
- https://github.com/DusanKasan/Parsemail
- https://github.com/emersion/go-imap
