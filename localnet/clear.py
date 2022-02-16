import os
import pathlib

if pathlib.PurePath(os.getcwd()).name != 'localnet':
    print('script must be run from localnet folder')
    exit(1)

os.system("rm ./node1/config/write-file-atomic-*")
os.system("rm ./node2/config/write-file-atomic-*")
os.system("rm ./node3/config/write-file-atomic-*")

os.system('spnd unsafe-reset-all --home ./node1')
os.system('spnd unsafe-reset-all --home ./node2')
os.system('spnd unsafe-reset-all --home ./node3')
