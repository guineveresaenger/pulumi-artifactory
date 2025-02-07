# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BackupArgs', 'Backup']

@pulumi.input_type
class BackupArgs:
    def __init__(__self__, *,
                 cron_exp: pulumi.Input[str],
                 key: pulumi.Input[str],
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Backup resource.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        :param pulumi.Input[str] key: The unique ID of the artifactory backup config.
        :param pulumi.Input[bool] create_archive: If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: A list of excluded repositories from the backup. Default is empty list.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        pulumi.set(__self__, "cron_exp", cron_exp)
        pulumi.set(__self__, "key", key)
        if create_archive is not None:
            pulumi.set(__self__, "create_archive", create_archive)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if exclude_new_repositories is not None:
            pulumi.set(__self__, "exclude_new_repositories", exclude_new_repositories)
        if excluded_repositories is not None:
            pulumi.set(__self__, "excluded_repositories", excluded_repositories)
        if retention_period_hours is not None:
            pulumi.set(__self__, "retention_period_hours", retention_period_hours)
        if send_mail_on_error is not None:
            pulumi.set(__self__, "send_mail_on_error", send_mail_on_error)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> pulumi.Input[str]:
        """
        A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: pulumi.Input[str]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The unique ID of the artifactory backup config.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="createArchive")
    def create_archive(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        """
        return pulumi.get(self, "create_archive")

    @create_archive.setter
    def create_archive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_archive", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to enable or disable the backup config. Default value is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="excludeNewRepositories")
    def exclude_new_repositories(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, new repositories will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "exclude_new_repositories")

    @exclude_new_repositories.setter
    def exclude_new_repositories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_new_repositories", value)

    @property
    @pulumi.getter(name="excludedRepositories")
    def excluded_repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of excluded repositories from the backup. Default is empty list.
        """
        return pulumi.get(self, "excluded_repositories")

    @excluded_repositories.setter
    def excluded_repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_repositories", value)

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> Optional[pulumi.Input[int]]:
        """
        The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        """
        return pulumi.get(self, "retention_period_hours")

    @retention_period_hours.setter
    def retention_period_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period_hours", value)

    @property
    @pulumi.getter(name="sendMailOnError")
    def send_mail_on_error(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        return pulumi.get(self, "send_mail_on_error")

    @send_mail_on_error.setter
    def send_mail_on_error(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_mail_on_error", value)


@pulumi.input_type
class _BackupState:
    def __init__(__self__, *,
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Backup resources.
        :param pulumi.Input[bool] create_archive: If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: A list of excluded repositories from the backup. Default is empty list.
        :param pulumi.Input[str] key: The unique ID of the artifactory backup config.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        if create_archive is not None:
            pulumi.set(__self__, "create_archive", create_archive)
        if cron_exp is not None:
            pulumi.set(__self__, "cron_exp", cron_exp)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if exclude_new_repositories is not None:
            pulumi.set(__self__, "exclude_new_repositories", exclude_new_repositories)
        if excluded_repositories is not None:
            pulumi.set(__self__, "excluded_repositories", excluded_repositories)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if retention_period_hours is not None:
            pulumi.set(__self__, "retention_period_hours", retention_period_hours)
        if send_mail_on_error is not None:
            pulumi.set(__self__, "send_mail_on_error", send_mail_on_error)

    @property
    @pulumi.getter(name="createArchive")
    def create_archive(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        """
        return pulumi.get(self, "create_archive")

    @create_archive.setter
    def create_archive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_archive", value)

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> Optional[pulumi.Input[str]]:
        """
        A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        """
        return pulumi.get(self, "cron_exp")

    @cron_exp.setter
    def cron_exp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_exp", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to enable or disable the backup config. Default value is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="excludeNewRepositories")
    def exclude_new_repositories(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, new repositories will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "exclude_new_repositories")

    @exclude_new_repositories.setter
    def exclude_new_repositories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_new_repositories", value)

    @property
    @pulumi.getter(name="excludedRepositories")
    def excluded_repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of excluded repositories from the backup. Default is empty list.
        """
        return pulumi.get(self, "excluded_repositories")

    @excluded_repositories.setter
    def excluded_repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_repositories", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID of the artifactory backup config.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> Optional[pulumi.Input[int]]:
        """
        The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        """
        return pulumi.get(self, "retention_period_hours")

    @retention_period_hours.setter
    def retention_period_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period_hours", value)

    @property
    @pulumi.getter(name="sendMailOnError")
    def send_mail_on_error(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        return pulumi.get(self, "send_mail_on_error")

    @send_mail_on_error.setter
    def send_mail_on_error(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_mail_on_error", value)


class Backup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Artifactory Backup Resource

        This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.

        When a Backup resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.  The backup process creates a time-stamped directory in the target backup directory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory Backup system config
        backup_config_name = artifactory.Backup("backupConfigName",
            create_archive=False,
            cron_exp="0 0 12 * * ?",
            enabled=True,
            exclude_new_repositories=True,
            excluded_repositories=[],
            key="backup_config_name",
            retention_period_hours=1000,
            send_mail_on_error=True)
        ```
        Note: `Key` argument has to match to the resource name.\
        Reference Link: [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups)

        ## Import

        Backup config can be imported using the key, e.g.

        ```sh
         $ pulumi import artifactory:index/backup:Backup backup_name backup_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_archive: If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: A list of excluded repositories from the backup. Default is empty list.
        :param pulumi.Input[str] key: The unique ID of the artifactory backup config.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Artifactory Backup Resource

        This resource can be used to manage the automatic and periodic backups of the entire Artifactory instance.

        When a Backup resource is configured and enabled to true, backup of the entire Artifactory system will be done automatically and periodically.  The backup process creates a time-stamped directory in the target backup directory.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory Backup system config
        backup_config_name = artifactory.Backup("backupConfigName",
            create_archive=False,
            cron_exp="0 0 12 * * ?",
            enabled=True,
            exclude_new_repositories=True,
            excluded_repositories=[],
            key="backup_config_name",
            retention_period_hours=1000,
            send_mail_on_error=True)
        ```
        Note: `Key` argument has to match to the resource name.\
        Reference Link: [JFrog Artifactory Backup](https://www.jfrog.com/confluence/display/JFROG/Backups)

        ## Import

        Backup config can be imported using the key, e.g.

        ```sh
         $ pulumi import artifactory:index/backup:Backup backup_name backup_name
        ```

        :param str resource_name: The name of the resource.
        :param BackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_archive: Optional[pulumi.Input[bool]] = None,
                 cron_exp: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
                 excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 retention_period_hours: Optional[pulumi.Input[int]] = None,
                 send_mail_on_error: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupArgs.__new__(BackupArgs)

            __props__.__dict__["create_archive"] = create_archive
            if cron_exp is None and not opts.urn:
                raise TypeError("Missing required property 'cron_exp'")
            __props__.__dict__["cron_exp"] = cron_exp
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["exclude_new_repositories"] = exclude_new_repositories
            __props__.__dict__["excluded_repositories"] = excluded_repositories
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["retention_period_hours"] = retention_period_hours
            __props__.__dict__["send_mail_on_error"] = send_mail_on_error
        super(Backup, __self__).__init__(
            'artifactory:index/backup:Backup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_archive: Optional[pulumi.Input[bool]] = None,
            cron_exp: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            exclude_new_repositories: Optional[pulumi.Input[bool]] = None,
            excluded_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            key: Optional[pulumi.Input[str]] = None,
            retention_period_hours: Optional[pulumi.Input[int]] = None,
            send_mail_on_error: Optional[pulumi.Input[bool]] = None) -> 'Backup':
        """
        Get an existing Backup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_archive: If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        :param pulumi.Input[str] cron_exp: A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        :param pulumi.Input[bool] enabled: Flag to enable or disable the backup config. Default value is `true`.
        :param pulumi.Input[bool] exclude_new_repositories: When set, new repositories will not be automatically added to the backup. Default value is `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_repositories: A list of excluded repositories from the backup. Default is empty list.
        :param pulumi.Input[str] key: The unique ID of the artifactory backup config.
        :param pulumi.Input[int] retention_period_hours: The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        :param pulumi.Input[bool] send_mail_on_error: If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupState.__new__(_BackupState)

        __props__.__dict__["create_archive"] = create_archive
        __props__.__dict__["cron_exp"] = cron_exp
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["exclude_new_repositories"] = exclude_new_repositories
        __props__.__dict__["excluded_repositories"] = excluded_repositories
        __props__.__dict__["key"] = key
        __props__.__dict__["retention_period_hours"] = retention_period_hours
        __props__.__dict__["send_mail_on_error"] = send_mail_on_error
        return Backup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createArchive")
    def create_archive(self) -> pulumi.Output[Optional[bool]]:
        """
        If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
        """
        return pulumi.get(self, "create_archive")

    @property
    @pulumi.getter(name="cronExp")
    def cron_exp(self) -> pulumi.Output[str]:
        """
        A valid CRON expression that you can use to control backup frequency. Eg: "0 0 12 * * ? "
        """
        return pulumi.get(self, "cron_exp")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to enable or disable the backup config. Default value is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="excludeNewRepositories")
    def exclude_new_repositories(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, new repositories will not be automatically added to the backup. Default value is `false`.
        """
        return pulumi.get(self, "exclude_new_repositories")

    @property
    @pulumi.getter(name="excludedRepositories")
    def excluded_repositories(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of excluded repositories from the backup. Default is empty list.
        """
        return pulumi.get(self, "excluded_repositories")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The unique ID of the artifactory backup config.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> pulumi.Output[Optional[int]]:
        """
        The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
        """
        return pulumi.get(self, "retention_period_hours")

    @property
    @pulumi.getter(name="sendMailOnError")
    def send_mail_on_error(self) -> pulumi.Output[Optional[bool]]:
        """
        If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
        """
        return pulumi.get(self, "send_mail_on_error")

