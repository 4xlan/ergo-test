## Project: "demo"

### Generated with
 - Types for the network messaging: true
 - Enabled Observer (http://localhost:9911): false

### Supervision Tree

Applications
 - `testApp{}` demo/apps/testapp/testapp.go
   - `testSup{}` demo/apps/testapp/testsup.go
     - `Act1{}` demo/apps/testapp/act1.go
     - `Act2{}` demo/apps/testapp/act2.go
     - `Act3{}` demo/apps/testapp/act3.go

Messages are generated for the networking in demo/types.go
- `Req{}`
- `Ans{}`


#### Used command

This project has been generated with the `ergo` tool. To install this tool, use the following command:

`$ go install ergo.services/tools/ergo@latest`

Below the command that was used to generate this project:

```$ ergo -path ./ -init demo -with-app testApp -with-sup testApp:testSup -with-actor testSup:Act1 -with-actor testSup:Act2 -with-actor testSup:Act3 -with-msg Req -with-msg Ans ```
