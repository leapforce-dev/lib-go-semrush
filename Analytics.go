package semrush

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Analytics map[string]int

type AnalyticsType string

const (
	AnalyticsTypeBacklinksOverview      AnalyticsType = "backlinks_overview"
	AnalyticsTypeBacklinksAscoreProfile AnalyticsType = "backlinks_ascore_profile"
)

type AnalyticsTargetType string

const (
	AnalyticsTargetTypeRootDomain AnalyticsTargetType = "root_domain"
	AnalyticsTargetTypeDomain     AnalyticsTargetType = "domain"
	AnalyticsTargetTypeUrl        AnalyticsTargetType = "url"
)

type AnalyticsExportColumn string

const (
	AnalyticsExportColumnAscore       AnalyticsExportColumn = "ascore"
	AnalyticsExportColumnTotal        AnalyticsExportColumn = "total"
	AnalyticsExportColumnDomainsNum   AnalyticsExportColumn = "domains_num"
	AnalyticsExportColumnUrlsNum      AnalyticsExportColumn = "urls_num"
	AnalyticsExportColumnIpsNum       AnalyticsExportColumn = "ips_num"
	AnalyticsExportColumnIpclasscNum  AnalyticsExportColumn = "ipclassc_num"
	AnalyticsExportColumnFollowsNum   AnalyticsExportColumn = "follows_num"
	AnalyticsExportColumnNofollowsNum AnalyticsExportColumn = "nofollows_num"
	AnalyticsExportColumnSponsoredNum AnalyticsExportColumn = "sponsored_num"
	AnalyticsExportColumnUgcNum       AnalyticsExportColumn = "ugc_num"
	AnalyticsExportColumnTextsNum     AnalyticsExportColumn = "texts_num"
	AnalyticsExportColumnImagesNum    AnalyticsExportColumn = "images_num"
	AnalyticsExportColumnFormsNum     AnalyticsExportColumn = "forms_num"
	AnalyticsExportColumnFramesNum    AnalyticsExportColumn = "frames_num"
)

type GetAnalyticsConfig struct {
	Type          AnalyticsType
	Target        string
	TargetType    AnalyticsTargetType
	ExportColumns *[]AnalyticsExportColumn
}

// GetAnalytics returns all analytics
func (service *Service) GetAnalytics(cfg *GetAnalyticsConfig) (*Analytics, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("GetAnalyticsConfig must not be nil")
	}
	values := url.Values{}

	values.Set("type", string(cfg.Type))
	values.Set("target", cfg.Target)
	values.Set("target_type", string(cfg.TargetType))

	if cfg.ExportColumns != nil {
		var exportColumns []string
		for _, ec := range *cfg.ExportColumns {
			exportColumns = append(exportColumns, string(ec))
		}

		values.Set("export_columns", strings.Join(exportColumns, ","))
	}

	requestConfig := go_http.RequestConfig{
		Method:     http.MethodGet,
		Url:        service.urlAnalytics(""),
		Parameters: &values,
	}

	_, response, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	lines := strings.Split(string(b), "\r\n")

	if len(lines) < 2 {
		return nil, errortools.ErrorMessage("Response contains less than two lines")
	}

	headers := strings.Split(lines[0], ";")
	values_ := strings.Split(lines[1], ";")

	if len(headers) != len(values_) {
		return nil, errortools.ErrorMessage("Number of headers and values is not equal")
	}

	var m = make(Analytics)

	for i, header := range headers {
		value, err := strconv.ParseInt(values_[i], 10, 64)
		if err != nil {
			errortools.CaptureError(err)
			continue
		}
		m[header] = int(value)
	}

	return &m, nil
}
