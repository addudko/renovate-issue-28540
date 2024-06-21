# 28540

Reproduction for [Renovate issue 28540](https://github.com/renovatebot/renovate/issues/28540). 
Renovate bot is self-hosted, version 37.381.6.

## Current behavior

Renovate can't get submodule from gitlab repository
```bash
DEBUG: Go lookup source url https://gitlab.com/test-nested/services/services-subgroup/module-test for module gitlab.com/test-nested/services/services-subgroup/module-test/api (repository=addudko/renovate-issue-28540)
DEBUG: Go lookup source url https://gitlab.com/test-nested/services/services-subgroup/module-test for module gitlab.com/test-nested/services/services-subgroup/module-test (repository=addudko/renovate-issue-28540)
DEBUG: GET https://gitlab.com/api/v4/projects/test-nested%2Fservices%2Fservices-subgroup%2Fmodule-test%2Fapi/repository/tags?per_page=100 = (code=ERR_NON_2XX_3XX_RESPONSE, statusCode=404 retryCount=0, duration=337) (repository=addudko/renovate-issue-28540)
DEBUG: GitLab API 404 (repository=addudko/renovate-issue-28540)
       "url": "https://gitlab.com/api/v4/projects/test-nested%2Fservices%2Fservices-subgroup%2Fmodule-test%2Fapi/repository/tags?per_page=100"
DEBUG: Datasource 404 (repository=addudko/renovate-issue-28540)
       "datasource": "go",
       "packageName": "gitlab.com/test-nested/services/services-subgroup/module-test/api",
       "url": "https://gitlab.com/api/v4/projects/test-nested%2Fservices%2Fservices-subgroup%2Fmodule-test%2Fapi/repository/tags?per_page=100"
DEBUG: Failed to look up go package gitlab.com/test-nested/services/services-subgroup/module-test/api (repository=addudko/renovate-issue-28540, packageFile=go.mod, dependency=gitlab.com/test-nested/services/services-subgroup/module-test/api)
```

## Expected behavior

Renovate checks repository submodule correctly.

## Link to the Renovate issue or Discussion

https://github.com/renovatebot/renovate/issues/28540
