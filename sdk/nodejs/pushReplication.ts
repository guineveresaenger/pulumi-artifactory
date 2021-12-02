// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Push replication configs can be imported using their repo key, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/pushReplication:PushReplication foo-rep provider_test_source
 * ```
 */
export class PushReplication extends pulumi.CustomResource {
    /**
     * Get an existing PushReplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PushReplicationState, opts?: pulumi.CustomResourceOptions): PushReplication {
        return new PushReplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/pushReplication:PushReplication';

    /**
     * Returns true if the given object is an instance of PushReplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PushReplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PushReplication.__pulumiType;
    }

    public readonly cronExp!: pulumi.Output<string>;
    public readonly enableEventReplication!: pulumi.Output<boolean>;
    public readonly replications!: pulumi.Output<outputs.PushReplicationReplication[] | undefined>;
    public readonly repoKey!: pulumi.Output<string>;

    /**
     * Create a PushReplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PushReplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PushReplicationArgs | PushReplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PushReplicationState | undefined;
            inputs["cronExp"] = state ? state.cronExp : undefined;
            inputs["enableEventReplication"] = state ? state.enableEventReplication : undefined;
            inputs["replications"] = state ? state.replications : undefined;
            inputs["repoKey"] = state ? state.repoKey : undefined;
        } else {
            const args = argsOrState as PushReplicationArgs | undefined;
            if ((!args || args.cronExp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cronExp'");
            }
            if ((!args || args.repoKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repoKey'");
            }
            inputs["cronExp"] = args ? args.cronExp : undefined;
            inputs["enableEventReplication"] = args ? args.enableEventReplication : undefined;
            inputs["replications"] = args ? args.replications : undefined;
            inputs["repoKey"] = args ? args.repoKey : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PushReplication.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PushReplication resources.
 */
export interface PushReplicationState {
    cronExp?: pulumi.Input<string>;
    enableEventReplication?: pulumi.Input<boolean>;
    replications?: pulumi.Input<pulumi.Input<inputs.PushReplicationReplication>[]>;
    repoKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PushReplication resource.
 */
export interface PushReplicationArgs {
    cronExp: pulumi.Input<string>;
    enableEventReplication?: pulumi.Input<boolean>;
    replications?: pulumi.Input<pulumi.Input<inputs.PushReplicationReplication>[]>;
    repoKey: pulumi.Input<string>;
}