# publish-md-to-gdocs-action
Action for publishing MarkDown files to Google Docs.


https://cloud.google.com/blog/products/identity-security/enabling-keyless-authentication-from-github-actions

```yaml
steps:
- id: 'auth'
  name: 'Authenticate to Google Cloud'
  uses: 'google-github-actions/auth@v0.4.0'
  with:
    workload_identity_provider: 'projects/123456789/locations/global/workloadIdentityPools/my-pool/providers/my-provider'
    service_account: 'my-service-account@my-project.iam.gserviceaccount.com'

- name: Publish

```

## Uses:
Go give a start to the following packages:  
[gomarkdown](https://github.com/gomarkdown/markdown)  
[bluemonday](https://github.com/microcosm-cc/bluemonday)