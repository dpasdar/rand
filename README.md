# Rand Fake Process
As an example of long running process, this golang program will behave randomly like a normal process with the following list of behavior:

 - It will run between 10 to 40 seconds randomly
 - It will sometimes crash and produce errors randomly
 - It can't be terminated instantly, it needs at least 3 seconds to do cleanup, so upon receiving SIGTERM it will wait 3 seconds before dying 

 
 To compile simply install go and run the following command inside the directory
 `go install`
 
 The binary will be located in default `$GOPATH/bin`
 
 The program is intended to be used with azrael process manager as a usecase to test different scenarios:
 [Azreal Process Manager](https://github.com/dpasdar/azrael)