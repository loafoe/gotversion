package gotversion

import (
	"encoding/json"
	"fmt"
)

type outputJSON struct {
	Major             int    `json:"Major"`
	Minor             int    `json:"Minor"`
	Patch             int    `json:"Patch"`
	SemVer            string `json:"SemVer,omitempty"`
	FullSemVer        string `json:"FullSemVer,omitempty"`
	MajorMinorPatch   string `json:"MajorMinorPatch,omitempty"`
	PreReleaseLabel   string `json:"PreReleaseLabel,omitempty"`
	BranchName        string `json:"BranchName,omitempty"`
	SHA               string `json:"Sha,omitempty"`
	CommitDate        string `json:"CommitDate,omitempty"`
	FullBuildMetaData string `json:"FullBuildMetaData,omitempty"`
}

// OutputJSON outputs gotversion results in JSON
func OutputJSON(base *Base) error {
	data := &outputJSON{
		Major:             base.Major(),
		Minor:             base.Minor(),
		Patch:             base.Patch(),
		SemVer:            base.Semver(),
		FullSemVer:        base.FullSemver(),
		SHA:               base.Commit(),
		CommitDate:        base.CommitDate(),
		MajorMinorPatch:   base.MajorMinorPatch(),
		BranchName:        base.BranchName(),
		FullBuildMetaData: base.FullBuildMetaData(),
		PreReleaseLabel:   base.PreReleaseLabel(),
	}
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
