#!/usr/bin/env python3

import yaml

SPEC_PATH = '../api/openapi.yaml'

# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for component_name, schema in data['components']['schemas'].items():
        if 'properties' in schema:
            # Remove "null" item from nullable enums
            for name, prop in schema['properties'].items():
                if 'enum' in prop and None in prop['enum']:
                    prop['enum'].remove(None)
                if 'properties' in prop and 'value' in prop['properties'] and 'enum' in prop['properties']['value'] and None in prop['properties']['value']['enum']:
                    prop['properties']['value']['enum'].remove(None)

            # Fix nullable types
            nullable_types = [
                'parent_device',
                'primary_ip',
            ]

            for ntype in nullable_types:
                if ntype in schema['properties']:
                    schema['properties'][ntype]['nullable'] = True

            # Fix non-nullable types
            # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
            non_nullable_types = [
                'front_image',
                'rear_image',
            ]

            for ntype in non_nullable_types:
                if ntype in schema['properties']:
                    if schema['properties'][ntype]['format'] == 'binary':
                        if 'nullable' in schema['properties'][ntype]:
                            schema['properties'][ntype].pop('nullable')

            change_type = {
                "BriefCustomFieldChoiceSet": {
                    "choices_count": "integer"
                },
                "CustomFieldChoiceSet": {
                    "choices_count": "integer"
                }
            }

            if component_name in change_type.keys():
                for propertie in change_type[component_name].keys():
                    schema['properties'][propertie]['type'] = change_type[component_name][propertie]

            if 'required' in schema:
                non_required = {
                    "BriefManufacturer": ["devicetype_count"],
                    "BriefRIR": ["aggregate_count"],
                    "ASN": ["created", "last_updated", "site_count", "provider_count"],
                    "BriefRackRole": ["rack_count"],
                    "BriefClusterType": ["cluster_count"],
                    "BriefClusterGroup": ["cluster_count"],
                    "BriefCluster": ["virtualmachine_count"],
                    "BriefDeviceRole": ["device_count", "virtualmachine_count"],
                    "BriefPlatform": ["device_count", "virtualmachine_count"],
                    "BriefVRF": ["prefix_count"],
                    "BriefVLANGroup": ["vlan_count"],
                    "VLAN": ["l2vpn_termination", "created", "last_updated", "prefix_count"],
                    "BriefRole": ["prefix_count", "vlan_count"],
                }

                if component_name in non_required.keys():
                    for r in non_required[component_name]:
                        if r in schema['required']:
                            schema['required'].remove(r)


# Save the spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)
