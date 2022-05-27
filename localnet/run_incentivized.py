import argparse
import time
from start_testnet import start_testnet
from initialize_rewards import initialize_rewards
from utils import cmd

parser = argparse.ArgumentParser(description='Start a testnet and connect it for reward')
parser.add_argument('--spn_chain_id',
                    help='Chain ID on SPN',
                    default='spn-1')
parser.add_argument('--orbit_chain_id',
                    help='Chain ID on Orbit',
                    default='orbit-1')
parser.add_argument('--debug',
                    action='store_true',
                    help='Set debug mode for module')
parser.add_argument('--spn_unbonding_period',
                    type=int,
                    default=1814400,
                    help='Unbonding period on spn',
                    )
parser.add_argument('--spn_revision_height',
                    type=int,
                    default=2,
                    help='Revision height for SPN IBC client',
                    )
parser.add_argument('--last_block_height',
                    type=int,
                    default=30,
                    help='Last block height for monitoring packet forwarding',
                    )
parser.add_argument('--max_validator',
                    type=int,
                    default=10,
                    help='Staking max validator set',
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
parser.add_argument('--unbonding_time',
                    default=1000, # 21 days = 1814400 seconds
                    type=int,
                    help='Staking unbonding time (unbonding period)',
                    )

if __name__ == "__main__":
    # Parse params
    args = parser.parse_args()
    debugMode = args.debug
    spnChainID = args.spn_chain_id
    chainID = args.orbit_chain_id
    spnUnbondingPeriod = args.spn_unbonding_period
    revisionHeight = args.spn_revision_height
    lastBlockHeight = args.last_block_height
    maxValidator = args.max_validator
    selfDelegationVal1 = args.self_delegation_1
    selfDelegationVal2 = args.self_delegation_2
    selfDelegationVal3 = args.self_delegation_3
    unbondingTime = args.unbonding_time

    # Initialize rewards
    print('intialize rewards')
    initialize_rewards(
        lastBlockHeight,
        selfDelegationVal1,
        selfDelegationVal2,
        selfDelegationVal3,
    )
    print('rewards initialized')

    cmd('spnd q ibc client self-consensus-state --height 2 > spncs.yaml')

    # Start the testnet
    print('start network')
    start(
        debugMode,
        spnChainID,
        chainID,
        spnUnbondingPeriod,
        revisionHeight,
        lastBlockHeight,
        maxValidator,
        selfDelegationVal1,
        selfDelegationVal2,
        selfDelegationVal3,
        unbondingTime,
        True,
    )
    print('network started')

    time.sleep(10)

    # Create verified IBC client on SPN
    print('create verified client')
    cmd('spnd q tendermint-validator-set 2 --node "tcp://localhost:26659" > vs.yaml')
    cmd('spnd q ibc client self-consensus-state --height 2 --node "tcp://localhost:26659" > cs.yaml')
    cmd('spnd tx monitoring-consumer create-client 1 cs.yaml vs.yaml --unbonding-period {} --revision-height 2 --from alice -y'.format(unbondingTime))

    # Perform IBC connection
    cmd('hermes -c ./hermes/config.toml create connection spn-1 --client-a 07-tendermint-0 --client-b 07-tendermint-0')
    cmd('hermes -c ./hermes/config.toml create channel spn-1 --connection-a connection-0 --port-a monitoring --port-b monitoring -o ordered --channel-version monitoring-1')

    # hermes -c ./hermes/config.toml start
