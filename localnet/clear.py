import os
import pathlib

def clear_home(home):
    if pathlib.PurePath(os.getcwd()).name != 'localnet':
        print('script must be run from localnet folder')
        exit(1)

    for i in [1,2,3]:
        os.system('rm {}/node{}/config/write-file-atomic-*'.format(home, i))
        os.system('rm {}/node{}/config/genesis.json'.format(home, i))
        os.system("rm {}/node{}/config/addrbook.json".format(home, i))
        os.system('spnd tendermint unsafe-reset-all --home {}/node{}'.format(home, i))

if __name__ == "__main__":
    clear_home('spn')
    clear_home('testnet')