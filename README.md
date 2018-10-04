# GotVersion
GotVersion is a small tool that examines your Git repositry and tries to come up with the most meaningful and consistent semantic version string. It's inspired by [GitVersion](https://github.com/GitTools/GitVersion). If you are looking for a mature and super configurable solution you should definitely look into it.

# Goals
The goal of GotVersion is to be lightweight, fast and friendly to CI pipelines usage. It currently output JSON and VSO renders of the calculated solution. There is currently no configuration which means it's a bit opiniated on your usage of Git tags.
If you use GitFlow or GitHubFlow it should fit perfectly. 

# Building

The easiest way to buld GotVersion is using Docker

```
$ git clone https://github.com/loafoe/gotversion.git
$ cd gotversion
$ docker build -t gotversion .
```

# Running

The Docker image will examine the Git repo mounted on `/repo`:

```
$ docker run --rm -v /path/to/git/repository:/repo gotversion 
```

By default it outputs VSO lines:

```
##vso[task.setvariable variable=GotSemVer;isOutput=true;]0.4.0
##vso[task.setvariable variable=GotFullSemVer;isOutput=true;]0.4.0
##vso[task.setvariable variable=GotSHA;isOutput=true;]0de2be93db13828f1d5f43896406f4902ab4159c
##vso[task.setvariable variable=GotCommitDate;isOutput=true;]2018-10-04
##vso[task.setvariable variable=GotBranchName;isOutput=true;]master
```

## Azure DevOPS

These lines, when executed as a Docker run task in an Azure DevOPS (VSTS) pipeline will inject the following variable for use in downstream tasks:

| Variable              | ENV name              | Description                           |
|-----------------------|-----------------------|---------------------------------------|
| `$(Docker.GotSemVer)`     | `$DOCKER_GOTSEMVER`     | Abbreviated semantic version of the current build |
| `$(Docker.GotFullSemVer)` | `$DOCKER_GOTFULLSEMVER` | The full semantic version (recommended) | 
| `$(Docker.GotSHA)`        | `$DOCKER_GOTSHA`        | The full SHA1 of the branch head |
| `$(Docker.GotCommitDate)` | `$DOCKER_COMMITDATE`    | The date of the commit in `YYYY-MM-DD` format |
| `$(Docker.GotBranchName)` | `$DOCKER_GOTBRANCHNAME` | The Git branch name |


You can also have it emit JSON:

```
$ docker run --rm -v /path/to/git/repository:/repo gotversion -json
```

Will output something like:

```json
{
  "Major": 0,
  "Minor": 4,
  "Patch": 0,
  "SemVer": "0.4.0",
  "FullSemVer": "0.4.0",
  "MajorMinorPatch": "0.4.0",
  "BranchName": "master",
  "Sha": "0de2be93db13828f1d5f43896406f4902ab4159c",
  "CommitDate": "2018-10-04",
  "FullBuildMetaData": "Branch.master.Sha.0de2be93db13828f1d5f43896406f4902ab4159c"
}
```

# License
License is MIT

