# Coordinator CLI Guide

Coordinators organize and launch new chains on `spn`

---

## Publish a chain

```shell
ignite n publish https://github.com/lubtd/planet
```

#### Output

```shell
✔ Source code fetched
✔ Blockchain set up
✔ Chain's binary built
✔ Blockchain initialized
✔ Genesis initialized
✔ Network published
⋆ Launch ID: 6
```

`LaunchID` identifies the published blockchain on Ignite blockchain

---

## List all published chains

```shell
ignite n chain list
```

### Output

```shell
Launch Id 	Chain Id 	Source ...
6 		      planet-1 	https://github.com/lubtd/planet
5 		      spn-11 		https://github.com/tendermint/spn.git
4 		      spn-11 		https://github.com/tendermint/spn.git
3 		      mars-1 		https://github.com/lubtd/planet.git
2 		      spn-10 		https://github.com/tendermint/spn.git
1 		      spn-1 		https://github.com/tendermint/spn
```

---

## Approve validator requests

First, list requests:

```shell
ignite n request list 6
```

*NOTE: here "6" is specifying the `LaunchID`*

#### Output

```shell
Id 	Status 		Type 			Content
1 	APPROVED 	Add Genesis Account 	spn1daefnhnupn85e8vv0yc5epmnkcr5epkqncn2le, 100000000stake
2 	APPROVED 	Add Genesis Validator 	e3d3ca59d8214206839985712282967aaeddfb01@84.118.211.157:26656, spn1daefnhnupn85e8vv0yc5epmnkcr5epkqncn2le, 95000000stake
3 	PENDING 	Add Genesis Account 	spn1daefnhnupn85e8vv0yc5epmnkcr5epkqncn2le, 95000000stake
4 	PENDING 	Add Genesis Validator 	b10f3857133907a14dca5541a14df9e8e3389875@84.118.211.157:26656, spn1daefnhnupn85e8vv0yc5epmnkcr5epkqncn2le, 95000000stake
```

Approve the requests.  Both syntaxes can be used: `1,2,3,4` and `1-3,4`.

```shell
ignite n request approve 6 3,4
```

#### Output

```shell
✔ Source code fetched
✔ Blockchain set up
✔ Requests format verified
✔ Blockchain initialized
✔ Genesis initialized
✔ Genesis built
✔ The network can be started
✔ Request(s) #3, #4 verified
✔ Request(s) #3, #4 approved
```

---

## Initiate the launch of a chain

```shell
ignite n chain launch 6
```

#### Output

```shell
✔ Chain 6 will be launched on 2022-09-21 09:54:28.145824 +0200 CEST m=+35.148432288
```

*This example output shows the launch time of the chain on the network.*

