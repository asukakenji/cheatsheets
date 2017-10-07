# Handling tags with Git commands

### Creating an annotated tag:

```
git tag -a v1.4 -m "my version 1.4"
```

An annoated tag is stored as a full object in the Git database.

### Creating a lightweight tag:

```
git tag v1.4
```

A lightweight tag is just a pointer to a specific commit.

### Sharing a tag:

```
git push origin v1.4
```

### Sharing all tags:

```
git push origin --tags
```

### Listing all tags:

```
git tag
```

### Listing all tags with timestamp

```
git for-each-ref --shell --format="ref=%(refname:short) dt=%(taggerdate:format:%s)" refs/tags/*
```

### Deleting a tag:

```
git tag -d v1.4
```
