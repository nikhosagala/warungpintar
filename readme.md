# Warung Pintar Test

This is a test for Warung Pintar hiring process.

## Requirements

* You need [Go](https://golang.org/) installed in your computer

## Build & running

```shell script
    $ go build
    $ ./warungpintar
    //or on windows
    $ warungpintar.exe
```

## Available path

* GET **/messages** -> getting all messages
* POST **/messages** -> add new messages
    
    request payload:
    ```json
      {
          "body": "hello world"
      }
    ```