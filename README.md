# check-css-selector

Check if a CSS selector exists in the HTML DOM of a specified URL.

## Synopsis

```
check-css-selector -U <URL> -S <CSS_SELECTOR>
```

## Description

This plugin fetches the specified URL and checks whether the given CSS selector exists in the HTML DOM. It returns:
- `OK` if the selector is found
- `CRITICAL` if the selector is not found
- `UNKNOWN` if there is an error accessing the URL or parsing the HTML

## Installation

### Using go install

```
go install github.com/kga/go-check-css-selector@latest
```

### Building from source

```
git clone https://github.com/kga/go-check-css-selector.git
cd go-check-css-selector
go build
```

## Usage

### Command line

```
check-css-selector -U https://example.com -S "div.content"
```

### With mackerel-agent

Add the following configuration to your `mackerel-agent.conf`:

```
[plugin.checks.check-css-selector-sample]
command = ["check-css-selector", "-U", "https://example.com", "-S", "div.content"]
```

## Options

```
  -U, --url=      URL to check (required)
  -S, --selector= CSS selector to find in the DOM (required)
```

## Examples

Check if a specific element exists on a page:
```
check-css-selector -U https://example.com -S "#main-content"
```

Check for a button with a specific class:
```
check-css-selector -U https://example.com -S "button.submit-btn"
```

Check for nested elements:
```
check-css-selector -U https://example.com -S "div.header nav.menu ul li"
```

## Author

[kga](https://github.com/kga)
