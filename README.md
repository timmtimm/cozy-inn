# cozy-inn

## Setting Environment Variable

### Firebase

Set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to the path of the JSON file containing your service account key. This variable only applies to your current shell session. So if you open a new session, set the variable back

For windows
```sh
$env:GOOGLE_APPLICATION_CREDENTIALS=".firebase\service-account.json"
```

For Linux or macOS
```sh
export GOOGLE_APPLICATION_CREDENTIALS=".firebase\service-account.json"
```

### Environment

Copy env example and set every environment variable

## Live Reload

Live Reload being used in this application is [Air](https://github.com/cosmtrek/air).

Steps to use live reload for local development:
1. Install Air
```sh
go install github.com/cosmtrek/air@latest
```

2. Run live reload
```sh
air
```


1. go install github.com/vektra/mockery/v2@latest
2. cd to folder
3. mockery --all --keeptree

go test . -v -coverprofile=cover; go tool cover -html=cover