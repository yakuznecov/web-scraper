# Simple Web scraper

This is a simple web scraper written in Go. It accepts a list of URLs as input, crawls these pages, extracts all the links (URLs) from each page, and outputs the unique URLs found.

## Features

- Parallel crawling using goroutines
- Extraction of links from HTML
- Collection of unique URLs
- Simple and efficient design

## Requirements

- Go 1.16 or later

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yakuznecov/web-scraper.git
   cd web-scraper
   ```

2. Build the project:

   ```sh
   go build -o scraper
   ```

## Usage

To run the scraper, provide a list of seed URLs as arguments:

```sh
./scraper http://example.com http://example.org

```
