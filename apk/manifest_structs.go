package apk

import "encoding/xml"

type ManifestPermission struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type ManifestQueryIntentAction struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type ManifestQueryIntentData struct {
	Text     string `xml:",chardata"`
	MimeType string `xml:"mimeType,attr"`
}

type ManifestQueryIntent struct {
	Text   string                     `xml:",chardata"`
	Action *ManifestQueryIntentAction `xml:"action"`
	Data   *ManifestQueryIntentData   `xml:"data"`
}

type ManifestQuery struct {
	Text   string               `xml:",chardata"`
	Intent *ManifestQueryIntent `xml:"intent"`
}

type ManifestFeature struct {
	Text     string `xml:",chardata"`
	Name     string `xml:"name,attr"`
	Required string `xml:"required,attr"`
}

type ManifestApplicationMetadata struct {
	Text     string `xml:",chardata"`
	Name     string `xml:"name,attr"`
	Value    string `xml:"value,attr"`
	Resource string `xml:"resource,attr"`
}

type ManifestApplicationProviderMetadata struct {
	Text                string `xml:",chardata"`
	Name                string `xml:"name,attr"`
	Resource            string `xml:"resource,attr"`
	Exported            string `xml:"exported,attr"`
	GrantUriPermissions string `xml:"grantUriPermissions,attr"`
}

type ManifestApplicationProvider struct {
	Text                string                               `xml:",chardata"`
	Authorities         string                               `xml:"authorities,attr"`
	Exported            string                               `xml:"exported,attr"`
	GrantUriPermissions string                               `xml:"grantUriPermissions,attr"`
	Name                string                               `xml:"name,attr"`
	DirectBootAware     string                               `xml:"directBootAware,attr"`
	InitOrder           string                               `xml:"initOrder,attr"`
	Multiprocess        string                               `xml:"multiprocess,attr"`
	Metadata            *ManifestApplicationProviderMetadata `xml:"meta-data"`
}

type ManifestApplicationActivityMetadata struct {
	Text     string `xml:",chardata"`
	Name     string `xml:"name,attr"`
	Resource string `xml:"resource,attr"`
}

type ManifestApplicationActivityIntentFilterAction struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type ManifestApplicationActivityIntentFilterCategory struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type ManifestApplicationActivityIntentFilterData struct {
	Text       string `xml:",chardata"`
	Scheme     string `xml:"scheme,attr"`
	Host       string `xml:"host,attr"`
	PathPrefix string `xml:"pathPrefix,attr"`
	Path       string `xml:"path,attr"`
}

type ManifestApplicationActivityIntentFilter struct {
	Text       string                                             `xml:",chardata"`
	AutoVerify string                                             `xml:"autoVerify,attr"`
	Label      string                                             `xml:"label,attr"`
	Action     *ManifestApplicationActivityIntentFilterAction     `xml:"action"`
	Category   []*ManifestApplicationActivityIntentFilterCategory `xml:"category"`
	Data       []*ManifestApplicationActivityIntentFilterData     `xml:"data"`
}

type ManifestApplicationActivity struct {
	Text                string                                     `xml:",chardata"`
	LaunchMode          string                                     `xml:"launchMode,attr"`
	Name                string                                     `xml:"name,attr"`
	NoHistory           string                                     `xml:"noHistory,attr"`
	ScreenOrientation   string                                     `xml:"screenOrientation,attr"`
	Theme               string                                     `xml:"theme,attr"`
	ConfigChanges       string                                     `xml:"configChanges,attr"`
	WindowSoftInputMode string                                     `xml:"windowSoftInputMode,attr"`
	Exported            string                                     `xml:"exported,attr"`
	ExcludeFromRecents  string                                     `xml:"excludeFromRecents,attr"`
	Enabled             string                                     `xml:"enabled,attr"`
	Process             string                                     `xml:"process,attr"`
	StateNotNeeded      string                                     `xml:"stateNotNeeded,attr"`
	ResizeableActivity  string                                     `xml:"resizeableActivity,attr"`
	SupportsRtl         string                                     `xml:"supportsRtl,attr"`
	IntentFilter        []*ManifestApplicationActivityIntentFilter `xml:"intent-filter"`
	Metadata            *ManifestApplicationActivityMetadata       `xml:"meta-data"`
}

