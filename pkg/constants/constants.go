package constants

import "time"

// Global constants
const (
	// A github token is required to run the tests. The token need to have permissions to the given github organization. By default the e2e use redhat-appstudio-qe github organization.
	GITHUB_TOKEN_ENV string = "GITHUB_TOKEN" // #nosec

	// The github organization is used to create the gitops repositories in Red Hat Appstudio.
	GITHUB_E2E_ORGANIZATION_ENV string = "MY_GITHUB_ORG" // #nosec

	// The quay organization is used to push container images using Red Hat Appstudio pipelines.
	QUAY_E2E_ORGANIZATION_ENV string = "QUAY_E2E_ORGANIZATION" // #nosec

	// The quay.io username to perform container builds and puush
	QUAY_OAUTH_USER_ENV string = "QUAY_OAUTH_USER" // #nosec

	// A quay organization where repositories for component images will be created.
	DEFAULT_QUAY_ORG_ENV string = "DEFAULT_QUAY_ORG" // #nosec

	// The quay.io token to perform container builds and push. The token must be correlated with the QUAY_OAUTH_USER environment
	QUAY_OAUTH_TOKEN_ENV string = "QUAY_OAUTH_TOKEN" // #nosec

	// The git repo url for the EC pipelines.
	EC_PIPELINES_REPO_URL_ENV string = "EC_PIPELINES_REPO_URL"

	// The git repo revision for the EC pipelines.
	EC_PIPELINES_REPO_REVISION_ENV string = "EC_PIPELINES_REPO_REVISION"

	// The private devfile sample git repository to use in certain HAS e2e tests
	PRIVATE_DEVFILE_SAMPLE string = "PRIVATE_DEVFILE_SAMPLE" // #nosec

	// The namespace where Tekton Chains and its secrets are deployed.
	TEKTON_CHAINS_NS string = "openshift-pipelines" // #nosec

	// User for running the end-to-end Tekton Chains tests
	TEKTON_CHAINS_E2E_USER string = "chains-e2e"

	// Name of the Secret Tekton Chains uses to read signing key
	TEKTON_CHAINS_SIGNING_SECRETS_NAME = "signing-secrets"

	//Cluster Registration namespace
	CLUSTER_REG_NS string = "cluster-reg-config" // #nosec

	// E2E test namespace where the app and component CRs will be created
	E2E_APPLICATIONS_NAMESPACE_ENV string = "E2E_APPLICATIONS_NAMESPACE"

	// Skip checking "ApplicationServiceGHTokenSecrName" secret
	SKIP_HAS_SECRET_CHECK_ENV string = "SKIP_HAS_SECRET_CHECK"

	// Sandbox kubeconfig user path
	USER_KUBE_CONFIG_PATH_ENV string = "USER_KUBE_CONFIG_PATH"
	// Release e2e auth for build and release quay keys

	QUAY_OAUTH_TOKEN_RELEASE_SOURCE string = "QUAY_OAUTH_TOKEN_RELEASE_SOURCE"

	QUAY_OAUTH_TOKEN_RELEASE_DESTINATION string = "QUAY_OAUTH_TOKEN_RELEASE_DESTINATION"

	// Key auth for accessing Pyxis stage external registry
	PYXIS_STAGE_KEY_ENV string = "PYXIS_STAGE_KEY"

	// Cert auth for accessing Pyxis stage external registry
	PYXIS_STAGE_CERT_ENV string = "PYXIS_STAGE_CERT"

	// Offline/refresh token used for getting Keycloak token in order to authenticate against stage/prod cluster
	// More details: https://access.redhat.com/articles/3626371
	OFFLINE_TOKEN_ENV = "OFFLINE_TOKEN"

	// Keycloak URL used for authentication against stage/prod cluster
	KEYLOAK_URL_ENV = "KEYLOAK_URL"

	// Toolchain API URL used for authentication against stage/prod cluster
	TOOLCHAIN_API_URL_ENV = "TOOLCHAIN_API_URL"

	// Dev workspace for release pipelines tests
	RELEASE_DEV_WORKSPACE_ENV = "RELEASE_DEV_WORKSPACE"

	// Managed workspace for release pipelines tests
	RELEASE_MANAGED_WORKSPACE_ENV = "RELEASE_MANAGED_WORKSPACE"

	// Bundle ref for overriding the default Java build bundle specified in BuildPipelineSelectorYamlURL
	CUSTOM_JAVA_PIPELINE_BUILD_BUNDLE_ENV string = "CUSTOM_JAVA_PIPELINE_BUILD_BUNDLE"

	// Bundle ref for a buildah-remote build
	CUSTOM_BUILDAH_REMOTE_PIPELINE_BUILD_BUNDLE_ENV string = "CUSTOM_BUILDAH_REMOTE_PIPELINE_BUILD_BUNDLE"

	// QE slack bot token used for delivering messages about critical failures during CI runs
	SLACK_BOT_TOKEN_ENV = "SLACK_BOT_TOKEN"

	// This variable is set by an automation in case Spray Proxy configuration fails in CI
	SKIP_PAC_TESTS_ENV = "SKIP_PAC_TESTS"

	// Test namespace's required labels
	ArgoCDLabelKey   string = "argocd.argoproj.io/managed-by"
	ArgoCDLabelValue string = "gitops-service-argocd"

	BuildPipelinesConfigMapDefaultNamespace = "build-templates"

	HostOperatorNamespace   string = "toolchain-host-operator"
	MemberOperatorNamespace string = "toolchain-member-operator"

	HostOperatorWorkload   string = "host-operator-controller-manager"
	MemberOperatorWorkload string = "member-operator-controller-manager"

	OLMOperatorNamespace string = "openshift-operator-lifecycle-manager"
	OLMOperatorWorkload  string = "olm-operator"

	OSAPIServerNamespace string = "openshift-apiserver"
	OSAPIServerWorkload  string = "apiserver"

	DefaultQuayOrg = "redhat-appstudio-qe"

	RegistryAuthSecretName = "redhat-appstudio-registry-pull-secret"

	QuayRepositorySecretName      = "quay-repository"
	QuayRepositorySecretNamespace = "e2e-secrets"

	JVMBuildImageSecretName = "jvm-build-image-secrets"
	JBSConfigName           = "jvm-build-config"

	BuildPipelineSelectorYamlURL = "https://raw.githubusercontent.com/redhat-appstudio/infra-deployments/main/components/build-service/base/build-pipeline-selectors/build-pipeline-selector.yaml"

	DefaultImagePushRepo         = "quay.io/" + DefaultQuayOrg + "/test-images"
	DefaultReleasedImagePushRepo = "quay.io/" + DefaultQuayOrg + "/test-release-images"

	BuildTaskRunName = "build-container"

	ReleasePipelineImageRef = "quay.io/hacbs-release/pipeline-release:0.20"

	FromIndex   = "registry-proxy.engineering.redhat.com/rh-osbs/iib-preview-rhtap:v4.13"
	TargetIndex = "quay.io/redhat/redhat----preview-operator-index:v4.13"
	BinaryImage = "registry.redhat.io/openshift4/ose-operator-registry:v4.13"

	StrategyConfigsRepo          = "strategy-configs"
	StrategyConfigsDefaultBranch = "main"
	StrategyConfigsRevision      = "caeaaae63a816ab42dad6c7be1e4b352ea8aabf4"

	// TODO
	// delete this constant and all its occurrences in the code base
	// once https://issues.redhat.com/browse/RHTAP-810 is completed
	OldTektonTaskTestOutputName = "HACBS_TEST_OUTPUT"

	TektonTaskTestOutputName = "TEST_OUTPUT"

	DefaultPipelineServiceAccount            = "appstudio-pipeline"
	DefaultPipelineServiceAccountRoleBinding = "appstudio-pipelines-runner-rolebinding"
	DefaultPipelineServiceAccountClusterRole = "appstudio-pipelines-runner"

	PaCPullRequestBranchPrefix = "appstudio-"

	// Expiration for image tags
	IMAGE_TAG_EXPIRATION_ENV  string = "IMAGE_TAG_EXPIRATION"
	DefaultImageTagExpiration string = "6h"

	PipelineRunPollingInterval = 10 * time.Second

	JsonStageUsersPath = "users.json"

	SamplePrivateRepoName = "test-private-repo"

	// Github App name is RHTAP-Qe-App. Note: this App ID is used in our CI and can't be used for local dev/testing.
	DefaultPaCGitHubAppID = "310332"

	// Error string constants for Namespace-backed environment test suite
	SEBAbsenceErrorString          = "no SnapshotEnvironmentBinding found in environment"
	EphemeralEnvAbsenceErrorString = "no matching Ephemeral Environment found"

	// #app-studio-ci-reports channel id
	SlackCIReportsChannelID = "C02M210JZ7B"

	DevReleaseTeam     = "dev-release-team"
	ManagedReleaseTeam = "managed-release-team"
)

var (
	ComponentDefaultLabel                       = map[string]string{"e2e-test": "true"}
	ComponentPaCRequestAnnotation               = map[string]string{"build.appstudio.openshift.io/request": "configure-pac"}
	ComponentTriggerSimpleBuildAnnotation       = map[string]string{"build.appstudio.openshift.io/request": "trigger-simple-build"}
	ImageControllerAnnotationRequestPublicRepo  = map[string]string{"image.redhat.com/generate": `{"visibility": "public"}`}
	ImageControllerAnnotationRequestPrivateRepo = map[string]string{"image.redhat.com/generate": `{"visibility": "private"}`}
	IntegrationTestScenarioDefaultLabels        = map[string]string{"test.appstudio.openshift.io/optional": "false"}
)
