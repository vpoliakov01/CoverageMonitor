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
![image](https://user-images.githubusercontent.com/53489500/84440266-85963000-abee-11ea-95e3-f7d88f1fcc63.png)

### Select a file
![image](https://user-images.githubusercontent.com/53489500/84440334-a65e8580-abee-11ea-96b4-547f3d8e3252.png)

### Click Run Tests
![image](https://user-images.githubusercontent.com/53489500/84440367-b8d8bf00-abee-11ea-8360-c3c7d0cff235.png)

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
