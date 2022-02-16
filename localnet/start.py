import os
import sys
import json
import subprocess

if pathlib.PurePath(os.getcwd()).name != 'localnet':
    print('script must be run from localnet folder')
    exit(1)

# Debug mode
debugMode = False

# Staking

# Validator set size, setting the value to 1 allows testing full validator set change
maxValidator = 10

# Self-delegation must be lower than 200000000stake
selfDelegationVal1 = '70000000stake'
selfDelegationVal2 = '60000000stake'
selfDelegationVal3 = '50000000stake'

# Unbonding time in seconds
unbondingTime = 1

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
genesis['app_state']['staking']['params']['max_validators'] = maxValidator
genesis['app_state']['staking']['params']['unbonding_time'] = str(unbondingTime)+"s"

# Create the gentxs
os.system('spnd gentx alice {} --chain-id spn-1 --moniker="bob" --home ./node1 --output-document ./gentx1.json'.format(selfDelegationVal1))
gentx1File = open('./gentx1.json')
gentx1 = json.load(gentx1File)

os.system('spnd gentx bob {} --chain-id spn-1 --moniker="carol" --home ./node2 --output-document ./gentx2.json'.format(selfDelegationVal2))
gentx2File = open('./gentx2.json')
gentx2 = json.load(gentx2File)

os.system('spnd gentx carol {} --chain-id spn-1 --moniker="dave" --home ./node3 --output-document ./gentx3.json'.format(selfDelegationVal3))
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
subprocess.Popen(["spnd", "start", "--home", "./node2"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
subprocess.Popen(["spnd", "start", "--home", "./node3"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
subprocess.run(["spnd start --home ./node1"], shell=True, check=True)

