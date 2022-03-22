# Alpaca User Documentation [:warning: Deprecated]

This repo is no longer maintained. For latest documentation and OpenAPI specs visit: https://github.com/alpacahq/alpaca-docs

This repository powers the documentation for Alpaca's users.

Its contents are hosted on Alpaca's site [here](https://docs.alpaca.markets).

## Submitting Improvements

We are accepting pull requests for this repository, so if you'd like to make some improvements to the documentation,
we'd certainly appreciate it! To submit a correction or any other sort of improvment, you can fork this repository
on GitHub, make your changes, and send them as a pull request.

## Building the Site Locally

You can make changes to the markdown files holding site content here on GitHub, but if you'd like to see your changes
locally before you submit, you can follow these steps.

1. Install [hugo](https://gohugo.io/). If on OSX, install [homebrew](https://brew.sh/) and run `brew install hugo`.
1. In the project directory, run `hugo server`.
1. If you follow the link in hugo's output, you'll be taken to a local version of the site with the changes you've made.

On your local version of the site, you can check any pages that will be affected by the changes you've made.

## Custom Shortcodes

Hugo supports [shortcodes](https://gohugo.io/content-management/shortcodes/), which allow you to add HTML to a markdown
file without making the code too messy. We've added a few custom shortcodes of our own.

### code-example
The code-example shortcode allows you to add a code block with tabs that can switch between different languages. This
allows us to show readers how to do the same thing in a wider variety of languages without cluttering the page. You will
need to add the example code files to the system and then provide the shortcode with a path to them. In order to support
multiple code blocks on the same page, all files in the same block should be named the same, only distinguished by their
extension, and the name you choose should also be given to the shortcode as `exampleId`.

Example usage, given a set of files like `ex1.py`, `ex1.js`, and `ex1.cs` stored in an `examples/` directory that you
want to include on the orders page:
```
{{< code-example exampleId="ex1" pathURL="/web-api/fun/examples" >}}
```

(Note that when you only have one example to show and the tabs are unneccessary, you can forego the shortcode and use
regular markdown syntax to format it as a code block.)

### rest-endpoint
This shortcode lets you show a formatted endpoint description. Descriptions are stored in `data/webapi/endpoints` as
YAML files.

```
{{< rest-endpoint resource="account" method="GET" path="/v1/account" >}}
```

### rest-entity-desc
This can be used to show a formatted API description. These descriptions should be stored in `data/webapi/entities` as
in the `spec` field of a YAML file.

```
{{< rest-entity-desc name="account">}}
```

### rest-entity-example
This can be used to show a formatted API example. These examples should be stored in `data/webapi/entities` in the `
`example` field of a YAML file.

```
{{< rest-entity-example name="account">}}
```