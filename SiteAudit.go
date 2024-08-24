package semrush

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type SiteAudit struct {
	Id                 int            `json:"id"`
	Name               string         `json:"name"`
	Url                string         `json:"url"`
	ProjectDomain      string         `json:"projectDomain"`
	Status             string         `json:"status"`
	Scheme             string         `json:"scheme"`
	Errors             int            `json:"errors"`
	Warnings           int            `json:"warnings"`
	Notices            int            `json:"notices"`
	Broken             int            `json:"broken"`
	BrokenDelta        int            `json:"brokenDelta"`
	Blocked            int            `json:"blocked"`
	BlockedDelta       int            `json:"blockedDelta"`
	Redirected         int            `json:"redirected"`
	RedirectedDelta    int            `json:"redirectedDelta"`
	Healthy            int            `json:"healthy"`
	HealthyDelta       int            `json:"healthyDelta"`
	HaveIssues         int            `json:"haveIssues"`
	HaveIssuesDelta    int            `json:"haveIssuesDelta"`
	ExcludedBetaChecks []interface{}  `json:"excludedBetaChecks"`
	Defects            map[string]int `json:"defects"`
	Markups            struct {
		TwitterCard     int `json:"twitterCard"`
		SchemaOrg       int `json:"schemaOrg"`
		NoMarkups       int `json:"noMarkups"`
		HasMarkups      int `json:"hasMarkups"`
		OpenGraph       int `json:"openGraph"`
		SchemaOrgJsonld int `json:"schemaOrgJsonld"`
		Microfomats     int `json:"microfomats"`
	} `json:"markups"`
	StructuredData struct {
		GroupByPages struct {
			NoItems int `json:"noItems"`
			Valid   int `json:"valid"`
			Invalid int `json:"invalid"`
		} `json:"groupByPages"`
		GroupByItems struct {
			Valid   int `json:"valid"`
			Invalid int `json:"invalid"`
		} `json:"groupByItems"`
		Items []struct {
			Type    string `json:"type"`
			Valid   int    `json:"valid"`
			Invalid int    `json:"invalid"`
		} `json:"items"`
	} `json:"structuredData"`
	Depths                          map[string]int `json:"depths"`
	CrawlSubdomains                 bool           `json:"crawlSubdomains"`
	RespectCrawlDelay               bool           `json:"respectCrawlDelay"`
	InternalIncomingLinksGroups     map[string]int `json:"internalIncomingLinksGroups"`
	InternalIncomingLinksGroupsList []struct {
		Id          int `json:"id"`
		LeftBorder  int `json:"leftBorder"`
		RightBorder int `json:"rightBorder"`
		Count       int `json:"count"`
	} `json:"internalIncomingLinksGroupsList"`
	Canonical    int `json:"canonical"`
	WithAmp      int `json:"withAmp"`
	ScheduleDay  int `json:"scheduleDay"`
	Crawlability struct {
		Budget struct {
			Total int `json:"total"`
			Stats struct {
				NonCanonicalPages      int `json:"nonCanonicalPages"`
				PermanentRedirects     int `json:"permanentRedirects"`
				RedirectChainsAndLoops int `json:"redirectChainsAndLoops"`
				TemporaryRedirects     int `json:"temporaryRedirects"`
				ContentDuplicates      int `json:"contentDuplicates"`
				TooLargePages          int `json:"tooLargePages"`
				TooSlowPages           int `json:"tooSlowPages"`
				BotBlocked             int `json:"botBlocked"`
				Code4Xx                int `json:"code4xx"`
				Code5Xx                int `json:"code5xx"`
			} `json:"stats"`
		} `json:"budget"`
	} `json:"crawlability"`
	CrawlSourceType             int           `json:"crawlSourceType"`
	GaStatus                    string        `json:"gaStatus"`
	GaSettings                  interface{}   `json:"gaSettings"`
	ProductGroup                string        `json:"productGroup"`
	IsPaid                      bool          `json:"isPaid"`
	CrawlDelaySettingTypeId     int           `json:"crawlDelaySettingTypeId"`
	IgnoreRobotsDisallow        bool          `json:"ignoreRobotsDisallow"`
	CampaignsCount              int           `json:"campaignsCount"`
	ProjectsCount               int           `json:"projectsCount"`
	JsRendering                 string        `json:"jsRendering"`
	JsImpactReportDataAvailable bool          `json:"jsImpactReportDataAvailable"`
	UserAgentType               int           `json:"user_agent_type"`
	LastAudit                   int64         `json:"last_audit"`
	LastFailedAudit             int           `json:"last_failed_audit"`
	NextAudit                   int64         `json:"next_audit"`
	RunningPagesCrawled         int           `json:"running_pages_crawled"`
	RunningPagesLimit           int           `json:"running_pages_limit"`
	PagesCrawled                int           `json:"pages_crawled"`
	PagesCrawledDelta           int           `json:"pages_crawled_delta"`
	PagesLimit                  int           `json:"pages_limit"`
	TotalChecks                 int           `json:"total_checks"`
	ErrorsDelta                 int           `json:"errors_delta"`
	WarningsDelta               int           `json:"warnings_delta"`
	NoticesDelta                int           `json:"notices_delta"`
	MaskAllow                   []interface{} `json:"mask_allow"`
	MaskDisallow                []interface{} `json:"mask_disallow"`
	RemovedParameters           []interface{} `json:"removedParameters"`
	ExcludedChecks              []interface{} `json:"excluded_checks"`
	CurrentSnapshot             struct {
		IsPagesLimitReached bool `json:"isPagesLimitReached"`
		ThematicScores      struct {
			Https        ValueDelta `json:"https"`
			IntSeo       ValueDelta `json:"intSeo"`
			Crawlability ValueDelta `json:"crawlability"`
			Performance  ValueDelta `json:"performance"`
			Linking      ValueDelta `json:"linking"`
			Markups      ValueDelta `json:"markups"`
		} `json:"thematicScores"`
		Errors    []Alert `json:"errors"`
		Warnings  []Alert `json:"warnings"`
		Notices   []Alert `json:"notices"`
		TopIssues []int   `json:"topIssues"`
		Hidden    struct {
		} `json:"hidden"`
		Fixed            map[string]int `json:"fixed"`
		StatusCodeGroups map[string]int `json:"statusCodeGroups"`
		ExcludedChecks   []interface{}  `json:"excludedChecks"`
		SettingsChanged  bool           `json:"settingsChanged"`
		JsRendering      string         `json:"jsRendering"`
		SnapshotId       string         `json:"snapshot_id"`
		SnapshotVersion  int            `json:"snapshot_version"`
		PagesCrawled     int            `json:"pages_crawled"`
		PagesLimit       int            `json:"pages_limit"`
		FinishDate       int64          `json:"finish_date"`
		Quality          struct {
			Value int `json:"value"`
			Delta int `json:"delta"`
		} `json:"quality"`
		UserAgentType int            `json:"user_agent_type"`
		New           map[string]int `json:"new"`
	} `json:"current_snapshot"`
	IsNotify  bool `json:"is_notify"`
	AmpTotal  int  `json:"ampTotal"`
	Hreflangs struct {
		TotalPages int `json:"totalPages"`
		BadPages   int `json:"badPages"`
		Missing    int `json:"missing"`
	} `json:"hreflangs"`
	Xdefault int `json:"xdefault"`
	Sitemaps struct {
		CrawledPages int    `json:"crawledPages"`
		TotalPages   int    `json:"totalPages"`
		Count        int    `json:"count"`
		ReprUrl      string `json:"reprUrl"`
		Status       string `json:"status"`
	} `json:"sitemaps"`
	Canonicalization struct {
		EmptyOrWithoutCanonicalLink int `json:"emptyOrWithoutCanonicalLink"`
		SelfCanonical               int `json:"selfCanonical"`
		CanonicalToOtherPage        int `json:"canonicalToOtherPage"`
	} `json:"canonicalization"`
	UploadUrlsCount   int         `json:"uploadUrlsCount"`
	CrawlSourceUrl    string      `json:"crawlSourceUrl"`
	PrecrawlError     interface{} `json:"precrawlError"`
	IsAllowedStopping bool        `json:"is_allowed_stopping"`
	Robots            struct {
		Status  int    `json:"status"`
		Link    string `json:"link"`
		Errors  int    `json:"errors"`
		Changes int    `json:"changes"`
	} `json:"robots"`
	AuthEnabled bool `json:"authEnabled"`
}

type ValueDelta struct {
	Value int `json:"value"`
	Delta int `json:"delta"`
}

type Alert struct {
	Id     int `json:"id"`
	Count  int `json:"count"`
	Delta  int `json:"delta"`
	Checks int `json:"checks"`
}

// GetSiteAudit returns all siteAudits
func (service *Service) GetSiteAudit(projectId int) (*SiteAudit, *errortools.Error) {
	var siteAudits SiteAudit

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.urlReports(fmt.Sprintf("projects/%v/siteaudit/info", projectId)),
		ResponseModel: &siteAudits,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &siteAudits, nil
}
