# Serverless - Azure Function with GO

This repo is just to learn the Serverless GO with Azure

## Guide - VSCode

1. Install Azure Function extension
2. Azure tab -> Workspace -> Create Project -> Create Function
3. Select a folder, usually your current folder.
4. In Select a language, select Custom Handler.
5. In Select a template for your first function, select HttpTrigger.
6. Give the app a name, such as hello.
7. Select an authorization level of anonymous. You can change that later if you want.

## Run the app

```go
go build server.go
```

Then update `defaultExecutablePath`
Then start the function

```go
func start
```
