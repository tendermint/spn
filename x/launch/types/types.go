package types

// REVERT_DELAY is the delay after the launch time when it is possible to revert the launch of the chain
// Chain launch can be reverted on-chain when the actual chain launch failed (incorrect gentx, etc...)
// This delay must be small be big enough to ensure nodes had the time to bootstrap\
// This currently corresponds to 1 hour
const RevertDelay int64 = 60 * 60
