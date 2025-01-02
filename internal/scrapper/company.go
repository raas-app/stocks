package scrapper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	raas "github.com/raas-app/stocks"
	"go.uber.org/zap"
)

type companyScrapper struct {
	Collector *colly.Collector
	Config    *raas.Config
	Logger    *zap.Logger
}

func NewCompanyScrapper(config *raas.Config, logger *zap.Logger) (CompanyScrapper, error) {
	return &companyScrapper{
		Collector: colly.NewCollector(),
		Config:    config,
		Logger:    logger,
	}, nil
}

func InitializeCompanyScrapper(scrapper CompanyScrapper) {
	symbol := "EFERT"

	_, err := scrapper.GetCompany(symbol)
	if err != nil {
		return // TODO: handle error
	}

}

func (c *companyScrapper) GetCompany(symbol string) (*Company, error) {
	company := &Company{
		Symbol: symbol,
		Profile: CompanyProfile{
			Executives: []Executive{},
		},
		Announcements: map[ReportType][]CompanyReport{},
		Financials:    []*CompanyReport{},
	}

	c.Collector.OnHTML("#quote .company__quote .quote__details", func(e *colly.HTMLElement) {
		company.Name = e.ChildText(".quote__name")
		company.Sector = Sector{
			Name: e.ChildText(".quote__sector"),
		}

	})
	c.Collector.OnHTML("#profile", func(e *colly.HTMLElement) {
		e.ForEach(".profile__item.profile__item--decription p", func(_ int, el *colly.HTMLElement) {
			company.Profile.Description = strings.TrimSpace(el.Text)
		})

		e.ForEach(".profile__item.profile__item--people .tbl__body tr", func(_ int, el *colly.HTMLElement) {
			name := el.ChildText("td:first-child strong")
			role := el.ChildText("td:nth-child(2)")
			if name != "" && role != "" {
				company.Profile.Executives = append(company.Profile.Executives, Executive{
					Name: name,
					Role: role,
				})
			}
		})

		// Scrape Address
		e.ForEach(".profile__item .item__head", func(_ int, el *colly.HTMLElement) {
			head := strings.ToUpper(strings.TrimSpace(el.Text))
			switch head {
			case "ADDRESS":
				company.Profile.Address = strings.TrimSpace(el.DOM.Next().Text())
			case "WEBSITE":
				company.Profile.Website = el.DOM.Next().Find("a").AttrOr("href", "")
			case "FISCAL YEAR END":
				company.Profile.FiscalYearEnd = strings.TrimSpace(el.DOM.Next().Text())
			default:
				c.Logger.Warn("Unrecognized head", zap.String("head", head))
			}
		})
	})

	c.Collector.OnHTML("#announcements", func(e *colly.HTMLElement) {
		announcementTypes := [3]string{"Board Meetings", "Financial Results", "Others"}
		// Scrape Announcements
		for _, announcementType := range announcementTypes {
			e.ForEach("[data-name='"+string(announcementType)+"'] .tbl__body tr", func(_ int, h *colly.HTMLElement) {
				date, err := time.Parse("Jan 2, 2006", h.ChildText("td:first-child"))
				if err != nil {
					c.Logger.Error("Failed to parse date", zap.String("date", h.ChildText("td:first-child")), zap.Error(err))
				}
				switch announcementType {
				case "Board Meetings":
					company.Announcements[BoardMeeting] = append(company.Announcements[BoardMeeting], CompanyReport{
						Date:  date.Unix(),
						Title: h.ChildText("td:nth-child(2)"),
						URL:   h.ChildAttrs("a", "href")[1],
						Type:  BoardMeeting,
					})
				case "Financial Results":
					company.Announcements[Financial] = append(company.Announcements[Financial], CompanyReport{
						Date:  date.Unix(),
						Title: h.ChildText("td:nth-child(2)"),
						URL:   h.ChildAttrs("a", "href")[1],
						Type:  Financial,
					})
				case "Others":
					company.Announcements[Other] = append(company.Announcements[Other], CompanyReport{
						Date:  date.Unix(),
						Title: h.ChildText("td:nth-child(2)"),
						URL:   h.ChildAttrs("a", "href")[1],
						Type:  Other,
					})
				}
			})
		}
	})

	c.Collector.OnHTML("#financials", func(e *colly.HTMLElement) {
		// Scrape Financials
	})
	// collector.OnHTML("#ratios", func(e *colly.HTMLElement) {
	// 	// Scrape Ratios
	// })

	// collector.OnHTML("#payouts", func(e *colly.HTMLElement) {

	// })

	// On Error
	c.Collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Failed to scrape %s: %v\n", r.Request.URL, err)
	})

	err := c.Collector.Visit(fmt.Sprintf("%s%s/%s", c.Config.Market.PSX.BaseURL, c.Config.Market.PSX.ScraperURL.Company, symbol))
	if err != nil {
		return nil, err
	}

	company.Financials, err = c.GetFinancials(symbol)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (c *companyScrapper) GetFinancials(symbol string) ([]*CompanyReport, error) {
	reports := make([]*CompanyReport, 0)
	c.Collector.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			date, err := time.Parse("Jan 2, 2006", el.ChildText("td:first-child"))
			if err != nil {
				c.Logger.Error("Failed to parse date", zap.String("date", el.ChildText("td:first-child")), zap.Error(err))
			}

			reports = append(reports, &CompanyReport{
				Date:  date.Unix(),
				Title: el.ChildText("td:nth-child(2)"),
				URL:   el.ChildAttr("a", "href"),
				Type: func() ReportType {
					if el.ChildText("td:nth-child(1)") == string(Quaterly) {
						return Quaterly
					} else {
						return Annual
					}
				}(),
			})
		})
	})

	err := c.Collector.Visit(fmt.Sprintf("%s%s/%s", c.Config.Market.PSX.BaseURL, c.Config.Market.PSX.ScraperURL.Reports, symbol))
	if err != nil {
		return nil, err
	}
	return reports, nil
}
