# Moving files from a Git repository to a new one

### Create a temp directory for working

```
mkdir temp
cd temp/
```

### Clone the repository for working

```
git clone https://github.com/asukakenji/cheatsheets.git
cd cheatsheets/
```

### Unlink the clone from GitHub to prevent affecting it by accident
```
git remote remove origin
```

### Filter the files (retain only `clash-royale*` here, any shell command will do)

```
git filter-branch --tree-filter 'for i in *; do case "$i" in clash-royale*) true; ;; *) rm "$i"; ;; esac; done' HEAD
```

### Prune the logs (otherwise unrelated logs will be retained)

```
git filter-branch --prune-empty -f
```

### Initialize a new repositry

```
cd ../
mkdir clash-royale
cd clash-royale/
git init
```

### Pull the files from the tidied up repository

```
git pull ../cheatsheets/
```

### Link the repository with GitHub (need to create a repository on GitHub first)

```
git remote add origin https://github.com/asukakenji/clash-royale.git
```

### Push the contents to GitHub

```
git push -u origin master
```

### Clean up the temp directory

```
cd ../..
rm -fr temp/
```
