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

To create or update a Task, please use `cURL` or `Postman` (I didn't have time to implement the front end handling for this).

This can be done with the following:
```
curl -H 'Content-Type: application/json' \
      -d '{ "title":"Task Title","description":"Task Description", "status": "todo"}' \
      -X POST \
      https://localhost:8080/tasks
```
or for Postman:
```
POST localhost:8080/tasks
body:
{
  "title": "Task Title",
  "description": "Task Description",
  "status": "todo"
}
```
You can replace `POST` with `PUT` for update. Make sure to supply the ID in the URL like so: 
```
PUT localhost:8080/tasks/507f1f77bcf86cd799439011
```
