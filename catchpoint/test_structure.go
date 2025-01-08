package catchpoint

type IdName struct {
	Id   int
	Name string
}

type TestConfig struct {
	TestType                       int
	Monitor                        int
	SimulateDevice                 int
	ChromeVersion                  IdName
	ChromeApplicationVersion       IdName
	Script                         TestRequestData
	DivisionId                     int
	ProductId                      int
	FolderId                       int
	TestName                       string
	DnsQueryType                   IdName
	DnsServer                      string
	EdnsSubnet                     string
	TestUrl                        string
	TestDescription                string
	GatewayAddressOrHost           string
	Labels                         []TestLabel
	TestTimeThresholdWarning       float64
	TestTimeThresholdCritical      float64
	AvailabilityThresholdWarning   float64
	AvailabilityThresholdCritical  float64
	EnforceCertificateKeyPinning   bool
	EnforceCertificatePinning      bool
	EnableTestDataWebhook          bool
	AlertsPaused                   bool
	StartTime                      string
	EndTime                        string
	TestStatus                     int
	RequestSettingType             int
	AuthenticationType             IdName
	AuthenticationPasswordIds      []int
	AuthenticationTokenIds         []int
	AuthenticationCertificateIds   []int
	TestHttpHeaderRequests         []TestHttpHeaderRequest
	InsightSettingType             int
	TracepointIds                  []int
	IndicatorIds                   []int
	ScheduleSettingType            int
	ScheduleRunScheduleId          int
	ScheduleMaintenanceScheduleId  int
	TestFrequency                  IdName
	NodeDistribution               IdName
	NodeIds                        []int
	NodeGroupIds                   []NodeGroup
	NoOfSubsetNodes                int
	AlertSettingType               IdName
	AlertRuleConfigs               []AlertRuleConfig
	AlertWebhookIds                []int
	AlertRecipientEmails           []string
	AlertContactGroups             []string
	AdvancedSettingType            int
	AppliedTestFlags               []int
	MaxStepRuntimeSecOverride      int
	AdditionalMonitorType          IdName
	BandwidthThrottling            IdName
	WaitForNoActivityOnDocComplete *int
	ViewportHeight                 int
	ViewportWidth                  int
	TracerouteFailureHopCount      int
	TraceroutePingCount            int
	AlertSubject                   string
}

type ProductConfig struct {
	DivisionId                     int
	ProductName                    string
	ProductStatus                  int
	TestDataWebhookId              int
	EdnsSubnet                     string
	AlertGroupId                   int
	RequestSettingType             int
	AuthenticationType             IdName
	AuthenticationPasswordIds      []int
	AuthenticationTokenIds         []int
	AuthenticationCertificateIds   []int
	TestHttpHeaderRequests         []TestHttpHeaderRequest
	InsightSettingType             int
	TracepointIds                  []int
	IndicatorIds                   []int
	ScheduleSettingType            int
	ScheduleRunScheduleId          int
	ScheduleMaintenanceScheduleId  int
	TestFrequency                  IdName
	NodeDistribution               IdName
	NodeIds                        []int
	NodeGroupIds                   []NodeGroup
	NoOfSubsetNodes                int
	AlertSettingType               int
	AlertRuleConfigs               []AlertRuleConfig
	AlertWebhookIds                []int
	AlertRecipientEmails           []string
	AlertContactGroups             []string
	AdvancedSettingType            int
	AppliedTestFlags               []int
	MaxStepRuntimeSecOverride      int
	AdditionalMonitorType          IdName
	BandwidthThrottling            IdName
	WaitForNoActivityOnDocComplete *int
	ViewportHeight                 int
	ViewportWidth                  int
	TracerouteFailureHopCount      int
	TraceroutePingCount            int
	VerifytestOnFailure            bool
	AlertSubject                   string
}

