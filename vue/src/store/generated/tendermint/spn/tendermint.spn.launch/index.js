import { txClient, queryClient, MissingWalletError } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { Chain } from "./module/types/launch/chain";
import { DefaultInitialGenesis } from "./module/types/launch/chain";
import { GenesisURL } from "./module/types/launch/chain";
import { ChainNameCount } from "./module/types/launch/chain_name_count";
import { GenesisAccount } from "./module/types/launch/genesis_account";
export { Chain, DefaultInitialGenesis, GenesisURL, ChainNameCount, GenesisAccount };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        ChainNameCount: {},
        ChainNameCountAll: {},
        GenesisAccount: {},
        GenesisAccountAll: {},
        Chain: {},
        ChainAll: {},
        _Structure: {
            Chain: getStructure(Chain.fromPartial({})),
            DefaultInitialGenesis: getStructure(DefaultInitialGenesis.fromPartial({})),
            GenesisURL: getStructure(GenesisURL.fromPartial({})),
            ChainNameCount: getStructure(ChainNameCount.fromPartial({})),
            GenesisAccount: getStructure(GenesisAccount.fromPartial({})),
        },
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(subscription);
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(subscription);
        }
    },
    getters: {
        getChainNameCount: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ChainNameCount[JSON.stringify(params)] ?? {};
        },
        getChainNameCountAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ChainNameCountAll[JSON.stringify(params)] ?? {};
        },
        getGenesisAccount: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.GenesisAccount[JSON.stringify(params)] ?? {};
        },
        getGenesisAccountAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.GenesisAccountAll[JSON.stringify(params)] ?? {};
        },
        getChain: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Chain[JSON.stringify(params)] ?? {};
        },
        getChainAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ChainAll[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: tendermint.spn.launch initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    await dispatch(subscription.action, subscription.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryChainNameCount({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryChainNameCount(key.chainName)).data;
                commit('QUERY', { query: 'ChainNameCount', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryChainNameCount', payload: { options: { all }, params: { ...key }, query } });
                return getters['getChainNameCount']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryChainNameCount', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryChainNameCountAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryChainNameCountAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryChainNameCountAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ChainNameCountAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryChainNameCountAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getChainNameCountAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryChainNameCountAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryGenesisAccount({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryGenesisAccount(key.chainID, key.address)).data;
                commit('QUERY', { query: 'GenesisAccount', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryGenesisAccount', payload: { options: { all }, params: { ...key }, query } });
                return getters['getGenesisAccount']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryGenesisAccount', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryGenesisAccountAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryGenesisAccountAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryGenesisAccountAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'GenesisAccountAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryGenesisAccountAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getGenesisAccountAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryGenesisAccountAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryChain({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryChain(key.chainID)).data;
                commit('QUERY', { query: 'Chain', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryChain', payload: { options: { all }, params: { ...key }, query } });
                return getters['getChain']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryChain', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryChainAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryChainAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryChainAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ChainAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryChainAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getChainAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryChainAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgCreateChain({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateChain(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateChain:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgCreateChain({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateChain(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateChain:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateChain:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
