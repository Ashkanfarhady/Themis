```
            ,ggg,                   gg                   ,ggg,
           d8P""8b                ,d88b,                d8""Y8b
           Y8b,__,,aadd88888bbaaa,888888,aaadd88888bbaa,,__,d8P
            "88888888888888888888I888888I88888888888888888888"
            /|\`""YY8888888PP""""`888888'""""YY8888888PP""'/|\
           / | \                  `WWWW'                  / | \
          /  |  \                 ,dMMb,                 /  |  \
         /   |   \                I8888I                /   |   \
        /    |    \               `Y88P'               /    |    \
       /     |     \               `YP'               /     |     \
      /      |      \               88               /      |      \
     /       |       \             i88i             /       |       \
    /        |        \            8888            /        |        \
"Y88888888888888888888888P"       i8888i       "Y88888888888888888888888P"
  `""Y888888888888888P""'        ,888888,        `""Y888888888888888P""'
                                 I888888I
                                 Y888888P
                                 `Y8888P'
                                  `WWWW'
                                   dMMb
                                _,ad8888ba,_
                    __,,aaaadd888888888888888bbaaaa,,__
                  d8888888888888888888888888888888888888b
```

[Themis](https://en.wikipedia.org/wiki/Themis) is the name of a load balancer, written in GoLang. 
It will forward the requests randomly to your backends

You have to create a config file in yaml format and put your backends addresses there. (Also there is an example in the repository)
And the address of that config file, shoud be export to environment as a variable.

You have to provide your desired port as an argument

Example:
```
./main 8585
```
Feature works:
Right now I have used the simplest algorithm for balancing the load. If yuo wish to contribute, you can implement other algorithms too and make me happy :D

