import subprocess

val_address = 'spnvaloper15rz2rwnlgr7nf6eauz52usezffwrxc0muf4z5n'

def cmd(command):
    subprocess.run([command], shell=True, check=True)

def cmd_devnull(command):
    subprocess.run([command], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)

def date_f(d):
    return d.isoformat("T") + "Z"

def initialize_campaign():
    cmd_devnull('spnd tx staking delegate {} 100000uspn --from bob -y'.format(val_address))
    cmd_devnull('spnd tx staking delegate {} 100000uspn --from carol -y'.format(val_address))
    cmd_devnull('spnd tx staking delegate {} 100000uspn --from dave -y'.format(val_address))
    cmd_devnull('spnd tx profile create-coordinator --from alice -y')
    cmd_devnull('spnd tx campaign create-campaign orbit 1000000orbit --from alice -y')
    cmd_devnull('spnd tx campaign mint-vouchers 1 100000orbit --from alice -y')