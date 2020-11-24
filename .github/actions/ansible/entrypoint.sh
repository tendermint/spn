#!/bin/sh

export HOME=/root
mkdir -p ~/.ssh && echo $INPUT_SSH_KEY_B64 | base64 -d > ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa

for host in $INPUT_KNOWN_HOSTS
do
    ssh-keyscan -H $host >> ~/.ssh/known_hosts;
done

if [ "$INPUT_INVENTORY_SCRIPT" == "true" ]
then
    chmod +x $INPUT_INVENTORY;
fi

sleep $INPUT_WAIT
ANSIBLE_HOST_KEY_CHECKING=False \
    ansible-playbook \
    -i $INPUT_INVENTORY \
    $INPUT_PLAYBOOK \
    -e "$INPUT_EXTRA_VARS" \
    --key-file ~/.ssh/id_rsa \
    --ssh-common-args="$INPUT_SSH_COMMON_ARGS"
