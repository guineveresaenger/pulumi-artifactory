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
    public sealed class PushReplicationReplication
    {
        public readonly bool? Enabled;
        public readonly string Password;
        public readonly string? PathPrefix;
        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        public readonly string? Proxy;
        public readonly int? SocketTimeoutMillis;
        public readonly bool? SyncDeletes;
        public readonly bool? SyncProperties;
        public readonly bool? SyncStatistics;
        public readonly string Url;
        public readonly string Username;

        [OutputConstructor]
        private PushReplicationReplication(
            bool? enabled,

            string password,

            string? pathPrefix,

            string? proxy,

            int? socketTimeoutMillis,

            bool? syncDeletes,

            bool? syncProperties,

            bool? syncStatistics,

            string url,

            string username)
        {
            Enabled = enabled;
            Password = password;
            PathPrefix = pathPrefix;
            Proxy = proxy;
            SocketTimeoutMillis = socketTimeoutMillis;
            SyncDeletes = syncDeletes;
            SyncProperties = syncProperties;
            SyncStatistics = syncStatistics;
            Url = url;
            Username = username;
        }
    }
}
