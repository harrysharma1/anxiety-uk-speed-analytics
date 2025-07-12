package cmd

type PageSpeedResponse struct {
	CaptchaResult           string           `json:"captchaResult"`
	Kind                    string           `json:"kind"`
	ID                      string           `json:"id"`
	LoadingExperience       Experience       `json:"loadingExperience"`
	OriginLoadingExperience Experience       `json:"originLoadingExperience"`
	LighthouseResult        LighthouseResult `json:"lighthouseResult"`
	AnalysisUTCTimeStamp    string           `json:"analysisUTCTimestamp"`
	Version                 Version          `json:"version"`
}

type Experience struct {
	ID              string            `json:"id"`
	Metrics         map[string]Metric `json:"metrics"`
	OverallCategory string            `json:"overall_category"`
	InitialURL      string            `json:"initial_url"`
}

type Metric struct {
	Percentile    int            `json:"percentile"`
	Distributions []Distribution `json:"distributions"`
	Category      string         `json:"category"`
}

type Distribution struct {
	Min        int     `json:"min"`
	Max        int     `json:"max"`
	Proportion float64 `json:"proportion"`
}

type LighthouseResult struct {
	RequestedURL      string                    `json:"requestedUrl"`
	FinalURL          string                    `json:"finalUrl"`
	LightHouseVersion string                    `json:"lighthouseVersion"`
	UserAgent         string                    `json:"userAgent"`
	FetchTime         string                    `json:"fetchTime"`
	Environment       Environment               `json:"environment"`
	RunWarnings       []interface{}             `json:"runWarnings"`
	ConfigSettings    ConfigSettings            `json:"configSettings"`
	Audits            map[string]Audits         `json:"audits"`
	Categories        map[string]Categories     `json:"categories"`
	CategoryGroups    map[string]CategoryGroups `json:"categoryGroups"`
	RuntimeError      RuntimeError              `json:"runtimeError"`
	Timing            Timing                    `json:"timing"`
	I18n              I18n                      `json:"i18n"`
}

type Environment struct {
	NetworkUserAgent string  `json:"networkUserAgent"`
	HostUserAgent    string  `json:"hostUserAgent"`
	BenchmarkIndex   float64 `json:"benchmarkIndex"`
}

type ConfigSettings struct {
	EmulatedFormFactor string      `json:"emulatedFormFactor"`
	Locale             string      `json:"locale"`
	OnlyCategories     interface{} `json:"onlyCategories"`
}

type Audits struct {
	ID               string                 `json:"id"`
	Title            string                 `json:"title"`
	Description      string                 `json:"description"`
	Score            interface{}            `json:"score"`
	ScoreDisplayMode string                 `json:"scoreDisplayMode"`
	DisplayValue     string                 `json:"displayValue"`
	Explanation      string                 `json:"explanation"`
	ErrorMessage     string                 `json:"errorMessage"`
	Warnings         interface{}            `json:"warnings"`
	Details          map[string]interface{} `json:"details"`
}

type Categories struct {
	ID                string      `json:"id"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	Score             interface{} `json:"score"`
	ManualDescription string      `json:"string"`
	AuditRefs         []AuditRef  `json:"auditRefs"`
}

type AuditRef struct {
	ID     string  `json:"id"`
	Weight float64 `json:"weight"`
	Group  string  `json:"group"`
}

type CategoryGroups struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RuntimeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Timing struct {
	Total float64 `json:"total"`
}

type I18n struct {
	RendererFormattedStrings RendererFormattedStrings `json:"rendererFormattedStrings"`
}

type RendererFormattedStrings struct {
	VarianceDisclaimer               string `json:"varianceDisclaimer"`
	OpportunityResourceColumnLabel   string `json:"opportunityResourceColumnLabel"`
	OpportunitySavingsColumnLabel    string `json:"opportunitySavingsColumnLabel"`
	ErrorMissingAuditInfo            string `json:"errorMissingAuditInfo"`
	ErrorLabel                       string `json:"errorLabel"`
	WarningHeader                    string `json:"warningHeader"`
	AuditGroupExpanTooltip           string `json:"auditGroupExpandTooltip"`
	PassedAuditsGroupTitle           string `json:"passedAuditsGroupTitle"`
	NotApplicableAuditsGroupTitle    string `json:"notApplicableAuditsGroupTitle"`
	ManualAuditsGroupTitle           string `json:"manualAuditsGroupTitle"`
	TopLevelWarningsMessage          string `json:"toplevelWarningsMessage"`
	ScorescaleLabel                  string `json:"scorescaleLabel"`
	CrcLongestDurationLabel          string `json:"crcLongestDurationLabel"`
	CrcInitialNavigation             string `json:"crcInitialNavigation"`
	LsPerformanceCategoryDescription string `json:"lsPerformanceCategoryDescription"`
	LabDataTitle                     string `json:"labDataTitle"`
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
}
