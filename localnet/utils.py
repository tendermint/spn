import subprocess

def cmd(command):
    subprocess.run([command], shell=True, check=True)

def cmd_devnull(command):
    subprocess.run([command], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)