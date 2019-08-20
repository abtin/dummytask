## dummytask simulates a CPU heavy task.

It can be used to simulate high CPU loads when testing compute grids.

To build:
  -  On Mac/Linux/Windows:

         $ go build
  - On Mac/Windows for linux:
  
        $ make linux
  - docker
    
        $ make docker
        
To execute:
  -  With default values:
    
         $ ./dummytask 
    
  - Print Usage:
  
        $ ./dummytask -h