type ManifestApplicationActivityAlias struct {
	Text              string `xml:",chardata"`
	LaunchMode        string `xml:"launchMode,attr"`
	Name              string `xml:"name,attr"`
	NoHistory         string `xml:"noHistory,attr"`
	ScreenOrientation string `xml:"screenOrientation,attr"`
	TargetActivity    string `xml:"targetActivity,attr"`
}

type ManifestApplicationServiceMetadata struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type ManifestApplicationServiceIntentFilterAction struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type ManifestApplicationServiceIntentFilter struct {
	Text     string                                        `xml:",chardata"`
	Priority string                                        `xml:"priority,attr"`
	Action   *ManifestApplicationServiceIntentFilterAction `xml:"action"`
}

type ManifestApplicationService struct {
	Text            string                                  `xml:",chardata"`
	DirectBootAware string                                  `xml:"directBootAware,attr"`
	Exported        string                                  `xml:"exported,attr"`
	Name            string                                  `xml:"name,attr"`
	Permission      string                                  `xml:"permission,attr"`
	StopWithTask    string                                  `xml:"stopWithTask,attr"`
	Enabled         string                                  `xml:"enabled,attr"`
	Metadata        []*ManifestApplicationServiceMetadata   `xml:"meta-data"`
	IntentFilter    *ManifestApplicationServiceIntentFilter `xml:"intent-filter"`
}

type ManifestApplicationReceiverIntentFilterAction struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type ManifestApplicationReceiverIntentFilter struct {
	Text   string                                         `xml:",chardata"`
	Action *ManifestApplicationReceiverIntentFilterAction `xml:"action"`
}

type ManifestApplicationReceiver struct {
	Text         string                                   `xml:",chardata"`
	Exported     string                                   `xml:"exported,attr"`
	Name         string                                   `xml:"name,attr"`
	Enabled      string                                   `xml:"enabled,attr"`
	Permission   string                                   `xml:"permission,attr"`
	IntentFilter *ManifestApplicationReceiverIntentFilter `xml:"intent-filter"`
}

type ManifestApplication struct {
	Text                string                            `xml:",chardata"`
	AllowBackup         string                            `xml:"allowBackup,attr"`
	AppComponentFactory string                            `xml:"appComponentFactory,attr"`
	ExtractNativeLibs   string                            `xml:"extractNativeLibs,attr"`
	HardwareAccelerated string                            `xml:"hardwareAccelerated,attr"`
	Icon                string                            `xml:"icon,attr"`
	IsSplitRequired     string                            `xml:"isSplitRequired,attr"`
	Label               string                            `xml:"label,attr"`
	LargeHeap           string                            `xml:"largeHeap,attr"`
	Name                string                            `xml:"name,attr"`
	RoundIcon           string                            `xml:"roundIcon,attr"`
	SupportsRtl         string                            `xml:"supportsRtl,attr"`
	Theme               string                            `xml:"theme,attr"`
	Metadata            []*ManifestApplicationMetadata    `xml:"meta-data"`
	Provider            []*ManifestApplicationProvider    `xml:"provider"`
	Activity            []*ManifestApplicationActivity    `xml:"activity"`
	ActivityAlias       *ManifestApplicationActivityAlias `xml:"activity-alias"`
	Service             []*ManifestApplicationService     `xml:"service"`
	Receiver            []*ManifestApplicationReceiver    `xml:"receiver"`
}

type Manifest struct {
	XMLName                   xml.Name              `xml:"manifest"`
	Text                      string                `xml:",chardata"`
	Android                   string                `xml:"android,attr"`
	CompileSdkVersion         string                `xml:"compileSdkVersion,attr"`
	CompileSdkVersionCodename string                `xml:"compileSdkVersionCodename,attr"`
	Package                   string                `xml:"package,attr"`
	PlatformBuildVersionCode  string                `xml:"platformBuildVersionCode,attr"`
	PlatformBuildVersionName  string                `xml:"platformBuildVersionName,attr"`
	UsesPermission            []*ManifestPermission `xml:"uses-permission"`
	Queries                   *ManifestQuery        `xml:"queries"`
	UsesFeature               []*ManifestFeature    `xml:"uses-feature"`
	Application               *ManifestApplication  `xml:"application"`
}
