package passes

import (
	"github.com/bflad/tfproviderlint/passes/AT001"
	"github.com/bflad/tfproviderlint/passes/AT002"
	"github.com/bflad/tfproviderlint/passes/AT003"
	"github.com/bflad/tfproviderlint/passes/AT004"
	"github.com/bflad/tfproviderlint/passes/AT005"
	"github.com/bflad/tfproviderlint/passes/AT006"
	"github.com/bflad/tfproviderlint/passes/AT007"
	"github.com/bflad/tfproviderlint/passes/AT008"
	"github.com/bflad/tfproviderlint/passes/R001"
	"github.com/bflad/tfproviderlint/passes/R002"
	"github.com/bflad/tfproviderlint/passes/R003"
	"github.com/bflad/tfproviderlint/passes/R004"
	"github.com/bflad/tfproviderlint/passes/R005"
	"github.com/bflad/tfproviderlint/passes/R006"
	"github.com/bflad/tfproviderlint/passes/R007"
	"github.com/bflad/tfproviderlint/passes/R008"
	"github.com/bflad/tfproviderlint/passes/R009"
	"github.com/bflad/tfproviderlint/passes/R010"
	"github.com/bflad/tfproviderlint/passes/R011"
	"github.com/bflad/tfproviderlint/passes/R012"
	"github.com/bflad/tfproviderlint/passes/R013"
	"github.com/bflad/tfproviderlint/passes/R014"
	"github.com/bflad/tfproviderlint/passes/S001"
	"github.com/bflad/tfproviderlint/passes/S002"
	"github.com/bflad/tfproviderlint/passes/S003"
	"github.com/bflad/tfproviderlint/passes/S004"
	"github.com/bflad/tfproviderlint/passes/S005"
	"github.com/bflad/tfproviderlint/passes/S006"
	"github.com/bflad/tfproviderlint/passes/S007"
	"github.com/bflad/tfproviderlint/passes/S008"
	"github.com/bflad/tfproviderlint/passes/S009"
	"github.com/bflad/tfproviderlint/passes/S010"
	"github.com/bflad/tfproviderlint/passes/S011"
	"github.com/bflad/tfproviderlint/passes/S012"
	"github.com/bflad/tfproviderlint/passes/S013"
	"github.com/bflad/tfproviderlint/passes/S014"
	"github.com/bflad/tfproviderlint/passes/S015"
	"github.com/bflad/tfproviderlint/passes/S016"
	"github.com/bflad/tfproviderlint/passes/S017"
	"github.com/bflad/tfproviderlint/passes/S018"
	"github.com/bflad/tfproviderlint/passes/S019"
	"github.com/bflad/tfproviderlint/passes/S020"
	"github.com/bflad/tfproviderlint/passes/S021"
	"github.com/bflad/tfproviderlint/passes/S022"
	"github.com/bflad/tfproviderlint/passes/S023"
	"github.com/bflad/tfproviderlint/passes/S024"
	"github.com/bflad/tfproviderlint/passes/S025"
	"github.com/bflad/tfproviderlint/passes/S026"
	"github.com/bflad/tfproviderlint/passes/S027"
	"github.com/bflad/tfproviderlint/passes/S028"
	"github.com/bflad/tfproviderlint/passes/S029"
	"github.com/bflad/tfproviderlint/passes/S030"
	"github.com/bflad/tfproviderlint/passes/S031"
	"github.com/bflad/tfproviderlint/passes/S032"
	"github.com/bflad/tfproviderlint/passes/S033"
	"github.com/bflad/tfproviderlint/passes/S034"
	"github.com/bflad/tfproviderlint/passes/S035"
	"github.com/bflad/tfproviderlint/passes/S036"
	"github.com/bflad/tfproviderlint/passes/S037"
	"github.com/bflad/tfproviderlint/passes/S038"
	"github.com/bflad/tfproviderlint/passes/V001"
	"github.com/bflad/tfproviderlint/passes/V002"
	"github.com/bflad/tfproviderlint/passes/V003"
	"github.com/bflad/tfproviderlint/passes/V004"
	"github.com/bflad/tfproviderlint/passes/V005"
	"github.com/bflad/tfproviderlint/passes/V006"
	"github.com/bflad/tfproviderlint/passes/V007"
	"github.com/bflad/tfproviderlint/passes/V008"
	"golang.org/x/tools/go/analysis"
)

// AllChecks contains all Analyzers that report issues
// This can be consumed via multichecker.Main(passes.AllChecks...) or by
// combining these Analyzers with additional custom Analyzers
var AllChecks = []*analysis.Analyzer{
	AT001.Analyzer,
	AT002.Analyzer,
	AT003.Analyzer,
	AT004.Analyzer,
	AT005.Analyzer,
	AT006.Analyzer,
	AT007.Analyzer,
	AT008.Analyzer,
	R001.Analyzer,
	R002.Analyzer,
	R003.Analyzer,
	R004.Analyzer,
	R005.Analyzer,
	R006.Analyzer,
	R007.Analyzer,
	R008.Analyzer,
	R009.Analyzer,
	R010.Analyzer,
	R011.Analyzer,
	R012.Analyzer,
	R013.Analyzer,
	R014.Analyzer,
	S001.Analyzer,
	S002.Analyzer,
	S003.Analyzer,
	S004.Analyzer,
	S005.Analyzer,
	S006.Analyzer,
	S007.Analyzer,
	S008.Analyzer,
	S009.Analyzer,
	S010.Analyzer,
	S011.Analyzer,
	S012.Analyzer,
	S013.Analyzer,
	S014.Analyzer,
	S015.Analyzer,
	S016.Analyzer,
	S017.Analyzer,
	S018.Analyzer,
	S019.Analyzer,
	S020.Analyzer,
	S021.Analyzer,
	S022.Analyzer,
	S023.Analyzer,
	S024.Analyzer,
	S025.Analyzer,
	S026.Analyzer,
	S027.Analyzer,
	S028.Analyzer,
	S029.Analyzer,
	S030.Analyzer,
	S031.Analyzer,
	S032.Analyzer,
	S033.Analyzer,
	S034.Analyzer,
	S035.Analyzer,
	S036.Analyzer,
	S037.Analyzer,
	S038.Analyzer,
	V001.Analyzer,
	V002.Analyzer,
	V003.Analyzer,
	V004.Analyzer,
	V005.Analyzer,
	V006.Analyzer,
	V007.Analyzer,
	V008.Analyzer,
}
