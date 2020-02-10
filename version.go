package version

import (
	"fmt"
	"runtime"
	"strings"
	"text/template"
)

// These variables are usually defined via LDFLags at compile-time
var (
	ApplicationName = "unknown"
	CommitHash      = "unknown"
	BuildDate       = "unknown"
	BuildTag        = "v0-dev"

	Template = "{{.ApplicationName}} {{.BuildTag}} ({{.CommitHash}}) {{.BuildTarget}} - BuildDate: {{.BuildDate}}"
)

var (
	t = template.Must(template.New("version").Parse(Template))
)

// Info contains all the versioning information for a package/tool
type Info struct {
	ApplicationName string `json:"application_name"`
	CommitHash      string `json:"commit_hash"`
	BuildDate       string `json:"build_date"`
	BuildTarget     string `json:"build_target"`
	BuildTag        string `json:"build_number"`
	GoVersion       string `json:"go_version"`
	GOOS            string `json:"goos"`
	GOARCH          string `json:"goarch"`
}

func (vi Info) String() string {
	s := &strings.Builder{}
	if err := t.Execute(s, vi); err != nil {
		panic(err)
	}

	return s.String()
}

// Version is the default version of the package/tool
var Version Info

func init() {
	Version = Info{
		ApplicationName: ApplicationName,
		CommitHash:      CommitHash,
		BuildDate:       BuildDate,
		BuildTag:        BuildTag,
		BuildTarget:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		GoVersion:       runtime.Version(),
		GOOS:            runtime.GOOS,
		GOARCH:          runtime.GOARCH,
	}
}

func String() string {
	return Version.String()
}
