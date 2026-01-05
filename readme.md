    1. In Go we have multiple packages. If the package is main it must have a main function as well and it’s a executable package where as if there is any other name it will be the reuseable package. 
    2. We can do go run, go build  and many others like go fmt (format code), test , install, get  etc. 
    3. Multiple files with same package name will be considered as that one package one can have multiple files all starting with package main

    4. We can use all the inbuilt packages or the packages created by the common people.

    5. A function can have multiple return values.

    6. using the os.WriteFile and ReadFile function we can write in to the file as a [] byte and its a ascii form of the string and get it back the same way.  It is not able to read we can also the error handling for the same.

    7.  At the same place we can provide the error handling options as well. like the exit codes and the status of the exit code can be non zero for handling the errors and thus can be made sure that its working as expected.

    8. The os.exit () fuction is supposed to get the status code 0 only for the pass in case of failed its expected to give the status code of 1-n

    9. No inbuilt package to shuffle the deck.

    10. Go does not proivide data like the 10 test cases were executed it just gives the name that the test case passed it only assumes that something  terriblke is not failing inside our function.

    STRUCTS

    Its kind of a dictionary in python and its there for the golang to make sure we can add a dataset that has like 2 strings related.
    so we have a struct that can have a person's name first name and last name both strings and that is a struct.

    If the variable is defined but no value defined it automatically deifnes the zero values like string has "" int has 0 etc. 

    A struct can also have a struct inside embedding of struct is allowed and its helpful to make a very detailed and complex data types

    





