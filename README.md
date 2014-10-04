# gotodo

Minimal webapp for displaying and editing [todo.txt](http://todotxt.com/) files.

## Server

- Read in todo.txt file and convert to JSON.
- Expose URL for clients to fetch JSON representation of todo.txt.
- TODO: Allow POST requests for changes to be written back to todo.txt file (renderer JSON -> todo.txt format exists).


## Client

- TODO: Fetch and display todo.txt JSON data in intervals.
- TODO: Allow in-browser editing and saving, i.e. sending JSON back to the server.

## Building

```
go build -o gotodo main.go
```
