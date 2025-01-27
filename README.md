# N-Able-Task-App
## Rory Raeper

## How to Run
To run the app, you'll need `Go`, `node/npm`, and `docker` installed.
Once these are present, install the application in your Go directory (for example `go/src/`) as `github.com/RoryRaeper/n-able-task-app`
Then all you need to do is run:
```
setup_app.sh
```
This should host the backend service on docker alongside the database on the mongoDB image.
And then boot the react frontend via `node start`
