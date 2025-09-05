package consts

const (
	// container media types
	OCIManifestSchema1        = "application/vnd.oci.image.manifest.v1+json"
	DockerManifestSchema2     = "application/vnd.docker.distribution.manifest.v2+json"
	DockerManifestListSchema2 = "application/vnd.docker.distribution.manifest.list.v2+json"
	OCIImageIndexSchema       = "application/vnd.oci.image.index.v1+json"
	DockerConfigJSON          = "application/vnd.docker.container.image.v1+json"
	DockerLayer               = "application/vnd.docker.image.rootfs.diff.tar.gzip"
	DockerForeignLayer        = "application/vnd.docker.image.rootfs.foreign.diff.tar.gzip"
	DockerUncompressedLayer   = "application/vnd.docker.image.rootfs.diff.tar"
	OCILayer                  = "application/vnd.oci.image.layer.v1.tar+gzip"
	OCIArtifact               = "application/vnd.oci.empty.v1+json"

	// helm chart media types
	ChartConfigMediaType = "application/vnd.cncf.helm.config.v1+json"
	ChartLayerMediaType  = "application/vnd.cncf.helm.chart.content.v1.tar+gzip"
	ProvLayerMediaType   = "application/vnd.cncf.helm.chart.provenance.v1.prov"

	// file media types
	FileLayerMediaType           = "application/vnd.content.freighter.file.layer.v1"
	FileLocalConfigMediaType     = "application/vnd.content.freighter.file.local.config.v1+json"
	FileDirectoryConfigMediaType = "application/vnd.content.freighter.file.directory.config.v1+json"
	FileHttpConfigMediaType      = "application/vnd.content.freighter.file.http.config.v1+json"

	// memory media types
	MemoryConfigMediaType = "application/vnd.content.freighter.memory.config.v1+json"

	// wasm media types
	WasmArtifactLayerMediaType = "application/vnd.wasm.content.layer.v1+wasm"
	WasmConfigMediaType        = "application/vnd.wasm.config.v1+json"

	// unknown media types
	UnknownManifest = "application/vnd.freighter.cattle.io.unknown.v1+json"
	UnknownLayer    = "application/vnd.content.freighter.unknown.layer"
	Unknown         = "unknown"

	// vendor prefixes
	OCIVendorPrefix       = "vnd.oci"
	DockerVendorPrefix    = "vnd.docker"
	FreighterVendorPrefix = "vnd.freighter"

	// annotation keys
	KindAnnotationName      = "kind"
	KindAnnotationImage     = "dev.cosignproject.cosign/image"
	KindAnnotationIndex     = "dev.cosignproject.cosign/imageIndex"
	ImageAnnotationKey      = "freighter.dev/key"
	ImageAnnotationPlatform = "freighter.dev/platform"
	ImageAnnotationRegistry = "freighter.dev/registry"
	ImageAnnotationTlog     = "freighter.dev/use-tlog-verify"

	// cosign keyless validation options
	ImageAnnotationCertIdentity                 = "freighter.dev/certificate-identity"
	ImageAnnotationCertIdentityRegexp           = "freighter.dev/certificate-identity-regexp"
	ImageAnnotationCertOidcIssuer               = "freighter.dev/certificate-oidc-issuer"
	ImageAnnotationCertOidcIssuerRegexp         = "freighter.dev/certificate-oidc-issuer-regexp"
	ImageAnnotationCertGithubWorkflowRepository = "freighter.dev/certificate-github-workflow-repository"

	// content kinds
	ImagesContentKind    = "Images"
	ChartsContentKind    = "Charts"
	FilesContentKind     = "Files"
	DriverContentKind    = "Driver"
	ImageTxtsContentKind = "ImageTxts"
	ChartsCollectionKind = "ThickCharts"

	// content groups
	ContentGroup    = "content.freighter.cattle.io"
	CollectionGroup = "collection.freighter.cattle.io"

	// environment variables
	FreighterDir          = "FREIGHTER_DIR"
	FreighterTempDir      = "FREIGHTER_TEMP_DIR"
	FreighterStoreDir     = "FREIGHTER_STORE_DIR"
	FreighterIgnoreErrors = "FREIGHTER_IGNORE_ERRORS"

	// container files and directories
	ImageManifestFile = "manifest.json"
	ImageConfigFile   = "config.json"

	// other constraints
	CarbideRegistry              = "rgcrprod.azurecr.us"
	DefaultNamespace             = "freighter"
	DefaultTag                   = "latest"
	DefaultStoreName             = "store"
	DefaultFreighterDirName      = ".freighter"
	DefaultFreighterTempDirName  = "freighter"
	DefaultRegistryRootDir       = "registry"
	DefaultRegistryPort          = 5000
	DefaultFileserverRootDir     = "fileserver"
	DefaultFileserverPort        = 8080
	DefaultFileserverTimeout     = 60
	DefaultFreighterArchiveName  = "haul.tar.zst"
	DefaultFreighterManifestName = "freighter-manifest.yaml"
	DefaultRetries               = 3
	RetriesInterval              = 5
	CustomTimeFormat             = "2006-01-02 15:04:05"
)
