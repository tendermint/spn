import os
import json
import subprocess
import pathlib
import yaml
from clear import clear_home
from utils import cmd_devnull

def save_genesis(genesis):
    with open('./spn/node1/config/genesis.json', 'w', encoding='utf-8') as f:
        json.dump(genesis, f, ensure_ascii=False, indent=4)
    with open('./spn/node2/config/genesis.json', 'w', encoding='utf-8') as f:
        json.dump(genesis, f, ensure_ascii=False, indent=4)
    with open('./spn/node3/config/genesis.json', 'w', encoding='utf-8') as f:
        json.dump(genesis, f, ensure_ascii=False, indent=4)

def start_spn():
    if pathlib.PurePath(os.getcwd()).name != 'localnet':
        print('script must be run from localnet folder')
        exit(1)
    clear_home('spn')

    # Load config
    confFile = open('./conf.yml')
    conf = yaml.safe_load(confFile)

    # Open the genesis template
    genesisFile = open('./spn/genesis_template.json')
    genesis = json.load(genesisFile)

    # Each node's home must contain a valid genesis in order to generate a gentx
    # The initial genesis template is therefore first saved in each home
    save_genesis(genesis)

    # Set timestamp
    genesis['genesis_time'] = "2022-02-10T10:29:59.410196Z"

    # Set chain ID
    genesis['chain_id'] = conf['chain_id']

    # Set staking params
    genesis['app_state']['staking']['params']['max_validators'] = conf['max_validators']
    genesis['app_state']['staking']['params']['unbonding_time'] = str(conf['unbonding_time'])+"s"

    # Create the gentxs
    for i in range(3):
        gentxCmd = 'spnd gentx {valName} {selfDelegation} --chain-id {chainID} --moniker="{valName}" --home ./spn/node{i} --output-document ./gentx.json'.format(
            valName=conf['validator_names'][i],
            selfDelegation=str(conf['validator_self_delegations'][i])+conf['staking_denom'],
            chainID=conf['chain_id'],
            i=str(i+1),
        )
        cmd_devnull(gentxCmd)
        gentxFile = open('./gentx.json')
        gentx = json.load(gentxFile)
        genesis['app_state']['genutil']['gen_txs'].append(gentx)
        os.remove('./gentx.json')

    # Save genesis
    save_genesis(genesis)

    print('Starting the network')
    subprocess.Popen(["spnd", "start", "--home", "./spn/node2"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    subprocess.Popen(["spnd", "start", "--home", "./spn/node3"], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    subprocess.run(["spnd start --home ./spn/node1"], shell=True, check=True)

if __name__ == "__main__":
    start_spn()