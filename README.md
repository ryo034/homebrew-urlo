# urlo CLI
urlo is a Command Line Interface (CLI) tool for managing and interacting with a local CSV file containing URLs.

The tool is designed to offer quick and convenient access to URLs, with options for adding new URLs, listing all URLs, selecting a URL to open, and more.

## Installation
On macOS, urlo can be installed with Homebrew:

```shell
> brew install urlo
```

## Usage
### Display the version
To display the version of urlo, use the -v option:

```shell
> urlo -v
```

### Add a new URL
To add a new URL to the local CSV file, use the add command:

```shell
> urlo add {title} {url}

# Example
> urlo add Google https://google.com
```
### List all URLs
To list all URLs from the local CSV file, use the list command:

```shell
> urlo list
google
yahoo
```

Use the -u option to also display the URLs:

```shell
> urlo list -u
google - https://google.com
yahoo  - https://yahoo.com
```
### Open a URL
To open a URL by its title, use the open command:

```shell
> urlo open {title}

# Example
> urlo open Google
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
urlo is released under the MIT License. See LICENSE for more information.
