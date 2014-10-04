# gotodo

Minimal webapp for displaying and editing [todo.txt](http://todotxt.com/) files.

## Server

- Read in todo.txt file and convert to JSON.
- Expose URL for clients to fetch JSON representation of todo.txt.
- TODO: Allow POST requests for changes to be written back to todo.txt file (renderer JSON -> todo.txt format exists).


## Client

- Fetch and display todo.txt JSON data.
- TODO: Re-fetch in intervals.
- TODO: Allow in-browser editing and saving, i.e. sending JSON back to the server.

## Building and running

```
go build -o gotodo main.go

Usage of ./gotodo:
  -port=4242: Local port to listen on
  -todofile="todo.txt": todo.txt file to use
```

- NB: Server will also serve static files relative to executable's dir (for _index.html_ and _js/*.js_).
