// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Artifactory Local Nuget Repository Resource
 *
 * Creates a local Nuget repository and allows for the creation of a
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const terraform_local_test_nuget_repo_basic = new artifactory.LocalNugetRepository("terraform-local-test-nuget-repo-basic", {
 *     forceNugetAuthentication: true,
 *     key: "terraform-local-test-nuget-repo-basic",
 *     maxUniqueSnapshots: 5,
 * });
 * ```
 */
export class LocalNugetRepository extends pulumi.CustomResource {
    /**
     * Get an existing LocalNugetRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocalNugetRepositoryState, opts?: pulumi.CustomResourceOptions): LocalNugetRepository {
        return new LocalNugetRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/localNugetRepository:LocalNugetRepository';

    /**
     * Returns true if the given object is an instance of LocalNugetRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocalNugetRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocalNugetRepository.__pulumiType;
    }

    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    public readonly archiveBrowsingEnabled!: pulumi.Output<boolean | undefined>;
    public readonly blackedOut!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly downloadDirect!: pulumi.Output<boolean | undefined>;
    public readonly excludesPattern!: pulumi.Output<string>;
    /**
     * - Force basic authentication credentials in order to use this repository.
     */
    public readonly forceNugetAuthentication!: pulumi.Output<boolean | undefined>;
    public readonly includesPattern!: pulumi.Output<string>;
    /**
     * - the identity key of the repo
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * - The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     */
    public readonly maxUniqueSnapshots!: pulumi.Output<number | undefined>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    public readonly repoLayoutRef!: pulumi.Output<string>;
    public readonly xrayIndex!: pulumi.Output<boolean>;

    /**
     * Create a LocalNugetRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalNugetRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalNugetRepositoryArgs | LocalNugetRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocalNugetRepositoryState | undefined;
            inputs["archiveBrowsingEnabled"] = state ? state.archiveBrowsingEnabled : undefined;
            inputs["blackedOut"] = state ? state.blackedOut : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["downloadDirect"] = state ? state.downloadDirect : undefined;
            inputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            inputs["forceNugetAuthentication"] = state ? state.forceNugetAuthentication : undefined;
            inputs["includesPattern"] = state ? state.includesPattern : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["maxUniqueSnapshots"] = state ? state.maxUniqueSnapshots : undefined;
            inputs["notes"] = state ? state.notes : undefined;
            inputs["packageType"] = state ? state.packageType : undefined;
            inputs["propertySets"] = state ? state.propertySets : undefined;
            inputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            inputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as LocalNugetRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            inputs["archiveBrowsingEnabled"] = args ? args.archiveBrowsingEnabled : undefined;
            inputs["blackedOut"] = args ? args.blackedOut : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["downloadDirect"] = args ? args.downloadDirect : undefined;
            inputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            inputs["forceNugetAuthentication"] = args ? args.forceNugetAuthentication : undefined;
            inputs["includesPattern"] = args ? args.includesPattern : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["maxUniqueSnapshots"] = args ? args.maxUniqueSnapshots : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["propertySets"] = args ? args.propertySets : undefined;
            inputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            inputs["xrayIndex"] = args ? args.xrayIndex : undefined;
            inputs["packageType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LocalNugetRepository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocalNugetRepository resources.
 */
export interface LocalNugetRepositoryState {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * - Force basic authentication credentials in order to use this repository.
     */
    forceNugetAuthentication?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * - the identity key of the repo
     */
    key?: pulumi.Input<string>;
    /**
     * - The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     */
    maxUniqueSnapshots?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    packageType?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a LocalNugetRepository resource.
 */
export interface LocalNugetRepositoryArgs {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * - Force basic authentication credentials in order to use this repository.
     */
    forceNugetAuthentication?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * - the identity key of the repo
     */
    key: pulumi.Input<string>;
    /**
     * - The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     */
    maxUniqueSnapshots?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}