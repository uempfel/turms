![Turms](https://user-images.githubusercontent.com/28337775/114446466-8a8e8c00-9bd1-11eb-88fd-38924dfdb0fc.png)

Turms ([the Etruscan messenger god, 𐌕𐌖𐌓𐌌𐌑](https://en.wikipedia.org/wiki/Turms)) is a CLI to send messages
to Microsoft Teams Channels via Incoming Webhooks.

## Installation
Assuming you have already [installed go](https://golang.org/doc/install):

```sh
export GO111MODULE=on
go get github.com/uempfel/turms
```

## Usage
For Information on how to configure an incoming webhook in Teams,
please refer to the [official docs at Microsoft](https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook).  

To configure the CLI, export the following environment variable:  
`export TURMS_URL=https://some-tenant.webhook.office.com/webhookb2/some-id/IncomingWebhook/some-other-id/yet-another-id`


```bash
Usage:
  turms [flags]

Flags:
  -b, --body string             text body to send. Required, if "--body-from-file" is not set
  -f, --body-from-file string   path to a markdown file to send as body (takes precedence over the "--body" flag)
                                Required, if "--body" is not set
  -c, --color string            theme color to display in the message (webcolors or hexcodes are supported)
  -h, --help                    help for turms
  -t, --title string            title to display in the message
  -u, --url string              webhook Url (overrides $TURMS_URL)
  ```
  
### Image Credits
* Gopher: [Maria Letta - Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack)
