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
    /// ## # Artifactory Backup Resource
    /// 
    /// This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.
    /// 
    /// When a artifactory.Backup resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.  The backup process creates a time-stamped directory in the target backup directory.
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
    ///         // Configure Artifactory Backup system config
    ///         var backupConfigName = new Artifactory.Backup("backupConfigName", new Artifactory.BackupArgs
    ///         {
    ///             CreateArchive = false,
    ///             CronExp = "0 0 12 * * ?",
    ///             Enabled = true,
    ///             ExcludeNewRepositories = true,
    ///             ExcludedRepositories = {},
    ///             Key = "backup_config_name",
    ///             RetentionPeriodHours = 1000,
    ///             SendMailOnError = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// Note: `Key` argument has to match to the resource name.\
    /// Reference Link: [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups)
    /// 
    /// ## Import
    /// 
    /// Backup config can be imported using the key, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/backup:Backup backup_name backup_name
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/backup:Backup")]
    public partial class Backup : Pulumi.CustomResource
    {
        /// <summary>
        /// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        /// </summary>
        [Output("createArchive")]
        public Output<bool?> CreateArchive { get; private set; } = null!;

        /// <summary>
        /// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        /// </summary>
        [Output("cronExp")]
        public Output<string> CronExp { get; private set; } = null!;

        /// <summary>
        /// Flag to enable or disable the backup config. Default value is `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// When set, new repositories will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Output("excludeNewRepositories")]
        public Output<bool?> ExcludeNewRepositories { get; private set; } = null!;

        /// <summary>
        /// A list of excluded repositories from the backup. Default is empty list.
        /// </summary>
        [Output("excludedRepositories")]
        public Output<ImmutableArray<string>> ExcludedRepositories { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the artifactory backup config.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        /// </summary>
        [Output("retentionPeriodHours")]
        public Output<int?> RetentionPeriodHours { get; private set; } = null!;

        /// <summary>
        /// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        /// </summary>
        [Output("sendMailOnError")]
        public Output<bool?> SendMailOnError { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/backup:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/backup:Backup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, state, options);
        }
    }

    public sealed class BackupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        /// </summary>
        [Input("createArchive")]
        public Input<bool>? CreateArchive { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        /// </summary>
        [Input("cronExp", required: true)]
        public Input<string> CronExp { get; set; } = null!;

        /// <summary>
        /// Flag to enable or disable the backup config. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When set, new repositories will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Input("excludeNewRepositories")]
        public Input<bool>? ExcludeNewRepositories { get; set; }

        [Input("excludedRepositories")]
        private InputList<string>? _excludedRepositories;

        /// <summary>
        /// A list of excluded repositories from the backup. Default is empty list.
        /// </summary>
        public InputList<string> ExcludedRepositories
        {
            get => _excludedRepositories ?? (_excludedRepositories = new InputList<string>());
            set => _excludedRepositories = value;
        }

        /// <summary>
        /// The unique ID of the artifactory backup config.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        /// </summary>
        [Input("retentionPeriodHours")]
        public Input<int>? RetentionPeriodHours { get; set; }

        /// <summary>
        /// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        /// </summary>
        [Input("sendMailOnError")]
        public Input<bool>? SendMailOnError { get; set; }

        public BackupArgs()
        {
        }
    }

    public sealed class BackupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        /// </summary>
        [Input("createArchive")]
        public Input<bool>? CreateArchive { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        /// <summary>
        /// Flag to enable or disable the backup config. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When set, new repositories will not be automatically added to the backup. Default value is `false`.
        /// </summary>
        [Input("excludeNewRepositories")]
        public Input<bool>? ExcludeNewRepositories { get; set; }

        [Input("excludedRepositories")]
        private InputList<string>? _excludedRepositories;

        /// <summary>
        /// A list of excluded repositories from the backup. Default is empty list.
        /// </summary>
        public InputList<string> ExcludedRepositories
        {
            get => _excludedRepositories ?? (_excludedRepositories = new InputList<string>());
            set => _excludedRepositories = value;
        }

        /// <summary>
        /// The unique ID of the artifactory backup config.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        /// </summary>
        [Input("retentionPeriodHours")]
        public Input<int>? RetentionPeriodHours { get; set; }

        /// <summary>
        /// If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        /// </summary>
        [Input("sendMailOnError")]
        public Input<bool>? SendMailOnError { get; set; }

        public BackupState()
        {
        }
    }
}
