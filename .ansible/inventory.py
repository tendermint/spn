#!/usr/bin/env python3

import os
import json

inventory = {
    "validators": {
        "hosts": os.environ["VALIDATORS_IPS"].split(",") if os.getenv("VALIDATORS_IPS", "") != "" else [],
        "vars": {}
    },
    "sentries": {
        "hosts": os.environ["SENTRIES_IPS"].split(",") if os.getenv("SENTRIES_IPS", "") != "" else [],
        "vars": {}
    }
}

print(json.dumps(inventory, indent=2))