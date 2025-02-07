// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class PermissionTargetReleaseBundle
    {
        /// <summary>
        /// -
        /// </summary>
        public readonly Outputs.PermissionTargetReleaseBundleActions? Actions;
        /// <summary>
        /// Pattern of artifacts to exclude
        /// </summary>
        public readonly ImmutableArray<string> ExcludesPatterns;
        /// <summary>
        /// Pattern of artifacts to include
        /// </summary>
        public readonly ImmutableArray<string> IncludesPatterns;
        /// <summary>
        /// List of repositories this permission target is applicable for
        /// </summary>
        public readonly ImmutableArray<string> Repositories;

        [OutputConstructor]
        private PermissionTargetReleaseBundle(
            Outputs.PermissionTargetReleaseBundleActions? actions,

            ImmutableArray<string> excludesPatterns,

            ImmutableArray<string> includesPatterns,

            ImmutableArray<string> repositories)
        {
            Actions = actions;
            ExcludesPatterns = excludesPatterns;
            IncludesPatterns = includesPatterns;
            Repositories = repositories;
        }
    }
}
