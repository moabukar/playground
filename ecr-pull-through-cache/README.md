# ECR pull through cache for images

Pull through cache rules for ECR have been first announced at re-Invent 2021. I believe this was a huge feature but at the time it was lacking support for Docker Hub which is essentially home for majority of images. This has changed recently and so I am going to demonstrate how to use it but first, what is pull through cache. With pull through cache rules, you can sync the contents of an upstream registry with your Amazon ECR private registry.


## How it works

For the upstream registries that require authentication, you must store your credentials in an AWS Secrets Manager secret. The Amazon ECR console makes it easy for you to create the Secrets Manager secret for each of the authenticated upstream registries. For more information on creating a Secrets Manager secret using the Secrets Manager console, see Storing your upstream repository credentials in an AWS Secrets Manager secret. I will also demonstrate how to create such secret later in the story.

After you’ve created a pull through cache rule for the upstream registry, simply pull an image from that upstream registry using your Amazon ECR private registry URI. Amazon ECR then creates a repository and caches that image in your private registry. On your subsequent pull requests of the cached image with a given tag, Amazon ECR checks the upstream registry to see if there is a new version of the image with that specific tag and attempts to update the image in your private registry at least once every 24 hours.

## validation

```sh
aws ecr validate-pull-through-cache-rule \
     --ecr-repository-prefix ecr \
     --region eu-west-2
```

