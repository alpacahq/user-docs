# Alpaca User Documentation

## Building the Site

1. Install [hugo](https://gohugo.io/) (`brew install hugo` on osx should do it)
1. In the project directory, run `hugo`
1. Open public/index.html

Yeah that's really it.

## Deployment

`make publish` will push to the s3 bucket using `aws` command.

Example:
```
$ make publish
```

Drone is set up so if you push to the branch `production` it publishes to the production bucket.

```
$ git push origin master master:production
```
