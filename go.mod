module github.com/restic/restic

replace (
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/form3tech-oss/jwt-go => github.com/golang-jwt/jwt/v4 v4.0.0
	golang.org/x/text => golang.org/x/text v0.3.7
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
)

require (
	bazil.org/fuse v0.0.0-20200407214033-5883e5a4b512
	cloud.google.com/go/storage v1.16.0
	github.com/Azure/azure-sdk-for-go v55.6.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.21 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/aws/aws-sdk-go v1.38.21
	github.com/cenkalti/backoff/v4 v4.1.1
	github.com/cespare/xxhash/v2 v2.1.1
	github.com/dchest/siphash v1.2.2
	github.com/dnaeon/go-vcr v1.2.0 // indirect
	github.com/elithrar/simple-scrypt v1.3.0
	github.com/go-ole/go-ole v1.2.5
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.1.0 // indirect
	github.com/google/go-cmp v0.5.6
	github.com/hashicorp/golang-lru v0.5.4
	github.com/juju/ratelimit v1.0.1
	github.com/kurin/blazer v0.5.3
	github.com/minio/minio-go/v7 v7.0.12
	github.com/minio/sha256-simd v1.0.0
	github.com/ncw/swift v1.0.53
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.6.0
	github.com/pkg/sftp v1.13.2
	github.com/pkg/xattr v0.4.3
	github.com/restic/chunker v0.4.0
	github.com/smartystreets/assertions v1.2.0 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211015200801-69063c4bb744
	golang.org/x/text v0.3.7
	google.golang.org/api v0.58.0
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
)

go 1.16
