# Changelog

This project adheres to semantic versioning and all major changes will
be noted in this file.

## [0.2.0] 2023-11-01

- Hide Conn.ReadCloser and Conn.WriteCloser
- go mod tidy

## [0.1.3] 2023-02-07

- Adapt Close to match io.Closer

## [0.1.2] 2023-02-07

- Fix Close; closing one side closes the other

## [0.1.1] 2023-02-07
## [0.1.0] 2023-02-07

- Add func RandIPv4
- Add func Dial to create connections
- Add types Conn and Addr
