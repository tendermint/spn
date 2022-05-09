import os
import pathlib

def clearSPN():
    if pathlib.PurePath(os.getcwd()).name != 'localnet':
        print('script must be run from localnet folder')
        exit(1)

    for i in 3:
        os.system('rm spn/node{}/config/write-file-atomic-*'.format(i))
        os.system('rm spn/node{}/config/genesis.json'.format(i))
        os.system("rm spn/node{}/config/addrbook.json")
        os.system('spnd tendermint unsafe-reset-all --home spn/node{}'.format(i))

if __name__ == "__main__":
    clearSPN()