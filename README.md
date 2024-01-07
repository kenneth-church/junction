# Junction
Junction serves as a bridge between SMTP and [Apprise](https://github.com/caronc/apprise). Any emails sent to Junction's SMTP server will be forwarded to the configured notification services.

This project was inspired by [Mailrise](https://github.com/YoRyan/mailrise). I wanted a solution that was more configurable, and not constrained to specific email addresses. And so, Junction was born.

Junction matches an incoming email to an Apprise URL by any combination of:
 * The address(es) the email was sent to
 * The address the email was sent from
 * The IP of the client that sent the email

Once matched, Junction will create the notification from the optionally provided template, and send it to the provided Apprise URL.

## Configuration
Junction is configured with a yaml file. By default, this file is read from `<APP DIRECTORY>/config/config.yaml`.

Available configuration options and any applicable defaults are described below:

`log-level:` Optional. Defaults to `info`. Can be set to `debug` to output more information during application runtime.

`port:` Optional. Defaults to `8025`. The port to listen on for emails. Do not change if using Docker.

`junctions:` Required. A list of configurations that received emails are matched against.

Junctions are configured with the following values.

`name:` Optional. Just used for easier identification of the junction used. Has no effect on application execution.

`apprise:` Required. The [Apprise URL](https://github.com/caronc/apprise/wiki/URLBasics) to send to.

`to:` Optional. If not included, every incoming email will match the this portion of the junction.

&nbsp;&nbsp;`emails:` A list of email addresses that the received email must be sent to.

&nbsp;&nbsp;`require-all:` `true` or `false`, defaults to `false`. If set to `true`, and multiple email addresses are listed, every email address listed must be present for the received email to match the junction. If unset, or `false`, only one of the listed email addresses needs to be present.

`from:` Optional. If not included, every incoming email will match this portion of the junction.

&nbsp;&nbsp;`email:` Optional. The email address that the received email must be sent from.

&nbsp;&nbsp;`ip:` Optional. The IP Address of the machine that the received email must be sent from.

`title:` Optional. What is displayed in the notification's title. Defaults to the received email's subject. See [templating](#templating) below for further information.

`body:` Optional. What is displayed in the notification's body. Defaults to the received email's subject. See [templating](#templating) below for further information.


**Junctions are matched top down. More specific conditions should be placed to the top, and broader to the bottom**


### Examples

Minimal:
```yaml
junctions:
  - apprise: <Apprise URL>
```
With this configuration, every email received will be sent to the provided Apprise URL.


Specific:
```yaml
log-level: debug
port: 25
junctions:
  - name: Very Specific
    apprise: <Apprise URL>
    to:
      emails:
        - person1@example.com
        - person2@example.com
      require-all: true
    from:
      email: server@example.com
      ip: 1.1.1.1
  - name: Less Specific
    apprise: <Apprise URL>
    to:
      emails:
        - person1@example.com
    from:
      email: server@example.com
      ip: 1.1.1.1
  - name: Less Specific 2
    apprise: <Apprise URL>
    to:
      emails:
        - person2@example.com
    from:
      email: server@example.com
      ip: 1.1.1.1
  - name: Less Less Specific
    apprise: <Apprise URL>
    to:
      emails:
        - person1@example.com
        - person2@example.com
    from:
      email: server@example.com
      ip: 1.1.1.1
  - name: No To
    apprise: <Apprise URL>
    from:
      email: server@example.com
      ip: 1.1.1.1
  - name: No To
    apprise: <Apprise URL>
    from:
      email: server@example.com
      ip: 1.1.1.1
  - name: No To 2
    apprise: <Apprise URL>
    from:
      ip: 1.1.1.1
  - name: Catch All
    apprise: <Apprise URL>
```
With this configuration:
- Extra information will be output by the application for debugging
- The application will listen for emails on port `25` instead of the default `8025`
- For an email to be forwarded to `Very Specific` it must be: 
  - Sent to both `person1@example.com` and `person2@example.com`
  - Have been sent from `server@example.com` and from a machine with the IP Address `1.1.1.1`
- For an email to be forwarded to `Less Specific` or `Less Specific 2` it must be:
  - Sent to either `person1@example.com` and `person2@example.com` respectively
  - Have been sent from `server@example.com` and from a machine with the IP Address `1.1.1.1`
- For an email to be forwarded to `Less Less Specific` it must be:
  - Sent to either `person1@example.com` and `person2@example.com`
  - Have been sent from `server@example.com` and from a machine with the IP Address `1.1.1.1`
- For an email to be forwarded to `No To 2` it must be:
  - Sent to any email address(es)
  - Have been sent from anmy email address and from a machine with the IP Address `1.1.1.1`
- For an email to be forwarded to `Catch All` it must be:
  - Sent to and from any combination that does not match any of the above. Such as:
  - To: `person3@example.com` From: `server2@example.com` with the sending machine having an IP of `8.8.8.8`


## Templating
Junction supports templating for `title` and `body` fields with Golang's [text/template](https://pkg.go.dev/text/template) package.

Currently available variables:
- `To`: The address(es) the email was sent to preformatted with a comma delimiter
- `From`: The address the email was sent from
- `IP`: The IP Address of the server that sent the email
- `Date`: The date the email was sent
- `Subject`: The email's subject
- `Body`: The email's body
- `RawTo`: The raw content of the email's to field as a slice of strings

Please note that you must use a `.` before the variable name. `{{ .Subject }}` will work. `{{ Subject }}` will not.

For example, an email with the contents:
```
To: person@example.com
From: person2@example.com
Date: 01/01/01
Subject: A subject line
Body: A body

```
That matches a junction with the following configuration:
```yaml
...
title: "Email Subject: {{ .Subject }}"
body: "Email Date: {{ .Date }}\nEmail Body: {{ .Body }}"
...
```
Will send a notification with the title and body:
```
Email Subject: A subject line
---
Email Date: 01/01/01
Email Body: A body
```
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