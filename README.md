# Gotversion

Gotversion is a small tool that examines your Git repositry and tries to come up with the most meaningful and consistent semantic version string. It's inspirted by [GitVersion](https://github.com/GitTools/GitVersion). If you are looking for a mature and super configurable solution you should probably look into it.

# Goals
The goal of Gotversion is to be lightweight, fast and friendly to CI pipelines usage. It currently output JSON and VSO renders of the calculated solution. The aim is to achieve this without any sort of configuration so it will be a bit opiniated on your usage of tags. If you use GitFlow or GitHubFlow it should fit perfectly.

# Building

The easiest way to buld Gotversion is using Docker

```
$ git clone https://github.com/loafoe/gotversion.git
$ cd gotversion
$ docker build -t gotversion .
```

# Running

```
$ docker run --rm -v /path/to/git/repository:/repo gotversion 
```

# License
License is MIT

