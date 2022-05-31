import sys
import subprocess
import yaml

if len(sys.argv) != 4:
    print('usage: delegate.py [val1_stake] [val2_stake] [val3_stake]')
    exit(0)

# Load config
confFile = open('./conf.yml')
conf = yaml.safe_load(confFile)

def delegate_cmd(valNumber, amount):
    cmd = ["spnd", "tx", "staking", "delegate"]
    cmd.append(conf['validator_addresses'][valNumber])

    stake = amount + conf['staking_denom']
    cmd.append(stake)

    cmd.append('--from')
    cmd.append(conf['validator_names'][valNumber])

    cmd.append('--chain-id')
    cmd.append(conf['chain_id'])

    cmd.append('-y')

    return cmd

def delegate(amounts):
    for s in amounts:
        if not s.isnumeric():
            print(s + ' must be a number')
            exit(1)

    i = 0
    for s in amounts:
        if int(s) > 0:
            print(i)
            cmd = delegate_cmd(i, s)
            print('running: ' + " ".join(cmd))
            subprocess.run(cmd, check=True)
        i += 1

if __name__ == "__main__":
    delegate(sys.argv[1:])

    print()
    print('delegation performed, to show validator set:')
    print('spnd q tendermint-validator-set')
    print()
    print('to show consensus state')
    print('spnd q ibc client self-consensus-state')