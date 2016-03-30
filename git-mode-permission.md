# Handling modes (permissions) with Git commands

### Listing Git directory at `origin` with modes (permissions):

```
git ls-tree origin
```

### Listing Git directory at `HEAD` with modes (permissions):

```
git ls-tree HEAD
```

### Listing Git staged objects with modes (permissions):

```
git ls-files --stage
```

### Enabling the execute permissions on the updated files:

```
git update-index --chmod=+x *.sh
```

### Disabling the execute permissions on the updated files:

```
git update-index --chmod=-x *.txt
```
