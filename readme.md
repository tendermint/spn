# Starport Network

> This branch contains the source code for the 1.0 version of Starport Network. To access the source code of the currently deployed proof of concept, visit the [`master-legacy`](https://github.com/tendermint/spn/tree/master-legacy) branch.

**spn** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/tendermint/spn@latest! | sudo bash
```
`tendermint/spn` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

Starport Network (SPN) is a free and open source product maintained by [Tendermint](https://tendermint.com). Here's where you can find us. Stay in touch.

- [Starport.com website](https://starport.com)
- [Starport docs](https://docs.starport.com/)
- [Starport Discord](https://discord.com/starport)
- [@StarportHQ on Twitter](https://twitter.com/StarportHQ)
- [Starport.com/blog](https://starport.com/blog/)
- [Starport YouTube](https://www.youtube.com/channel/UCXMndYLK7OuvjvElSeSWJ1Q)
- [Cosmos SDK docs](https://docs.cosmos.network)
