# Junction
Junction serves as a bridge between SMTP and [Apprise](https://github.com/caronc/apprise). Any emails sent to Junction's SMTP server will be forwarded to the configured notification services.

This project was inspired by [Mailrise](https://github.com/YoRyan/mailrise). I wanted a solution that was more configurable, and not constrained to specific email addresses. And so, Junction was born.

Junction matches an incoming email to an Apprise URL by any combination of:
 * The address(es) the email was sent to
 * The address the email was sent from
 * The IP of the client that sent the email

Once matched, Junction will create the notification from the optionally provided template, and send it to the provided Apprise URL.

## Configuration
Junction is configured with a simple yaml file. By default, this file is read from `<APP DIRECTORY>/config/config.yaml`.
The config is composed of the following keys:
| Key          | Value                 | Default | Required |
| ---          | -----                 | ------- | -------- |
| `log-level:` | `debug` or `info`     | `info`  | No       |
| `port:`      | Any open port number  | `8025`  | No       |
| `junctions:` | An array of Junctions | None    | Yes      |

Each junction is configured as follows:
| Key                      | Value                 | Default | Required |
| ---                      | -----                 | ------- | -------- |
| `name:`                  | Any name for your use | None    | No       |
| `apprise:`               | The [Apprise URL](https://github.com/caronc/apprise/wiki/URLBasics) to send the notification to | None | Yes |
| `to:`                    | Configuration for the received email's To addresses | none | No |
| `to.emails:`      | An email address, or array of addresses, that the received email was sent to | None | No |
| `to.require-all:` | `true` or `false`. If `true` this will only match if the email was sent to *all* of the listed emails. If unset, or `false`, this will match if any address listed is present  | `false` | No |
| `from:`                  | Configuration for the received email's From address | None | No |
| `from.email:`     | An email address that the received email was from | None | No |
| `from.ip:`        | The IP address of the client that sent the email | None | No |
| `title:`                 | The text used in the notification's title | The email's subject | No |
| `body:`                  | The text used in the notification's body  | The email's body    | No |

Any not required field without a default value that is left blank will be considered as matched by default.

Junction searches for matches top to bottom. Junctions that are more specific should be placed higher.

### Example
```
log-level: debug
port: 825
junctions:
  - name: Anything
    apprise: service://configuration/?parameters
    to:
      emails: junction@domain.tld
  - name: Anything2
    apprise: service://configuration/?parameters
    from:
      email: junction@domain.tld
  - apprise: service://configuration/?parameters
    from:
      ip: 1.1.1.1
  - apprise: service://configuration/?parameters
```
With this configuration:
* Any email sent to the address junction@domain.tld will match the first Junction
* Any email sent to a different address, but comes from junction@domain.tld will match the second Junction
* Any email that is not to or from junction@domain.tld, but came from the IP 1.1.1.1 will match the third Junction
* Any email not matching the first three will match the last Junction as a catch all

### Templating
Junction supports templating for `title` and `body` with Golang's [text/template](https://pkg.go.dev/text/template) package.
Available variables:
* Subject: The email's subject
* To: The address(es) the email was sent to 
* From: The address the email was sent from
* Date: The date the email was sent

Please note that you must use a `.` before the variable name. `{{ .Subject }}` will work. `{{ Subject }}` will not.

## Installation
Junction can be run as a container, or directly from the binary file. Once installed and configured, simply set your applications outbound SMTP server to Junction's IP and port.

### Docker
A container image is available through GitHub Container Registry.  
To use it, mount your `config.yaml` to `/app/config/config.yaml` and map container port `8025` to a host port of your choosing.

```docker run -p 8025:8025 -v /local-path/config.yaml:/app/config/config.yaml ghcr.io/kenneth-church/junction ```

### Binary
Download the latest release and place it where you'd like it, create a `config` directory and place your `config.yaml` within it.

If desired, you can change the location of the configuration file with the `CONF_PATH` environment variable.

```CONF_PATH='/whatever/path/you/want ./junction```

[Apprise](https://github.com/caronc/apprise) will also need to be installed and available on the machine under the `apprise` command.

## Planned Features
- [ ] Support for Apprise configuration files
- [ ] SMTP server authentication
- [ ] Optional configuration web UI
- [ ] Option to save emails in a database