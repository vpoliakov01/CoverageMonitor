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
![image](https://user-images.githubusercontent.com/53489500/84440830-8aa7af00-abef-11ea-90bc-8975f90b6997.png)

### Select a file
![image](https://user-images.githubusercontent.com/53489500/84440503-f63d4c80-abee-11ea-94fa-8898b0e7ac8c.png)

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
