<!--
order: 0
title: Genesis Overview
parent:
  title: "genesis"
-->

# `genesis`

## Abstract

The module `genesis` allows a Cosmos SDK application to store and handle data related to the genesis of a Cosmos SDK blockchain.
This module is used by the `spn` blockchain and the `starport network` command line application to let blockchain creators to start and coordinate the launch of their blockchain.

## Contents

1. **[State](01_state.md)**
    - [Chain](01_state.md#chain)
    - [Proposal](01_state.md#proposal)
    - [ProposalPools](01_state.md#proposalpools)
    - [GenesisInternalState](01_state.md#genesisinternalstate)
2. **[Messages](02_messages.md)**
    - [MsgChainCreate](02_messages.md#msgchaincreate)
    - [MsgProposalAddAccount](02_messages.md#msgproposaladdaccount)
    - [MsgProposalAddValidator](02_messages.md#msgproposaladdvalidator)
    - [MsgApprove](02_messages.md#msgapprove)
    - [MsgReject](02_messages.md#msgreject)