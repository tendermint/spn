import argparse
from utils import cmd_devnull

parser = argparse.ArgumentParser(description='Initialize the rewards on SPN for the testnet')
parser.add_argument('--last_block_height',
                    type=int,
                    default=100,
                    help='Last block for the reward pool',
                    )
parser.add_argument('--self_delegation_1',
                    default='10000000uspn',
                    help='Self delegation for validator 1',
                    )
parser.add_argument('--self_delegation_2',
                    default='10000000uspn',
                    help='Self delegation for validator 2',
                    )
parser.add_argument('--self_delegation_3',
                    default='10000000uspn',
                    help='Self delegation for validator 3',
                    )

def initialize_rewards(lastBlockHeight, selfDelegationVal1, selfDelegationVal2, selfDelegationVal3):
    cmd_devnull('spnd tx profile create-coordinator --from alice -y')
    cmd_devnull('spnd tx launch create-chain orbit-1 orbit.com 0xaaa --from alice -y')
    cmd_devnull('spnd tx campaign create-campaign orbit 1000000orbit --from alice -y')
    cmd_devnull('spnd tx campaign mint-vouchers 1 50000orbit --from alice -y')
    cmd_devnull('spnd tx reward set-rewards 1 50000v/1/orbit {} --from alice -y'.format(lastBlockHeight))

    gentx1 = './testnet/node1/config/gentx/gentx.json'
    gentx2 = './testnet/node2/config/gentx/gentx.json'
    gentx3 = './testnet/node3/config/gentx/gentx.json'
    pub1 = '"Q5D7koejne/P2F1iIcSSVo6M4siL5anwHH7iopX66ps="'
    pub2 = '"JzzB4Kr09x3k1MdatVL7MBMrZUn0D3Lx9AK+nHWjbq0="'
    pub3 = '"4TwlBGJhu4ZDRBDK57GiFyAFafDAapa6nVQ0VvG5rjA="'
    val1 = 'spn1aqn8ynvr3jmq67879qulzrwhchq5dtrvtx0nhe'
    val2 = 'spn1pkdk6m2nh77nlaep84cylmkhjder3arey7rll5'
    val3 = 'spn1twckcceyw43da9j247pfs3yhqsv25j38grh68q'

    cmd_devnull('spnd tx launch request-add-validator 1 {} {} {} aaa foo.com --validator-address {} --from alice -y'.format(gentx1, pub1, selfDelegationVal1, val1)),
    cmd_devnull('spnd tx launch request-add-validator 1 {} {} {} aaa foo.com --validator-address {} --from alice -y'.format(gentx2, pub2, selfDelegationVal2, val2)),
    cmd_devnull('spnd tx launch request-add-validator 1 {} {} {} aaa foo.com --validator-address {} --from alice -y'.format(gentx3, pub3, selfDelegationVal3, val3)),

    # Uncomment for testing incomplete validator set
    # cmd_devnull('spnd tx launch request-add-validator 1 ./node3/config/gentx/gentx.json "FyTmyvZhwRjwqhY6eWykTfiE+0mwe+U0aSo3ti8DCW8=" 16000000stake aaa foo.com --validator-address spn1ezptsm3npn54qx9vvpah4nymre59ykr9exx2ul --from alice -y')

    cmd_devnull('spnd tx launch trigger-launch 1 5 --from alice -y')


if __name__ == "__main__":
    # Parse params
    args = parser.parse_args()
    lastBlockHeight = args.last_block_height
    selfDelegationVal1 = args.self_delegation_1
    selfDelegationVal2 = args.self_delegation_2
    selfDelegationVal3 = args.self_delegation_3

    rewards(lastBlockHeight, selfDelegationVal1, selfDelegationVal2, selfDelegationVal3)