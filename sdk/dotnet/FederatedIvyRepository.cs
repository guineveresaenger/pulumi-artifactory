// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// ## # Artifactory Federated Ivy Repository Resource
    /// 
    /// Creates a federated Ivy repository
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var terraform_federated_test_ivy_repo = new Artifactory.FederatedIvyRepository("terraform-federated-test-ivy-repo", new Artifactory.FederatedIvyRepositoryArgs
    ///         {
    ///             Key = "terraform-federated-test-ivy-repo",
    ///             Members = 
    ///             {
    ///                 new Artifactory.Inputs.FederatedIvyRepositoryMemberArgs
    ///                 {
    ///                     Enabled = true,
    ///                     Url = "http://tempurl.org/artifactory/terraform-federated-test-ivy-repo",
    ///                 },
    ///                 new Artifactory.Inputs.FederatedIvyRepositoryMemberArgs
    ///                 {
    ///                     Enabled = true,
    ///                     Url = "http://tempurl2.org/artifactory/terraform-federated-test-ivy-repo-2",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/federatedIvyRepository:FederatedIvyRepository")]
    public partial class FederatedIvyRepository : Pulumi.CustomResource
    {
        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Output("archiveBrowsingEnabled")]
        public Output<bool?> ArchiveBrowsingEnabled { get; private set; } = null!;

        [Output("blackedOut")]
        public Output<bool?> BlackedOut { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("downloadDirect")]
        public Output<bool?> DownloadDirect { get; private set; } = null!;

        [Output("excludesPattern")]
        public Output<string> ExcludesPattern { get; private set; } = null!;

        [Output("includesPattern")]
        public Output<string> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// - the identity key of the repo
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.FederatedIvyRepositoryMember>> Members { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("packageType")]
        public Output<string> PackageType { get; private set; } = null!;

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Output("priorityResolution")]
        public Output<bool?> PriorityResolution { get; private set; } = null!;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        [Output("projectEnvironments")]
        public Output<ImmutableArray<string>> ProjectEnvironments { get; private set; } = null!;

        /// <summary>
        /// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
        /// with project key, separated by a dash.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        [Output("propertySets")]
        public Output<ImmutableArray<string>> PropertySets { get; private set; } = null!;

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Output("repoLayoutRef")]
        public Output<string?> RepoLayoutRef { get; private set; } = null!;

        [Output("xrayIndex")]
        public Output<bool> XrayIndex { get; private set; } = null!;


        /// <summary>
        /// Create a FederatedIvyRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FederatedIvyRepository(string name, FederatedIvyRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/federatedIvyRepository:FederatedIvyRepository", name, args ?? new FederatedIvyRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FederatedIvyRepository(string name, Input<string> id, FederatedIvyRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/federatedIvyRepository:FederatedIvyRepository", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FederatedIvyRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FederatedIvyRepository Get(string name, Input<string> id, FederatedIvyRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new FederatedIvyRepository(name, id, state, options);
        }
    }

    public sealed class FederatedIvyRepositoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// - the identity key of the repo
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<Inputs.FederatedIvyRepositoryMemberArgs>? _members;

        /// <summary>
        /// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.FederatedIvyRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.FederatedIvyRepositoryMemberArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
        /// with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public FederatedIvyRepositoryArgs()
        {
        }
    }

    public sealed class FederatedIvyRepositoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// - the identity key of the repo
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("members")]
        private InputList<Inputs.FederatedIvyRepositoryMemberGetArgs>? _members;

        /// <summary>
        /// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.FederatedIvyRepositoryMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.FederatedIvyRepositoryMemberGetArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
        /// with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public FederatedIvyRepositoryState()
        {
        }
    }
}
