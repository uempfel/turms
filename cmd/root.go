/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/atc0005/go-teams-notify/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var webHookUrl string
var text string
var color string
var title string
var fileBody string
var overriddenUrl string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "turms",
	Short: "Send messages to Teams channels",
	Long: `Send messages to Teams channels via incoming webhooks.	

For Information on how to configure an incoming webhook in Teams,
please refer to https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook

To configure the CLI, export the following environment variable:
export TURMS_URL=https://some-tenant.webhook.office.com/webhookb2/some-id/IncomingWebhook/some-other-id/yet-another-id
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args [] string) {

		if len(text) == 0 && len (fileBody) == 0 {
			fmt.Println("❌ ", "No body provided")
			fmt.Println("Either --body (-b) or --body-from-file (-f) must be supplied")
			os.Exit(1)
		}

		if len(webHookUrl) == 0 && len (overriddenUrl) == 0 {
			fmt.Println("❌ ", "No webhook url provided")
			fmt.Println("To configure the webhook url, either export TURMS_URL or use the --url (-u) flag")
			os.Exit(1)
		}
		
		colorMap := getColorMap()
		
		c := colorMap[color]
		if len(c) > 0 {
			color = c
		}

		if len(overriddenUrl) > 0 {
			webHookUrl = overriddenUrl
		}
		_, err := goteamsnotify.IsValidWebhookURL(webHookUrl)
		if err != nil {
			fmt.Println("❌ ", err)
			os.Exit(1)
		}

		if len(fileBody) > 0 {

			source, err := ioutil.ReadFile(fileBody)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			text = string(source)
		}

		err = sendTheMessage()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("📬 ", "Message Sent!")
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	webHookUrl = os.Getenv("TURMS_URL")
	rootCmd.Flags().StringVarP(&text, "body", "b", "", "text body to send. Required, if \"--body-from-file\" is not set")
	rootCmd.Flags().StringVarP(&color, "color", "c", "", "theme color to display in the message (webcolors or hexcodes are supported)")
	rootCmd.Flags().StringVarP(&title, "title", "t", "", "title to display in the message")
	rootCmd.Flags().StringVarP(&fileBody, "body-from-file", "f", "", "path to a markdown file to send as body (takes precedence over the \"--body\" flag)\nRequired, if \"--body\" is not set")
	rootCmd.Flags().StringVarP(&overriddenUrl, "url", "u", "", "webhook Url (overrides $TURMS_URL)")
}

func sendTheMessage() error {
	// init the client
	mstClient := goteamsnotify.NewClient()
	mstClient.SkipWebhookURLValidationOnSend(true)
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = title
	msgCard.Text = text
	msgCard.ThemeColor = color

	// send
	return mstClient.Send(webHookUrl, msgCard)
}

func getColorMap() map[string]string {
   
	//from https://github.com/bahamas10/css-color-names/blob/master/css-color-names.json
	return map[string]string { 
		"aliceblue":            "#f0f8ff",
		"antiquewhite":         "#faebd7",
		"aqua":                 "#00ffff",
		"aquamarine":           "#7fffd4",
		"azure":                "#f0ffff",
		"beige":                "#f5f5dc",
		"bisque":               "#ffe4c4",
		"black":                "#000000",
		"blanchedalmond":       "#ffebcd",
		"blue":                 "#0000ff",
		"blueviolet":           "#8a2be2",
		"brown":                "#a52a2a",
		"burlywood":            "#deb887",
		"cadetblue":            "#5f9ea0",
		"chartreuse":           "#7fff00",
		"chocolate":            "#d2691e",
		"coral":                "#ff7f50",
		"cornflowerblue":       "#6495ed",
		"cornsilk":             "#fff8dc",
		"crimson":              "#dc143c",
		"cyan":                 "#00ffff",
		"darkblue":             "#00008b",
		"darkcyan":             "#008b8b",
		"darkgoldenrod":        "#b8860b",
		"darkgray":             "#a9a9a9",
		"darkgreen":            "#006400",
		"darkgrey":             "#a9a9a9",
		"darkkhaki":            "#bdb76b",
		"darkmagenta":          "#8b008b",
		"darkolivegreen":       "#556b2f",
		"darkorange":           "#ff8c00",
		"darkorchid":           "#9932cc",
		"darkred":              "#8b0000",
		"darksalmon":           "#e9967a",
		"darkseagreen":         "#8fbc8f",
		"darkslateblue":        "#483d8b",
		"darkslategray":        "#2f4f4f",
		"darkslategrey":        "#2f4f4f",
		"darkturquoise":        "#00ced1",
		"darkviolet":           "#9400d3",
		"deeppink":             "#ff1493",
		"deepskyblue":          "#00bfff",
		"dimgray":              "#696969",
		"dimgrey":              "#696969",
		"dodgerblue":           "#1e90ff",
		"firebrick":            "#b22222",
		"floralwhite":          "#fffaf0",
		"forestgreen":          "#228b22",
		"fuchsia":              "#ff00ff",
		"gainsboro":            "#dcdcdc",
		"ghostwhite":           "#f8f8ff",
		"goldenrod":            "#daa520",
		"gold":                 "#ffd700",
		"gray":                 "#808080",
		"green":                "#008000",
		"greenyellow":          "#adff2f",
		"grey":                 "#808080",
		"honeydew":             "#f0fff0",
		"hotpink":              "#ff69b4",
		"indianred":            "#cd5c5c",
		"indigo":               "#4b0082",
		"ivory":                "#fffff0",
		"khaki":                "#f0e68c",
		"lavenderblush":        "#fff0f5",
		"lavender":             "#e6e6fa",
		"lawngreen":            "#7cfc00",
		"lemonchiffon":         "#fffacd",
		"lightblue":            "#add8e6",
		"lightcoral":           "#f08080",
		"lightcyan":            "#e0ffff",
		"lightgoldenrodyellow": "#fafad2",
		"lightgray":            "#d3d3d3",
		"lightgreen":           "#90ee90",
		"lightgrey":            "#d3d3d3",
		"lightpink":            "#ffb6c1",
		"lightsalmon":          "#ffa07a",
		"lightseagreen":        "#20b2aa",
		"lightskyblue":         "#87cefa",
		"lightslategray":       "#778899",
		"lightslategrey":       "#778899",
		"lightsteelblue":       "#b0c4de",
		"lightyellow":          "#ffffe0",
		"lime":                 "#00ff00",
		"limegreen":            "#32cd32",
		"linen":                "#faf0e6",
		"magenta":              "#ff00ff",
		"maroon":               "#800000",
		"mediumaquamarine":     "#66cdaa",
		"mediumblue":           "#0000cd",
		"mediumorchid":         "#ba55d3",
		"mediumpurple":         "#9370db",
		"mediumseagreen":       "#3cb371",
		"mediumslateblue":      "#7b68ee",
		"mediumspringgreen":    "#00fa9a",
		"mediumturquoise":      "#48d1cc",
		"mediumvioletred":      "#c71585",
		"midnightblue":         "#191970",
		"mintcream":            "#f5fffa",
		"mistyrose":            "#ffe4e1",
		"moccasin":             "#ffe4b5",
		"navajowhite":          "#ffdead",
		"navy":                 "#000080",
		"oldlace":              "#fdf5e6",
		"olive":                "#808000",
		"olivedrab":            "#6b8e23",
		"orange":               "#ffa500",
		"orangered":            "#ff4500",
		"orchid":               "#da70d6",
		"palegoldenrod":        "#eee8aa",
		"palegreen":            "#98fb98",
		"paleturquoise":        "#afeeee",
		"palevioletred":        "#db7093",
		"papayawhip":           "#ffefd5",
		"peachpuff":            "#ffdab9",
		"peru":                 "#cd853f",
		"pink":                 "#ffc0cb",
		"plum":                 "#dda0dd",
		"powderblue":           "#b0e0e6",
		"purple":               "#800080",
		"rebeccapurple":        "#663399",
		"red":                  "#ff0000",
		"rosybrown":            "#bc8f8f",
		"royalblue":            "#4169e1",
		"saddlebrown":          "#8b4513",
		"salmon":               "#fa8072",
		"sandybrown":           "#f4a460",
		"seagreen":             "#2e8b57",
		"seashell":             "#fff5ee",
		"sienna":               "#a0522d",
		"silver":               "#c0c0c0",
		"skyblue":              "#87ceeb",
		"slateblue":            "#6a5acd",
		"slategray":            "#708090",
		"slategrey":            "#708090",
		"snow":                 "#fffafa",
		"springgreen":          "#00ff7f",
		"steelblue":            "#4682b4",
		"tan":                  "#d2b48c",
		"teal":                 "#008080",
		"thistle":              "#d8bfd8",
		"tomato":               "#ff6347",
		"turquoise":            "#40e0d0",
		"violet":               "#ee82ee",
		"wheat":                "#f5deb3",
		"white":                "#ffffff",
		"whitesmoke":           "#f5f5f5",
		"yellow":               "#ffff00",
		"yellowgreen":          "#9acd32",
	}
}
