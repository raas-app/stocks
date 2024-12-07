package scrapperfx

import (
	config "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/scrapper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideCompanyScrapper(cfg *config.Config, logger *zap.Logger) (scrapper.CompanyScrapper, error) {
	return scrapper.NewCompanyScrapper(cfg, logger)
}

var Providers = fx.Module("scrappers",
	fx.Provide(ProvideCompanyScrapper))
