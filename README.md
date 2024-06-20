# 28540

Reproduction for [Renovate issue 28540](https://github.com/renovatebot/renovate/issues/28540). 
Renovate bot is self-hosted, version 37.381.6.

## Current behavior

Renovate can't get module from subgroup repository

## Expected behavior

If try to directly call the same url - you also get 404, but if you add after sub-group the repository name - you'll get the correct page:
```
https://gitlab.com/api/v4/projects/test-nested%2Fservices%2Fmodule-test/repository/tags?per_page=100
```

## Link to the Renovate issue or Discussion

https://github.com/renovatebot/renovate/issues/28540
