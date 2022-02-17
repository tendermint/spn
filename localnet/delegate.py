import sys
import subprocess
import yaml

if len(sys.argv) != 4:
    print('usage: delegate.py [val1_stake] [val2_stake] [val3_stake]')

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

# Perform delegation
for s in sys.argv[1:]:
    if not s.isnumeric():
        print(s + ' must be a number')
        exit(1)

i = 0
for s in sys.argv[1:]:
    if int(s) > 0:
        print(i)
        cmd = delegate_cmd(i, s)
        print('running: ' + " ".join(cmd))
        subprocess.run(cmd, check=True)
    i += 1

print()
print()
print('delegation performed, to show validator set:')
print('spnd q tendermint-validator-set')