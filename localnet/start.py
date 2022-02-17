import os
import json
import subprocess
import pathlib
import yaml

if pathlib.PurePath(os.getcwd()).name != 'localnet':
    print('script must be run from localnet folder')
    exit(1)
subprocess.run(["python3", "./clear.py"], check=True)

# Load config
confFile = open('./conf.yml')
conf = yaml.safe_load(confFile)

# Open the genesis template
genesisFile = open('./genesis_template.json')
genesis = json.load(genesisFile)

# Set timestamp
genesis['genesis_time'] = "2022-02-10T10:29:59.410196Z"

# Set chain ID
genesis['chain_id'] = conf['chain_id']

# Set monitoring param
genesis['app_state']['monitoringc']['params']['debugMode'] = conf['debug_mode']

# Set staking params
genesis['app_state']['staking']['params']['max_validators'] = conf['max_validators']
genesis['app_state']['staking']['params']['unbonding_time'] = str(conf['unbonding_time'])+"s"

# Create the gentxs
for i in range(3):
    gentxCmd = 'spnd gentx {valName} {selfDelegation} --chain-id {chainID} --moniker="{valName}" --home ./node{i} --output-document ./gentx.json'.format(
        valName=conf['validator_names'][i],
        selfDelegation=str(conf['validator_self_delegations'][i])+conf['staking_denom'],
        chainID=conf['chain_id'],
        i=str(i+1),
    )
    os.system(gentxCmd)
    gentxFile = open('./gentx.json')
    gentx = json.load(gentxFile)
    genesis['app_state']['genutil']['gen_txs'].append(gentx)
    os.remove('./gentx.json')

# Save genesis
with open('./node1/config/genesis.json', 'w', encoding='utf-8') as f:
    json.dump(genesis, f, ensure_ascii=False, indent=4)
with open('./node2/config/genesis.json', 'w', encoding='utf-8') as f:
    json.dump(genesis, f, ensure_ascii=False, indent=4)
with open('./node3/config/genesis.json', 'w', encoding='utf-8') as f:
    json.dump(genesis, f, ensure_ascii=False, indent=4)

print('Starting the network')
subprocess.Popen(["spnd", "start", "--home", "./node2"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
subprocess.Popen(["spnd", "start", "--home", "./node3"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
subprocess.run(["spnd start --home ./node1"], shell=True, check=True)

