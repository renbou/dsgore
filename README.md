# dsgore
Ever wanted those annoying .DS_Store files to just be gone in an instant? Or maybe you wanted to be able to forget about them and never see them again in your hard-worked-for repositories, commits and diffs? Always .DS_Store with your whole soul?

Well now you can achieve all that and more! Bringing you:  

<img src="design.svg" width="800" height="200">

# Usage
You may be wondering *Well how do I get this amazing tool to be part of my arsenal?* - and the answer, my friend, is damn simple it is!  
Here are **multiple** ways of launching it if you want to:

Install it for multiple usages with
  - Go:
    ```bash
    go get github.com/renbou/dsgore
    # Now launch it!
    dsgore
    ```
  - curl:
    ```bash
    curl -s https://renbou.ru/static/get-dsgore.sh | bash
    ```

Launch it once with
  - Go:
    ```bash
    go run github.com/renbou/dsgore
    ```
  - curl:
    ```bash
    # If you want to cleanse the current directory
    curl -s https://renbou.ru/static/run-dsgore.sh | bash
    # If you want to clear some other directory or pass other args, then run it like this
    curl -s https://renbou.ru/static/run-dsgore.sh | bash -s - -d {directory}
    ```

## Removing .DS_Store files
Simply spin this bad boy right up by running `dsgore` in your terminal, which will remove all the .DS_Store files in the directory tree beginning in the current directory.  
Now if you want to point it to a specific directory - just use the `-d/--directory` flag like so:
```bash
dsgore --directory ~
```

# Roadmap
- Automatically write .DS_Store into the current repo's gitignore
- Automatically setup a git hook which deletes .DS_Store before any pushes
- Create a github action which launches this tool (e.g. to clean up pushes or pull requests)
- Setup autorelease action
