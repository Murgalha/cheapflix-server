# Cheapflix Server

## About
Cheapflix is a piece of software that aims to transform your movie collection
into an online platform to watch them.

This repo contains the files to run the server side of Cheapflix, responsible for creating
and serving a RESTful API from your collection. The client can be found
[here](https://github.com/Murgalha/cheapflix-client).

## Movie
The server is responsible for reading a movie directory and providing
a RESTful API with the movies' data and their information. The server
expects the movies folder to have the following structure:
```
<root>/
├─ <movie-1>/
│  ├─ <movie-1 img>
│  ├─ <movie-1 movie>
│  ├─ <movie-1 subtitle>
├─ <movie-2>/
│  ├─ <movie-2 img>
│  ├─ <movie-2 movie>
│  ├─ <movie-2 subtitle>
├─ <movie-3>/
│  ├─ <movie-3 img>
│  ├─ <movie-3 movie>
│  ├─ <movie-3 subtitle>
```
This way, the server will be able to provide the following JSON for each movie:
```js
{
    "id": 1163,
    "name": "Movie Name",
    "year": 1997,
    "subtitle": "<server-url>:3000/movie/1163/sub",
    "url": "<server-url>:3000/movie/1163/movie",
    "poster": "<server-url>:3000/movie/1163/poster"
}
```

### Endpoints
The following endpoints are available to consume the movie data:
```
/movie/all -> Get data from all movies
/movie/:id -> Get data from movie with id ":id"
/movie/:id/sub -> Get the subtitle from movie with id ":id"
/movie/:id/movie -> Get the movie itself from movie with id ":id"
/movie/:id/poster -> Get the poster from movie with id ":id"
```

## Running
Clone this repo, install cheapflix dependencies and run it with
the following commands:
```sh
git clone https://github.com/Murgalha/cheapflix-server ~/cheapflix-server
cd ~/cheapflix-server
go get ./...
go run src/*.go <movie-dir>
```
