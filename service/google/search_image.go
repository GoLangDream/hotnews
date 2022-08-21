package google

import (
	"context"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gookit/config/v2"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func GetSearchImage(query string) string {
	apiKey := config.String("application.google.search.apiKey")
	cx := config.String("application.google.search.cx")

	ctx := context.Background()

	customSearchService, err := customsearch.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Info(err)
		return ""
	}

	resp, err := customSearchService.Cse.List().Cx(cx).Q(query).ImgSize("XLARGE").SearchType("image").Do()
	if err != nil {
		log.Info(err)
		return ""
	}

	if len(resp.Items) > 0 {
		return resp.Items[0].Link
	}

	return ""
}
