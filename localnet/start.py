import os
import sys
import json
import subprocess
import pathlib

if pathlib.PurePath(os.getcwd()).name != 'localnet':
    print('script must be run from localnet folder')
    exit(1)

# Debug mode
debugMode = False

# Staking

# Validator set size, setting the value to 1 allows testing full validator set change
maxValidators = 100

# Self-delegation must be lower than 200000000
selfDelegationVal1 = 10000000
selfDelegationVal2 = 10000000
selfDelegationVal3 = 10000000

# Unbonding time in seconds
# Default: 21 days = 1814400 seconds
unbondingTime = 1814400

# Denom
denom = 'uspn'

# Reset all nodes
os.system('spnd unsafe-reset-all --home ./node1')
os.system('spnd unsafe-reset-all --home ./node2')
os.system('spnd unsafe-reset-all --home ./node3')

# Open the genesis template
genesisFile = open('./genesis_template.json')
genesis = json.load(genesisFile)

# Set timestamp
genesis['genesis_time'] = "2022-02-10T10:29:59.410196Z"

# Set monitoring param
genesis['app_state']['monitoringc']['params']['debugMode'] = debugMode

# Set staking params
genesis['app_state']['staking']['params']['max_validators'] = maxValidators
# genesis['app_state']['staking']['params']['unbonding_time'] = str(unbondingTime)+"s"

# Create the gentxs
os.system('spnd gentx joe {} --chain-id spn-1 --moniker="joe" --home ./node1 --output-document ./gentx1.json'.format(str(selfDelegationVal1)+denom))
gentx1File = open('./gentx1.json')
gentx1 = json.load(gentx1File)

os.system('spnd gentx steve {} --chain-id spn-1 --moniker="steve" --home ./node2 --output-document ./gentx2.json'.format(str(selfDelegationVal2)+denom))
gentx2File = open('./gentx2.json')
gentx2 = json.load(gentx2File)

os.system('spnd gentx olivia {} --chain-id spn-1 --moniker="olivia" --home ./node3 --output-document ./gentx3.json'.format(str(selfDelegationVal3)+denom))
gentx3File = open('./gentx3.json')
gentx3 = json.load(gentx3File)

# Collect gentxs
genesis['app_state']['genutil']['gen_txs'].append(gentx1)
genesis['app_state']['genutil']['gen_txs'].append(gentx2)
genesis['app_state']['genutil']['gen_txs'].append(gentx3)

os.remove('./gentx1.json')
os.remove('./gentx2.json')
os.remove('./gentx3.json')

# Save genesis
with open('./node1/config/genesis.json', 'w', encoding='utf-8') as f:
    json.dump(genesis, f, ensure_ascii=False, indent=4)
with open('./node2/config/genesis.json', 'w', encoding='utf-8') as f:
    json.dump(genesis, f, ensure_ascii=False, indent=4)
with open('./node3/config/genesis.json', 'w', encoding='utf-8') as f:
    json.dump(genesis, f, ensure_ascii=False, indent=4)

print('Starting the network')
# subprocess.Popen(["spnd", "start", "--home", "./node2"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
# subprocess.Popen(["spnd", "start", "--home", "./node3"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
# subprocess.run(["spnd start --home ./node1"], shell=True, check=True)
#
