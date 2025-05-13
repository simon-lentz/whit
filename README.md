# whit
"Whit" is a Swiss army-knife type utility application for various Wyrth Research commands such as:

* Initialization of db
* Running scrapers/data-collectors
* Processing data
* Updating db with new data

Note: At this time, this repo is under transition from being an instantiation of the "goapp-template" and until the transition
is complete there will be some "hello world" and example boilerplate code present.

## How to run

With the root of the repo as the current directory:
```
go run .
```
This will show the built-in help over available commands and options.
## How to run tests

With the root of the repo as the current directory:
```
go test ./... -count=1 -v
```

This will run all tests in the repo after having forced a new build (the option `-count=1`) with the output of the outcome of each test (the option `-v` verbose).

## Working with this repository
### Configure git and GitHub

* You must have a GitHub account
* Configure your GitHub account (see instructions at GitHub)
  * so you can use SSH (generate a key and upload to GitHub)
  * to use "GitHub anonymous email" (select this option in your account and copy the email you got)
  * add one or more of your valid email addresses to your account and select the one to be your primary email
* Configure your local GitHub environment
  * to use the GitHub anonymous email for your commits
  * also set your name (full, real name, or GitHub moniker)

For example: (obviously replace "Henrik" specifics with your name/email):
```
git config --global user.email 563066+hlindberg@users.noreply.github.com
git config --global user.name "Henrik Lindberg"
```

Verify by viewing your settings:
```
git config --get-regexp user.*
``` 

Note: The name and email can instead be set at the repository level if you don't want it for all your git repos (see git docs for how to do that). If you want this, you need to set it up in your clone after you completed the steps below.


### Fork, clone, and configure the repository

* At GitHub, fork the repository to create your personal fork. All your work should be pushed to your fork at GitHub and you incorporate it into the main repo using a Pull Request (PR) at GitHub. We do this to reduce the number of branches in the main repo and to avoid mistakes. It is your job to do keep your fork as tidy as you want it to be.
* When you have forked, you clone the repo to your local development machine and configure the repository. You want your clone in the location where go code is placed on your machine, as this makes it easier to work with. You should make this clone with the git command line git and not in your IDE since certain IDEs will configure the repo with settings we don't want!
```
cd ~/go/src/github.com/wyrth-io
git@github.com:wyrth-io/whit.git
```
* You now have the cloned repository and can configure it to allow you to work with both the main repo and your fork at GitHub.
  Note: If you already cloned your fork to your local machine it will have an `origin` set that refers to your fork and not the main repo. **If you did this, delete the repo and clone the main repo as shown above**. That, and the steps below ensures we all work the same way. In the instructions below uou must replace `hlindberg` with your GitHub id!<br>
  These settings ensure that the `main` branch is always a reference to the `origin` repo, that merges to it go to the expected location in your file system, and that if you accidentally commited to your `origin main` branch it will replay your commits with a `rebase` after having fetched the official remote branch rather than doing a merge (which is harder to unwind).
```
cd ~/go/src/github.com/wyrth-io/whit
git remote add hlindberg git@github.com:hlindberg/whit.git
git config branch.main.remote origin
git config branch.main.merge refs/heads/main
git config branch.main.rebase true
```

You can now test pulling the remote repos - this will fetch all changes and apply to the local branches:
```
pull origin
pull hlindberg
```

You are now up to date and can start working on the repo.

## Working with the repo
* checkout out `origin main` branch and pull it (to always start with the latest)
* create a branch to work on and name the branch strting with a ticket number, or if you don't have a ticket number with a keyword such as `maint` or `docs` for code maintenance or docs/typos fixes. Always use a short sentence with just enough information from the ticket to make it identifiable. For example `123-update-version`, `maint-update-foolib-to-latest`. Do not mix in your username since this will be apparent since branches are coming from your fork.
* The first time you push a branch to a repo it will be bound to that repo. IDE's typically help with this task, but make sure it does it the right way - you don't want to push your work branches to `origin`. On the command line you can do this to start working on a branch to fix an issue:
```
git checkout main
git pull
git checkout -b 666-fix-leakage-problem
git push -u hlindberg 666-fix-leakage-problem
```
* Now, you can make changes to the code and commit. When making a commit you *must* make a commit comment on the following format:
  * Always start with ticket number (or keyword) in parentheses, and then an explanation what the commit does using imperative form for the first line. For example `(666) Fix leak in foo.go`. Think of the statement as the continuation of the phrase `When this commit is applied it will...`. Then (on additional lines) continue with a longer explanation *why* the change is made - this as opposed to trying to describe how source was changed as that is apparent from the code diff. You are writing this comment for your collegues and for the "future you" and they are important as they document how the logic came to be the way it is.
* Break up your changes into smaller commits - if you for example find typos as you are reading code, do not mix in those changes with changes to logic (unless possibly if they are for the lines of logic you are working on). Later when you are reviewing other peoples code you will know why this is of great value.
* Learn how to use `git rebase -i` for interactively rebasing your commits. This allows you to squash and split up commits, reorder them, etc. etc. and it is an invaluable tool to make sure the history of the code is clean and understandable.
* Try to make changes that you "have to geet out of the way" first - for example, you see porly formatted code with typos, address those first before you work on that file. You know that you will need an extra parameter in a function for your logic and as a consequence you will need to change the calls throughout the code base - make this a separate commit and then add your new logic. If you forgot or got carried away, `rebase -i` is your friend.


When you are done with your commits, push the branch to your fork at github. Actually it is a good habit to do this now and then to make sure you have a backup of your work - especially when you are new to working with git. For example, if you are going to do a `git rebase -i`, do a push first so you can retrieve your work from github if you screwed up your interactive rebase.

If you did the push as showed above your subsequent pushes are simple when you have the branch checked out simply type:
```
git push -u hlindberg HEAD
```
which is a short form of:
```
git push -u hlindberg 666-fix-leakage-problem
```
where you have to spell out the branch name.

Your IDE helps a lot with this, just remember to do the initial push to the correct repo.

### Code Quality
In order for code to be merged it must pass all tests and all configured linting. The linters check both hard and
stylistic issues. Any reported problem must be fixed before merging.

**Setup**</br>
* We build with latest go (currently 1.91) - if you do not have that installed follow the steps [here](https://go.dev/dl)
* We run lint with `golangci-lint` and you should install it locally so you can check your code. Follow instructions [here](https://golangci-lint.run/welcome/install/)

**Continously, or at least before PR**</br>
* Run `go test ./... -count=1 -v` and make sure all tests are green / ok.
* Run `golangci-lint run` which runs linting on the entire project, make sure there are no warning or errors
