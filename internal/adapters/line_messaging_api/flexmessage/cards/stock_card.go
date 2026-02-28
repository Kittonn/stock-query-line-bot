package cards

import (
	"fmt"

	"github.com/Kittonn/stock-query-line-bot/internal/adapters/line_messaging_api/flexmessage"
	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
)

const (
	colorGreen = "#16c784"
	colorRed   = "#ea3943"
	colorGray  = "#8b949e"
)

func priceChangeStyle(change float64) (color, arrow string) {
	switch {
	case change > 0:
		return colorGreen, "▲"
	case change < 0:
		return colorRed, "▼"
	default:
		return colorGray, "▬"
	}
}

func formatExchange(exchange string) string {
	switch exchange {
	case "NEW YORK STOCK EXCHANGE, INC.":
		return "NYSE"
	case "NASDAQ NMS - GLOBAL MARKET":
		return "NASDAQ"
	default:
		return exchange
	}
}

func BuildStockCard(stock *domain.StockSummary) flexmessage.Component {
	color, arrow := priceChangeStyle(stock.PriceChange)

	header := flexmessage.NewBox("vertical", []flexmessage.Component{
		flexmessage.NewBox("horizontal", []flexmessage.Component{
			flexmessage.NewText(
				stock.Ticker,
				flexmessage.WithTextWeight("bold"),
				flexmessage.WithTextSize("3xl"),
				flexmessage.WithTextColor("#ffffff"),
				flexmessage.WithTextFlex(0),
			),
			flexmessage.NewText(
				formatExchange(stock.Exchange),
				flexmessage.WithTextColor("#ffffff99"),
				flexmessage.WithTextSize("sm"),
				flexmessage.WithTextAlign("end"),
				flexmessage.WithTextGravity("bottom"),
			),
		}),
		flexmessage.NewText(
			stock.Name,
			flexmessage.WithTextColor("#ffffff99"),
			flexmessage.WithTextSize("xs"),
			flexmessage.WithTextMargin("sm"),
		),
	},
		flexmessage.WithBoxBackgroundColor("#181A20"),
		flexmessage.WithBoxPaddingAll("20px"),
		flexmessage.WithBoxPaddingBottom("16px"),
	)

	body := flexmessage.NewBox("vertical", []flexmessage.Component{
		flexmessage.NewBox("horizontal", []flexmessage.Component{
			flexmessage.NewText(
				fmt.Sprintf("%.2f", stock.CurrentPrice),
				flexmessage.WithTextWeight("bold"),
				flexmessage.WithTextSize("4xl"),
				flexmessage.WithTextColor("#181A20"),
				flexmessage.WithTextFlex(0),
				flexmessage.WithTextAdjustMode("shrink-to-fit"),
			),

			flexmessage.NewText(
				"THB",
				flexmessage.WithTextSize("md"),
				flexmessage.WithTextColor("#8b949e"),
				flexmessage.WithTextGravity("bottom"),
				flexmessage.WithTextMargin("sm"),
				flexmessage.WithTextFlex(0),
			),
		}),

		flexmessage.NewBox("horizontal", []flexmessage.Component{
			flexmessage.NewText(
				arrow,
				flexmessage.WithTextSize("sm"),
				flexmessage.WithTextColor(color),
				flexmessage.WithTextFlex(0),
				flexmessage.WithTextMargin("none"),
			),

			flexmessage.NewText(
				fmt.Sprintf("%.2f (%.2f%%)", stock.PriceChange, stock.PercentChange),
				flexmessage.WithTextSize("lg"),
				flexmessage.WithTextColor(color),
				flexmessage.WithTextWeight("bold"),
				flexmessage.WithTextMargin("md"),
				flexmessage.WithTextFlex(0),
			),

			flexmessage.NewText("Today",
				flexmessage.WithTextSize("xs"),
				flexmessage.WithTextColor("#8b949e"),
				flexmessage.WithTextAlign("end"),
				flexmessage.WithTextGravity("center"),
			),
		},
			flexmessage.WithBoxMargin("md"),
			flexmessage.WithBoxAlignItems("center"),
		),

		flexmessage.NewSeperator(flexmessage.WithSeparatorMargin("xl"), flexmessage.WithSeparatorColor("#e1e4e8")),

		flexmessage.NewBox("vertical", []flexmessage.Component{
			flexmessage.NewBox("horizontal", []flexmessage.Component{
				flexmessage.NewText(
					"High",
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#8b949e"),
				),

				flexmessage.NewText(
					fmt.Sprintf("%.2f", stock.HighPriceOfDay),
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#181A20"),
					flexmessage.WithTextAlign("end"),
					flexmessage.WithTextWeight("bold"),
				),
			}),

			flexmessage.NewBox("horizontal", []flexmessage.Component{
				flexmessage.NewText(
					"Low",
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#8b949e"),
				),

				flexmessage.NewText(
					fmt.Sprintf("%.2f", stock.LowPriceOfDay),
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#181A20"),
					flexmessage.WithTextAlign("end"),
					flexmessage.WithTextWeight("bold"),
				),
			}),

			flexmessage.NewBox("horizontal", []flexmessage.Component{
				flexmessage.NewText("Open",
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#8b949e"),
				),

				flexmessage.NewText(fmt.Sprintf("%.2f", stock.OpenPriceOfDay),
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#181A20"),
					flexmessage.WithTextAlign("end"),
					flexmessage.WithTextWeight("bold"),
				),
			}),

			flexmessage.NewBox("horizontal", []flexmessage.Component{
				flexmessage.NewText("Prev Close",
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#8b949e"),
				),

				flexmessage.NewText(fmt.Sprintf("%.2f", stock.PreviousClosePrice),
					flexmessage.WithTextSize("sm"),
					flexmessage.WithTextColor("#181A20"),
					flexmessage.WithTextAlign("end"),
					flexmessage.WithTextWeight("bold"),
				),
			}),
		},
			flexmessage.WithBoxMargin("xl"),
			flexmessage.WithBoxSpacing("sm"),
		),
	}, flexmessage.WithBoxPaddingAll("20px"))

	footer := flexmessage.NewBox("horizontal", []flexmessage.Component{
		flexmessage.NewButton(
			flexmessage.NewURIAction("View Chart", "https://line.me/"),
			flexmessage.WithButtonStyle("primary"),
			flexmessage.WithButtonColor("#181A20"),
			flexmessage.WithButtonHeight("sm"),
		),
	})

	return flexmessage.NewBubble(
		flexmessage.WithBubbleHeader(header),
		flexmessage.WithBubbleBody(body),
		flexmessage.WithBubbleFooter(footer),
		flexmessage.WithBubbleStyles(
			flexmessage.Component{
				"footer": flexmessage.Component{
					"separator": false,
				},
			},
		),
	)
}
