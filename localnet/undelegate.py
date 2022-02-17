import sys
import subprocess

if len(sys.argv) != 4:
    print('usage: undelegate.py [val1_stake] [val2_stake] [val3_stake]')

chainID = 'spn-1'
denom = 'uspn'

valName = [
    "joe",
    "steve",
    "olivia",
]

valAddr = [
    "spnvaloper15rz2rwnlgr7nf6eauz52usezffwrxc0muf4z5n",
    "spnvaloper1mhyps2hlkm0nz6k2puumn69928cnvgg4nznru5",
    "spnvaloper1hmx8eakt2948szjgmksvpv9ha0q9s6w09pdeer",
]

def delegate_cmd(valNumber, amount):
    cmd = ["spnd", "tx", "staking", "unbond"]
    cmd.append(valAddr[valNumber])

    stake = amount + denom
    cmd.append(stake)

    cmd.append('--from')
    cmd.append(valName[valNumber])

    cmd.append('--chain-id')
    cmd.append(chainID)

    cmd.append('-y')

    return cmd

# Perform unbonding
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
print('unbonding performed, to show validator set:')
print('spnd q tendermint-validator-set')