## dummytask simulates a CPU heavy task.

It can be used to simulate high CPU loads when testing compute grids.

To build:
  -  Directly:

         $ go build
  - Via make for linux:
  
        $ make linux
  - docker
    
        $ make docker
        
To execute:
  -  With default values:
    
         $ ./dummytask 
    
  - Print Usage:
  
        $ ./dummytask -h
