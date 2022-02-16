import os
import pathlib

if pathlib.PurePath(os.getcwd()).name != 'localnet':
    print('script must be run from localnet folder')
    exit(1)

os.system('spnd unsafe-reset-all --home ./node1')
os.system('spnd unsafe-reset-all --home ./node2')
os.system('spnd unsafe-reset-all --home ./node3')
