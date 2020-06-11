# CoverageMonitor
CoverageMonitor displays a Go repository and its test coverage

## Prereqs:
* docker
* make

## To run:
* `make docker` (due to node-modules, might take 5-10min)
* go to `localhost:8321` (preferably with chrome)

## What it looks like:

### Select a github repository
![image](https://user-images.githubusercontent.com/53489500/83067744-0f090800-a01c-11ea-80fb-642305b5a7fa.png)

### Select a file
![image](https://user-images.githubusercontent.com/53489500/83067826-2d6f0380-a01c-11ea-960e-358df5d079ed.png)

### Click Run Tests
![image](https://user-images.githubusercontent.com/53489500/83067907-4a0b3b80-a01c-11ea-85fa-b3a3f3f836c4.png)

## What's happening on the back end:
When the user enters the repo, the back end uses github APIs for
* Meta info (number of watchers and repo language)
* Repo Files
* File content

collects the needed information and replies with the acquired data to the front end (api calls are run in parallel, so the response is almost immediate).

When the user clicks `RUN TESTS`, the back end clones the repo, runs `go test -coverprofile` on it, parses the results and replies to the front end (takes ~1 sec for `/google/uuid` due to cloning and testing).

## Technologies Used
* React with TypeScript (tsx) + CSS
* Go
* Nginx
* Docker
