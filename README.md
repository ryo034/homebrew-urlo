# urlo CLI
![license](https://img.shields.io/badge/license-Apache2.0-blue)
![made with language](https://img.shields.io/badge/Made%20with-Go-00ADD8.svg)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/ryo034/homebrew-urlo)](https://github.com/ryo034/homebrew-urlo/releases)

urlo is a Command Line Interface (CLI) tool for managing and interacting with a local Json file containing URLs.

The tool is designed to offer quick and convenient access to URLs, with options for adding new URLs, listing all URLs, selecting a URL to open, and more.

urlo enables developers to effortlessly share a list of URLs.

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
> urlo add google https://www.google.com
```

#### Override the URL
```shell
> urlo add {title} {url} -o

# Example
> urlo add google https://www.google.co.jp -o
```

### Bulk add URLs
To add multiple URLs to the local Json file, use the add command:

```shell
> urlo bulk-add {json string}
```

### List all URLs
To list all URLs from the local Json file, use the list command:

```shell
> urlo list
google - https://www.google.com
yahoo  - https://www.yahoo.com
```

Use the -j or -s option to display the list in JSON format:

```shell
> urlo list -j
[
  {
    "title": "google",
    "url": "https://www.google.com"
  },
  {
    "title": "yahoo",
    "url": "https://www.yahoo.com"
  }
]

> urlo list -s
'[{title: "google", url: "https://www.google.com"},{title: "yahoo", url: "https://www.yahoo.com"}]'
```

### Set the list
To set with override the list of URLs from a JSON string, use the set command:

```shell
> urlo list
bing - https://www.bing.com

> urlo set '[{"title": "google", "url": "https://www.google.com"},{"title": "yahoo", "url": "https://www.yahoo.com"}]'

> urlo list
google - https://www.google.com
yahoo  - https://www.yahoo.com
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

## License
urlo is released under the Apache License 2.0 License. See LICENSE for more information.

## Data Privacy and Security
One of the key features of our library is how it handles data. Rather than relying on external databases or cloud storage, this library manages all data using local files on your system.
This ensures that your data stays where it belongs: with you.
With this approach, we prioritize your privacy and security, giving you full control over your data without compromising its accessibility.

## Share your list
You can share your list with others by using the pbcopy command:

```shell
> urlo list -s | pbcopy
> urlo set -s "{set output json string}"
```

## Maintenance and Support
As an open-source tool, we greatly value and rely on our community. The feedback, questions, and involvement from users not only help us to constantly improve and develop, but also make this project more robust and user-friendly.

In an open-source project, maintaining quality and providing consistent support are key. We are dedicated to fixing bugs, adding new features, and providing support as questions or issues arise. However, please understand that this is a community project, and responses or solutions may take time.

We appreciate your patience and encourage active contribution to help improve our tool. After all, the strength of open-source lies in the collective power of a community working together.
