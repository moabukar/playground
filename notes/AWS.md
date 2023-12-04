# AWS

## API GW vs ALB

- Can implement rate limiting on API GW but not ALB
- Both can integrate with WAF for protection
- API GW accepts HTTPS traffic only, ALB can accept HTTP and HTTPS
- API GW able to do request validation, request-response mapping, ALB cannot
- API GW can export/import cross API platforms using swagger, Open API 3.0, ALB cannot
- API GW can do caching, ALB cannot
- API GW time out is 30 seconds, ALB is 4000 seconds
- API GW integrates with almost all AWS services, ALB only integrates with EC2, Lambda, IP targets
- API GW can't do health checks, ALB can
- API GW pay per use unless you have caching enable, ALB pay per use (underlying EC2 instances)
- Price is not straightforward, depends on the number of requests, data transfer, caching, etc.
  - If your workload is spiky and idle for a long time, API GW is probably cheaper
  - But if your workload is consistently using high volume, ALB is probably cheaper


