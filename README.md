# sitehealth

Simple tool to check a hosts HTTP status and the time a request takes.

## Installation

Build the app:<br />
`go build`

Run:<br />
`./sitehealth [host]` (*Unix/Mac*)<br />
`sitehealth.exe [host]` (*Windows*)

## Usage

```shell
$ sitehealth [host]
        host    A valid URL to any hostname prefixed by "http://" or "https://"
```