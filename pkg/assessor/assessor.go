package assessor

import (
	"github.com/goodwithtech/docker-guard/pkg/assessor/contentTrust"

	"github.com/goodwithtech/docker-guard/pkg/assessor/group"
	"github.com/goodwithtech/docker-guard/pkg/assessor/manifest"
	"github.com/goodwithtech/docker-guard/pkg/assessor/passwd"
	"github.com/goodwithtech/docker-guard/pkg/assessor/user"

	"github.com/goodwithtech/docker-guard/pkg/log"
	"github.com/goodwithtech/docker-guard/pkg/types"
	"github.com/knqyf263/fanal/extractor"
)

var assessors []Assessor

type Assessor interface {
	Assess(extractor.FileMap) ([]types.Assessment, error)
	RequiredFiles() []string
}

func init() {
	RegisterAssessor(passwd.PasswdAssessor{})
	RegisterAssessor(user.UserAssessor{})
	RegisterAssessor(group.GroupAssessor{})
	RegisterAssessor(manifest.ManifestAssessor{})
	RegisterAssessor(contentTrust.ContentTrustAssessor{})
}

func GetAssessments(files extractor.FileMap) (assessments []types.Assessment) {
	for _, assessor := range assessors {
		results, err := assessor.Assess(files)
		if err != nil {
			log.Logger.Error(err)
		}
		assessments = append(assessments, results...)
	}
	return assessments
}

func RegisterAssessor(a Assessor) {
	assessors = append(assessors, a)
}

func LoadRequiredFiles() (filenames []string) {
	for _, assessor := range assessors {
		filenames = append(filenames, assessor.RequiredFiles()...)
	}
	return filenames
}