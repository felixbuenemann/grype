package java

import (
	"fmt"
	search2 "github.com/anchore/grype/grype/db/search"
	"net/http"

	"github.com/anchore/grype/grype/distro"
	"github.com/anchore/grype/grype/match"
	"github.com/anchore/grype/grype/pkg"
	"github.com/anchore/grype/grype/vulnerability"
	"github.com/anchore/grype/internal/log"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

const (
	sha1Query = `1:"%s"`
)

type Matcher struct {
	MavenSearcher
	cfg MatcherConfig
}

type ExternalSearchConfig struct {
	SearchMavenUpstream bool
	MavenBaseURL        string
}

type MatcherConfig struct {
	ExternalSearchConfig
	UseCPEs bool
}

func NewJavaMatcher(cfg MatcherConfig) *Matcher {
	return &Matcher{
		cfg: cfg,
		MavenSearcher: &mavenSearch{
			client:  http.DefaultClient,
			baseURL: cfg.MavenBaseURL,
		},
	}
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.JavaPkg, syftPkg.JenkinsPluginPkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.JavaMatcher
}

func (m *Matcher) Match(store vulnerability.Provider, d *distro.Distro, p pkg.Package) ([]match.Match, error) {
	var matches []match.Match
	if m.cfg.SearchMavenUpstream {
		upstreamMatches, err := m.matchUpstreamMavenPackages(store, d, p)
		if err != nil {
			log.Debugf("failed to match against upstream data for %s: %v", p.Name, err)
		} else {
			matches = append(matches, upstreamMatches...)
		}
	}
	criteria := search2.CommonCriteria
	if m.cfg.UseCPEs {
		criteria = append(criteria, search2.ByCPE)
	}
	criteriaMatches, err := search2.ByCriteria(store, d, p, m.Type(), criteria...)
	if err != nil {
		return nil, fmt.Errorf("failed to match by exact package: %w", err)
	}

	matches = append(matches, criteriaMatches...)
	return matches, nil
}

func (m *Matcher) matchUpstreamMavenPackages(store vulnerability.Provider, d *distro.Distro, p pkg.Package) ([]match.Match, error) {
	var matches []match.Match

	if metadata, ok := p.Metadata.(pkg.JavaMetadata); ok {
		for _, digest := range metadata.ArchiveDigests {
			if digest.Algorithm == "sha1" {
				indirectPackage, err := m.GetMavenPackageBySha(digest.Value)
				if err != nil {
					return nil, err
				}
				indirectMatches, err := search2.ByPackageLanguage(store, d, *indirectPackage, m.Type())
				if err != nil {
					return nil, err
				}
				matches = append(matches, indirectMatches...)
			}
		}
	}

	match.ConvertToIndirectMatches(matches, p)

	return matches, nil
}
