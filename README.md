# cozy-inn

## Setting Environment Variable

Set the GOOGLE_APPLICATION_CREDENTIALS environment variable to the path of the JSON file containing your service account key. This variable only applies to your current shell session. So if you open a new session, set the variable back

For windows
```sh
$env:GOOGLE_APPLICATION_CREDENTIALS=".firebase\service-account.json"
```

For Linux or macOS
```sh
export GOOGLE_APPLICATION_CREDENTIALS=".firebase\service-account.json"
```