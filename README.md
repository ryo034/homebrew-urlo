# urlo CLI
urlo is a Command Line Interface (CLI) tool for managing and interacting with a local Json file containing URLs.

The tool is designed to offer quick and convenient access to URLs, with options for adding new URLs, listing all URLs, selecting a URL to open, and more.

## Installation
On macOS, urlo can be installed with Homebrew:

```shell
> brew install ryo034/urlo/homebrew-urlo
```

## Usage
### Display the version
To display the version of urlo, use the -v option:

```shell
> urlo -v
```

### Add a new URL
To add a new URL to the local Json file, use the add command:

```shell
> urlo add {title} {url}

# Example
> urlo add Google https://google.com
```

### List all URLs
To list all URLs from the local Json file, use the list command:

```shell
> urlo list
google - https://google.com
yahoo  - https://yahoo.com
```

Use the -j or -s option to display the list in JSON format:

```shell
> urlo list -j
[
  {
    "title": "google",
    "url": "https://google.com"
  },
  {
    "title": "yahoo",
    "url": "https://yahoo.com"
  }
]

> urlo list -s
'[{title: "google", url: "https://google.com"},{title: "yahoo", url: "https://yahoo.com"}]'
```

if you want share the list with others, use the pbcopy command:

```shell
> urlo list -s | pbcopy
> urlo set -s "{set output json string}"
```

### Set the list
To set the list of URLs from a JSON string, use the set command:

```shell
> urlo list
No records found

> urlo set '[{"title": "google", "url": "https://google.com"},{"title": "yahoo", "url": "https://yahoo.com"}]'

> urlo list
google - https://google.com
yahoo  - https://yahoo.com
```

### Open a URL
To open a URL by its title, use the open command:

```shell
> urlo open {title}

# Example
> urlo open google
```

### Select a URL to open
To display a list of all URLs and select one to open, use the select command:

```shell
> urlo select
? Select a Website: 
  ▸ google
    yahoo
```
Use the -q option to filter the list with a regular expression:

```shell
> urlo select -q g
? Select a Website: 
  ▸ google
  
> urlo select -q y
? Select a Website: 
  ▸ yahoo
```

### License
urlo is released under the GPT-3.0 License. See LICENSE for more information.
