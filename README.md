# clerk

`clerk` is a very simple declarative development framework.  
The major version of `clerk` has not yet reached v1 and is still in the verification phase.  
Also, it is still unknown how `clerk` will solve development problems.  
However, `clerk`'s approach to development, declarative development,   
has the potential to make source code more maintainable in that it is implemented while writing specifications.

## feature

- Declarative Development
- Support for multiple dynamically typed languages (currently only python)

## installation

**for Mac OS**  

```
brew tap asweed888/homebrew-clerk
brew install clerk
```

**for Linux**

```
sudo curl -L https://github.com/asweed888/clerk/releases/download/{Any version}/clerk_linux_x86_64.tar.gz -o - | tar -xzvf - && sudo mv ./clerk /bin
```

**other**  

It can be installed from the release page.  
https://github.com/asweed888/clerk/releases

## Usage
The use of **clerk** is very simple.
All you need to do is create a file called **clerk.yml** and describe a simple structure.
