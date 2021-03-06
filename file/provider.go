package file

import (
	"fmt"

	"github.com/pivotalservices/file-downloader-resource/types"
)

//go:generate counterfeiter -o fakes/fake_provider.go provider.go Provider

// Provider - defines the interface for how to fetch configuration
type Provider interface {
	DownloadFile(targetDirectory, productSlug, version, pattern string, unpack bool) error
}

const maxRetries = 12

// FromSource - factory to return appropriate driver based on configuration
func FromSource(source types.Source) (Provider, error) {

	switch source.FileProvider {

	case types.FileProviderUnspecified, types.FileProviderPivnet:
		return NewPivnetProvider(source.PivnetToken)

	case types.FileProviderS3:
		return NewS3Provider(source.AccessKeyID, source.SecretAccessKey, source.RegionName, source.Endpoint, source.Bucket, source.SkipSSLVerification, source.DisableSSL, source.UseV2Signing)

	case types.FileProviderHTTP:
		return NewHTTPProvider(source.SkipSSLVerification, source.BaseHTTPURI)

	default:
		return nil, fmt.Errorf("unknown provider: %s", source.FileProvider)
	}
}
