# Pushing a local repository to a newly created remote repository

## Step 1: Create a new repository on GitHub

## Step 2: Initialize the local directory as a Git repository

```bash
$ git init
```

## Step 3: Set upstream for the local repository

```bash
$ git remote add origin git@github.com:harshshekhar15/golang-dev.git
```

## Step 4: Set upstream branch

```bash
$ git push --set-upstream origin main
```

## Step 5: Rename the branch to main

```bash
$ git branch -m master main
```
