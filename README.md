# go-skeleton

Base template that will be used for developing go web apps

## Project Structure

<pre>
  |── app            // Application main folder
  │   ├── action     // Application Main Controller
  │   │   └── api    // Application API Controller
  │   ├── entity     // Application Entity
  │   ├── repo       // Application DAO
  │   └── service    // Application Business Logic
  ├── cmd            // Application commandline command
  ├── conf           // Application Main Configuration
  └── util           // General Function
</pre>

## Build project

1. Create a copy of *.sample file and rename the newly copied to without *.sample
2. run `go build`

## Run project

1. run `./go-skeleton serve`

## Logging

currently log file is created on project root with name `skeleton.go`

## Migration 

Migration will have 3 command to migrate database `--up`, `--down` `--step <number>`
* --up will migrate all database to the latest version
* --down will migrate down database to the very first one **be careful**
* --step <n>, if n > 0 then it will migrate up by n number, if n < 0 then it will migrate down by n number