type FolderConfig struct {
	DivisionId                     int
	ProductId                      int
	ParentId                       int
	FolderName                     string
	EdnsSubnet                     string
	RequestSettingType             int
	AuthenticationType             IdName
	AuthenticationPasswordIds      []int
	AuthenticationTokenIds         []int
	AuthenticationCertificateIds   []int
	TestHttpHeaderRequests         []TestHttpHeaderRequest
	InsightSettingType             int
	TracepointIds                  []int
	IndicatorIds                   []int
	ScheduleSettingType            int
	ScheduleRunScheduleId          int
	ScheduleMaintenanceScheduleId  int
	TestFrequency                  IdName
	NodeDistribution               IdName
	NodeIds                        []int
	NodeGroupIds                   []NodeGroup
	NoOfSubsetNodes                int
	AlertSettingType               IdName
	AlertRuleConfigs               []AlertRuleConfig
	AlertWebhookIds                []int
	AlertRecipientEmails           []string
	AlertContactGroups             []string
	AdvancedSettingType            int
	AppliedTestFlags               []int
	MaxStepRuntimeSecOverride      int
	AdditionalMonitorType          IdName
	BandwidthThrottling            IdName
	WaitForNoActivityOnDocComplete *int
	ViewportHeight                 int
	ViewportWidth                  int
	TracerouteFailureHopCount      int
	TraceroutePingCount            int
	VerifytestOnFailure            bool
	AlertSubject                   string
}

type TestHttpHeaderRequest struct {
	RequestHeaderType IdName
	RequestValue      string
	ChildHostPattern  string
}

type TestRequestData struct {
	TestId                int
	RequestData           string
	TransactionScriptType int
	TestType              int
	Monitor               int
}

type AlertRuleConfig struct {
	AlertNodeThresholdType          IdName
	AlertThresholdNumOfRuns         int
	AlertConsecutiveNumOfRuns       int
	AlertThresholdPercentOfRuns     float64
	AlertThresholdNumOfFailingNodes int
	TriggerType                     IdName
	OperationType                   IdName
	StatisticalType                 IdName
	TrailingHistoricalInterval      IdName
	Expression                      string
	AlertWarningTrigger             float64
	AlertCriticalTrigger            float64
	AlertEnableConsecutive          bool
	AlertWarningReminder            IdName
	AlertCriticalReminder           IdName
	AlertThresholdInterval          IdName
	AlertUseRollingWindow           bool
	AlertNotificationType           int
	AlertType                       IdName
	AlertSubType                    IdName
	AlertEnforceTestFailure         bool
	AlertOmitScatterplot            bool
	DnsResolvedName                 string
	DnsTTL                          int
	DnsRecordType                   IdName
	AllMatchRecords                 bool
	FilterType                      IdName
	FilterValue                     string
	NotificationGroups              []NotificationGroupStruct
}

type TestLabel struct {
	Name   string
	Values []string
}

type TestConfigUpdate struct {
	UpdatedFieldValue              string
	UpdatedAdvancedSettingsSection AdvancedSetting
	UpdatedRequestSettingsSection  RequestSetting
	UpdatedScheduleSettingsSection ScheduleSetting
	UpdatedInsightSettingsSection  InsightDataStruct
	UpdatedAlertSettingsSection    AlertGroupStruct
	UpdatedLabels                  []Label
	UpdatedTestThresholds          Thresholds
	UpdatedTestRequestData         TestRequestDataStruct
	SectionToUpdate                string
}

type ProductConfigUpdate struct {
	UpdatedFieldValue              string
	UpdatedAdvancedSettingsSection AdvancedSetting
	UpdatedRequestSettingsSection  RequestSetting
	UpdatedScheduleSettingsSection interface{}
	UpdatedInsightSettingsSection  []map[string]int
	UpdatedAlertSettingsSection    AlertGroupStruct
	SectionToUpdate                string
}

type FolderConfigUpdate struct {
	UpdatedFieldValue              string
	UpdatedAdvancedSettingsSection AdvancedSetting
	UpdatedRequestSettingsSection  RequestSetting
	UpdatedScheduleSettingsSection interface{}
	UpdatedInsightSettingsSection  []map[string]int
	UpdatedAlertSettingsSection    AlertGroupStruct
	SectionToUpdate                string
}
