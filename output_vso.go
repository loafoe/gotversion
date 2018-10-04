package gotversion

import "fmt"

// OutputVSO outputs variables to Azure DevOPS output
func OutputVSO(base *Base) error {
	fmt.Printf(vsoString("GotSemVer", base.Semver()))
	fmt.Printf(vsoString("GotSHA", base.Commit()))
	fmt.Printf(vsoString("GotCommitDate", base.CommitDate()))
	return nil
}

func vsoString(variable, value string) string {
	return fmt.Sprintf("##vso[task.setvariable variable=%s;isOutput=true;]%s\n", variable, value)
}

func vsoInt(variable string, value int) string {
	return fmt.Sprintf("##vso[task.setvariable variable=%s;isOutput=true;]%d\n", variable, value)
}
