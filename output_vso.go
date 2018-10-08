package gotversion

import "fmt"

// OutputVSO outputs variables to Azure DevOPS output
func OutputVSO(base *Base) error {
	fmt.Printf(vsoString("SemVer", base.Semver()))
	fmt.Printf(vsoString("FullSemVer", base.FullSemver()))
	fmt.Printf(vsoString("SHA", base.Commit()))
	fmt.Printf(vsoString("CommitDate", base.CommitDate()))
	fmt.Printf(vsoString("BranchName", base.BranchName()))
	return nil
}

func vsoString(variable, value string) string {
	return fmt.Sprintf("##vso[task.setvariable variable=%s;isOutput=true;]%s\n", variable, value)
}

func vsoInt(variable string, value int) string {
	return fmt.Sprintf("##vso[task.setvariable variable=%s;isOutput=true;]%d\n", variable, value)
}
