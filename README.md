# 28540

Reproduction for [Renovate issue 28540](https://github.com/renovatebot/renovate/issues/28540). 
Renovate bot is self-hosted, version 37.381.6.
Read access to private gitlab modules:
```bash
login renovejttest
password glpat-syFsz-ZZsra91QkizCdm
```

```bash
GOPROXY=https://proxy.golang.org, direct
GOPRIVATE=https://gitlab.com/test-renovate-group*
```

## Current behavior

Renovate can't get submodule from gitlab repository
```bash
DEBUG: GET https://proxy.golang.org/gitlab.com/test-renovate-group/test-renovate-subgroup/module-test/api/@v/list = (code=ERR_NON_2XX_3XX_RESPONSE, statusCode=404 retryCount=0, duration=2130) (repository=addudko/renovate-issue-28540)
DEBUG: GET https://proxy.golang.org/gitlab.com/test-renovate-group/test-renovate-subgroup/module-test/@v/list = (code=ERR_NON_2XX_3XX_RESPONSE, statusCode=404 retryCount=0, duration=2144) (repository=addudko/renovate-issue-28540)
DEBUG: Go lookup source url https://gitlab.com/test-renovate-group/test-renovate-subgroup for module gitlab.com/test-renovate-group/test-renovate-subgroup/module-test/api (repository=addudko/renovate-issue-28540)
DEBUG: Go lookup source url https://gitlab.com/test-renovate-group/test-renovate-subgroup for module gitlab.com/test-renovate-group/test-renovate-subgroup/module-test (repository=addudko/renovate-issue-28540)
DEBUG: GET https://gitlab.com/api/v4/projects/test-renovate-group%2Ftest-renovate-subgroup%2Fmodule-test%2Fapi/repository/tags?per_page=100 = (code=ERR_NON_2XX_3XX_RESPONSE, statusCode=404 retryCount=0, duration=328) (repository=addudko/renovate-issue-28540)
DEBUG: GitLab API 404 (repository=addudko/renovate-issue-28540)
       "url": "https://gitlab.com/api/v4/projects/test-renovate-group%2Ftest-renovate-subgroup%2Fmodule-test%2Fapi/repository/tags?per_page=100"
DEBUG: Goproxy error: trying next URL provided with GOPROXY (repository=addudko/renovate-issue-28540)
       "err": {
         "name": "HTTPError",
         "code": "ERR_NON_2XX_3XX_RESPONSE",
         "timings": {
           "start": 1719473887452,
           "socket": 1719473887452,
           "lookup": 1719473887453,
           "connect": 1719473887484,
           "secureConnect": 1719473887540,
           "upload": 1719473887540,
           "response": 1719473887780,
           "end": 1719473887780,
           "phases": {
             "wait": 0,
             "dns": 1,
             "tcp": 31,
             "tls": 56,
             "request": 0,
             "firstByte": 240,
             "download": 0,
             "total": 328
           }
         },
         "message": "Response code 404 (Not Found)",
         "stack": "HTTPError: Response code 404 (Not Found)\n    at Request.<anonymous> (/usr/local/renovate/node_modules/.pnpm/got@11.8.6/node_modules/got/dist/source/as-promise/index.js:118:42)\n    at processTicksAndRejections (node:internal/process/task_queues:95:5)",
         "options": {
           "headers": {
             "user-agent": "RenovateBot/37.387.2 (https://github.com/renovatebot/renovate)",
             "accept": "application/json",
             "accept-encoding": "gzip, deflate, br"
           },
           "url": "https://gitlab.com/api/v4/projects/test-renovate-group%2Ftest-renovate-subgroup%2Fmodule-test%2Fapi/repository/tags?per_page=100",
           "hostType": "gitlab-tags",
           "username": "",
           "password": "",
           "method": "GET",
           "http2": false
         },
         "response": {
           "statusCode": 404,
           "statusMessage": "Not Found",
           "body": {"message": "404 Project Not Found"},
           "headers": {
             "date": "Thu, 27 Jun 2024 07:38:07 GMT",
             "content-type": "application/json",
             "content-length": "35",
             "connection": "close",
             "cache-control": "no-cache",
             "content-security-policy": "default-src 'none'",
             "vary": "Origin, Accept-Encoding",
             "x-content-type-options": "nosniff",
             "x-frame-options": "SAMEORIGIN",
             "x-gitlab-meta": "{\"correlation_id\":\"066dcbcbf54d3a2990ecf7811f3cef06\",\"version\":\"1\"}",
             "x-request-id": "066dcbcbf54d3a2990ecf7811f3cef06",
             "x-runtime": "0.023730",
             "strict-transport-security": "max-age=31536000",
             "referrer-policy": "strict-origin-when-cross-origin",
             "gitlab-lb": "haproxy-main-24-lb-gprd",
             "gitlab-sv": "api-gke-us-east1-b",
             "cf-cache-status": "MISS",
             "report-to": "{\"endpoints\":[{\"url\":\"https:\\/\\/a.nel.cloudflare.com\\/report\\/v4?s=awVeYEmJEG9wRvNCCj1rY%2B5aUEyHTpBih3B19KAfPf%2BBPvTDqc%2BhYSoTocFaT2DnnV5R8uvkYlmEnGJVc4nvHHMyTXxl1HyKqrO39vu9cTRnnk58LlwnWxwhBsk%3D\"}],\"group\":\"cf-nel\",\"max_age\":604800}",
             "nel": "{\"success_fraction\":0.01,\"report_to\":\"cf-nel\",\"max_age\":604800}",
             "set-cookie": [
               "_cfuvid=0zhAsHF3yVdMVGqU.xkL2y1QEBTc2O7fPnmaZXaCAHo-1719473887760-0.0.1.1-604800000; path=/; domain=.gitlab.com; HttpOnly; Secure; SameSite=None"
             ],
             "server": "cloudflare",
             "cf-ray": "89a3c6954b1eabf9-KLD"
           },
           "httpVersion": "1.1",
           "retryCount": 0
         }
       }
DEBUG: Failed to look up go package gitlab.com/test-renovate-group/test-renovate-subgroup/module-test/api (repository=addudko/renovate-issue-28540, packageFile=go.mod, dependency=gitlab.com/test-renovate-group/test-renovate-subgroup/module-test/api)
DEBUG: GET https://gitlab.com/api/v4/projects/test-renovate-group%2Ftest-renovate-subgroup%2Fmodule-test/repository/tags?per_page=100 = (code=ERR_NON_2XX_3XX_RESPONSE, statusCode=404 retryCount=0, duration=308) (repository=addudko/renovate-issue-28540)
DEBUG: GitLab API 404 (repository=addudko/renovate-issue-28540)
       "url": "https://gitlab.com/api/v4/projects/test-renovate-group%2Ftest-renovate-subgroup%2Fmodule-test/repository/tags?per_page=100"
DEBUG: Goproxy error: trying next URL provided with GOPROXY (repository=addudko/renovate-issue-28540)
       "err": {
         "name": "HTTPError",
         "code": "ERR_NON_2XX_3XX_RESPONSE",
         "timings": {
           "start": 1719473887551,
           "socket": 1719473887552,
           "lookup": 1719473887553,
           "connect": 1719473887582,
           "secureConnect": 1719473887629,
           "upload": 1719473887629,
           "response": 1719473887859,
           "end": 1719473887859,
           "phases": {
             "wait": 1,
             "dns": 1,
             "tcp": 29,
             "tls": 47,
             "request": 0,
             "firstByte": 230,
             "download": 0,
             "total": 308
           }
         },
         "message": "Response code 404 (Not Found)",
         "stack": "HTTPError: Response code 404 (Not Found)\n    at Request.<anonymous> (/usr/local/renovate/node_modules/.pnpm/got@11.8.6/node_modules/got/dist/source/as-promise/index.js:118:42)\n    at processTicksAndRejections (node:internal/process/task_queues:95:5)",
         "options": {
           "headers": {
             "user-agent": "RenovateBot/37.387.2 (https://github.com/renovatebot/renovate)",
             "accept": "application/json",
             "accept-encoding": "gzip, deflate, br"
           },
           "url": "https://gitlab.com/api/v4/projects/test-renovate-group%2Ftest-renovate-subgroup%2Fmodule-test/repository/tags?per_page=100",
           "hostType": "gitlab-tags",
           "username": "",
           "password": "",
           "method": "GET",
           "http2": false
         },
         "response": {
           "statusCode": 404,
           "statusMessage": "Not Found",
           "body": {"message": "404 Project Not Found"},
           "headers": {
             "date": "Thu, 27 Jun 2024 07:38:07 GMT",
             "content-type": "application/json",
             "content-length": "35",
             "connection": "close",
             "cache-control": "no-cache",
             "content-security-policy": "default-src 'none'",
             "vary": "Origin, Accept-Encoding",
             "x-content-type-options": "nosniff",
             "x-frame-options": "SAMEORIGIN",
             "x-gitlab-meta": "{\"correlation_id\":\"72c35d64c9ce51451df7cbebb33fb94f\",\"version\":\"1\"}",
             "x-request-id": "72c35d64c9ce51451df7cbebb33fb94f",
             "x-runtime": "0.022792",
             "strict-transport-security": "max-age=31536000",
             "referrer-policy": "strict-origin-when-cross-origin",
             "gitlab-lb": "haproxy-main-26-lb-gprd",
             "gitlab-sv": "api-gke-us-east1-d",
             "cf-cache-status": "MISS",
             "report-to": "{\"endpoints\":[{\"url\":\"https:\\/\\/a.nel.cloudflare.com\\/report\\/v4?s=zEbP3St7reJs0sfEWB40n%2BXcHBiJS3PrByHk4V%2FfbA539Qo2TsULAEztfJ8O5VBf0yBTglgGzkXqvI9eXcy81PoHjn5HAjZP6C92O1jO5%2Fp6iVhJDkY1IvNnDHY%3D\"}],\"group\":\"cf-nel\",\"max_age\":604800}",
             "nel": "{\"success_fraction\":0.01,\"report_to\":\"cf-nel\",\"max_age\":604800}",
             "set-cookie": [
               "_cfuvid=x1FbeVvjY1N.MssZUOm1i2DGhWv96Lbiq_pyuMjaoYs-1719473887839-0.0.1.1-604800000; path=/; domain=.gitlab.com; HttpOnly; Secure; SameSite=None"
             ],
             "server": "cloudflare",
             "cf-ray": "89a3c695dbeeabfe-KLD"
           },
           "httpVersion": "1.1",
           "retryCount": 0
         }
       }
DEBUG: Failed to look up go package gitlab.com/test-renovate-group/test-renovate-subgroup/module-test (repository=addudko/renovate-issue-28540, packageFile=go.mod, dependency=gitlab.com/test-renovate-group/test-renovate-subgroup/module-test)
```

## Expected behavior

Renovate checks private module and submodule in the same repo correctly.

## Link to the Renovate issue or Discussion

https://github.com/renovatebot/renovate/issues/28540
